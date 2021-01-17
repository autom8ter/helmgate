package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/kdeploy/helpers"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Manager) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return nil, err
		}
		ctx = c.tokenToContext(ctx, token)
		tokenHash := helpers.Hash([]byte(token))
		if val, ok := c.jwtCache.Get(tokenHash); ok {
			payload := val.(map[string]interface{})
			ctx, err := c.checkRequest(ctx, info.FullMethod, req, payload)
			if err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
		ctx = c.methodToContext(ctx, info.FullMethod)
		userinfoReq, err := http.NewRequest(http.MethodGet, c.userInfoEndpoint, nil)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo(%s): %s", c.userInfoEndpoint, err.Error())
		}
		userinfoReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(userinfoReq)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %v", resp.StatusCode)
		}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		payload := map[string]interface{}{}
		if err := json.Unmarshal(bits, &payload); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		c.jwtCache.Set(tokenHash, payload, 1*time.Hour)
		ctx, err = c.checkRequest(ctx, info.FullMethod, req, payload)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (c *Manager) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := c.methodToContext(ss.Context(), info.FullMethod)
		token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return err
		}
		ctx = c.tokenToContext(ctx, token)
		tokenHash := helpers.Hash([]byte(token))
		if val, ok := c.jwtCache.Get(tokenHash); ok {
			payload := val.(map[string]interface{})
			ctx, err := c.checkRequest(ctx, info.FullMethod, srv, payload)
			if err != nil {
				return err
			}
			wrapped := grpc_middleware.WrapServerStream(ss)
			wrapped.WrappedContext = ctx
			return handler(srv, wrapped)
		}
		userinfoReq, err := http.NewRequest(http.MethodGet, c.userInfoEndpoint, nil)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		userinfoReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(userinfoReq)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %v", resp.StatusCode)
		}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "failed to get userinfo")
		}
		payload := map[string]interface{}{}
		if err := json.Unmarshal(bits, &payload); err != nil {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		c.jwtCache.Set(tokenHash, payload, 1*time.Hour)
		ctx, err = c.checkRequest(ctx, info.FullMethod, srv, payload)
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func (a *Manager) userToContext(ctx context.Context, payload map[string]interface{}) (context.Context, error) {
	return context.WithValue(ctx, authCtxKey, payload), nil
}

func (s *Manager) getIdentity(ctx context.Context) map[string]interface{} {
	val, ok := ctx.Value(authCtxKey).(map[string]interface{})
	if ok {
		return val
	}
	return map[string]interface{}{}
}

func (r *Manager) getMethod(ctx context.Context) string {
	val, ok := ctx.Value(methodCtxKey).(string)
	if ok {
		return val
	}
	return ""
}

func (r *Manager) methodToContext(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, methodCtxKey, path)
}

func (r *Manager) tokenToContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtxKey, token)
}

func (r *Manager) getToken(ctx context.Context) string {
	val, ok := ctx.Value(tokenCtxKey).(string)
	if ok {
		return val
	}
	return ""
}

func (c *Manager) checkRequest(ctx context.Context, method string, req interface{}, payload map[string]interface{}) (context.Context, error) {
	ctx = c.methodToContext(ctx, method)
	ctx, err := c.userToContext(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}
	if payload["email"] != nil && c.isRootUser(payload["email"].(string)) {
		return ctx, nil
	}
	if len(c.requestAuthorizers) == 0 {
		return ctx, nil
	}
	authorizer := map[string]interface{}{
		"user":   payload,
		"method": method,
	}
	meta := map[string]interface{}{}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, val := range md {
			if len(val) > 0 {
				meta[k] = val[0]
			}
		}
	}
	authorizer["metadata"] = meta
	msg := req.(proto.Message)
	bits, _ := helpers.MarshalJSON(msg)
	reqMap := map[string]interface{}{}
	if err := json.Unmarshal(bits, &reqMap); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	authorizer["request"] = reqMap
	for _, a := range c.requestAuthorizers {
		if err := a.Eval(authorizer); err == nil {
			return ctx, nil
		}
	}
	return nil, status.Errorf(codes.PermissionDenied, "request from %s.%s  authorization = denied", payload["sub"], payload["email"])
}

func (c *Manager) isRootUser(email string) bool {
	return helpers.ContainsString(email, c.rootUsers)
}
