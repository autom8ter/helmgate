package core

import (
	"bytes"
	"context"
	"fmt"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
)

func (m *Manager) StreamLogs(ctx context.Context, opts *meshpaaspb.LogOpts) (chan string, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}

	pods, err := m.kclient.Pods(cast.ToString(usr[m.namespaceClaim])).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		Watch:         false,
		LabelSelector: fmt.Sprintf("meshpaas.app = %s", opts.Name),
	})
	if err != nil {
		return nil, err
	}
	if len(pods.Items) == 0 {
		return nil, errors.New("zero replicas")
	}
	var (
		sinceSeconds *int64
		tailLines    *int64
	)

	if opts.SinceSeconds > 0 {
		sinceSeconds = &opts.SinceSeconds
	}
	if opts.TailLines > 0 {
		sinceSeconds = &opts.TailLines
	}
	return m.streamLogs(ctx, pods, &corev1.PodLogOptions{
		TypeMeta:     v1.TypeMeta{},
		Container:    opts.Container,
		Follow:       opts.Stream,
		Previous:     opts.Previous,
		SinceSeconds: sinceSeconds,
		Timestamps:   true,
		TailLines:    tailLines,
	})
}

func (m *Manager) streamLogs(ctx context.Context, pods *corev1.PodList, opts *corev1.PodLogOptions) (chan string, error) {
	logChan := make(chan string)
	var streamMu = sync.RWMutex{}
	for _, pod := range pods.Items {
		go func(p corev1.Pod) {
			closer, err := m.kclient.GetLogs(ctx, p.Name, p.Namespace, opts)
			defer closer.Close()
			if err != nil {
				m.logger.Error("failed to stream pod logs",
					zap.Error(err),
					zap.Any("options", opts),
				)
				return
			}
			streamctx, cancel := context.WithCancel(ctx)
			defer cancel()
			for {
				select {
				case <-streamctx.Done():
					return
				default:
					buf := make([]byte, 1024)
					_, err := closer.Read(buf)
					if err != nil {
						if err == io.EOF {
							return
						}
					}
					streamMu.Lock()
					logChan <- string(bytes.Trim(buf, "\x00"))
					streamMu.Unlock()
				}
			}
		}(pod)
	}
	return logChan, nil
}
