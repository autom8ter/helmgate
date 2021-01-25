package gql

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/autom8ter/meshpaas/gen/gql/go/generated"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/logger"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client meshpaaspb.MeshPaasServiceClient
	logger *logger.Logger
}

func NewResolver(client meshpaaspb.MeshPaasServiceClient, logger *logger.Logger) *Resolver {
	return &Resolver{
		client: client,
		logger: logger,
	}
}

func (r *Resolver) QueryHandler() http.Handler {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers:  r,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			auth := initPayload.Authorization()
			ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", auth)
			return ctx, nil
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(&apollotracing.Tracer{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	return r.authMiddleware(srv)
}

func (r *Resolver) authMiddleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		for k, arr := range req.Header {
			if len(arr) > 0 {
				ctx = metadata.AppendToOutgoingContext(ctx, k, arr[0])
			}
		}
		handler.ServeHTTP(w, req.WithContext(ctx))
	}
}

//
//func (r *Resolver) Playground() http.HandlerFunc {
//	return func(w http.ResponseWriter, req *http.Request) {
//		w.Header().Add("Content-Type", "text/html")
//		var playground = template.Must(template.New("playground").Parse(`<!DOCTYPE html>
//<html>
//
//<head>
//  <meta charset=utf-8/>
//  <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
//  <title>Graphik Playground</title>
//  <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css" />
//  <link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react/build/favicon.png" />
//  <script src="//cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
//</head>
//
//<body>
//  <div id="root">
//    <style>
//      body {
//        background-color: rgb(23, 42, 58);
//        font-family: Open Sans, sans-serif;
//        height: 90vh;
//      }
//
//      #root {
//        height: 100%;
//        width: 100%;
//        display: flex;
//        align-items: center;
//        justify-content: center;
//      }
//
//      .loading {
//        font-size: 32px;
//        font-weight: 200;
//        color: rgba(255, 255, 255, .6);
//        margin-left: 20px;
//      }
//
//      img {
//        width: 78px;
//        height: 78px;
//      }
//
//      .title {
//        font-weight: 400;
//      }
//    </style>
//    <img src='//cdn.jsdelivr.net/npm/graphql-playground-react/build/logo.png' alt=''>
//    <div class="loading"> Loading
//      <span class="title">Graphik Playground</span>
//    </div>
//  </div>
//  <script>window.addEventListener('load', function (event) {
// 		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
//      GraphQLPlayground.init(document.getElementById('root'), {
//		endpoint: location.protocol + '//' + location.host + '/graphql',
//		subscriptionsEndpoint: wsProto + '//' + location.host + '/graphql',
//		shareEnabled: true,
//		settings: {
//			'request.credentials': 'same-origin',
//			'prettier.useTabs': true
//		}
//      })
//    })</script>
//</body>
//
//</html>
//`))
//
//		playground.Execute(w, map[string]string{})
//	}
//}
