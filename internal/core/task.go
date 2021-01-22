package core

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateTask(ctx context.Context, task *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sTask{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	tsk, err := m.toTask(usr, task)
	if err != nil {
		return nil, err
	}
	cronJob, err := m.kclient.CronJobs(cast.ToString(usr[m.namespaceClaim])).Create(ctx, tsk, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) UpdateTask(ctx context.Context, task *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sTask{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	cronJob, err := m.kclient.CronJobs(cast.ToString(usr[m.namespaceClaim])).Get(ctx, task.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cronJob, err = overwriteTask(cronJob, task)
	if err != nil {
		return nil, err
	}
	cronJob, err = m.kclient.CronJobs(cast.ToString(usr[m.namespaceClaim])).Update(ctx, cronJob, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) GetTask(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Task, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sTask{}
	ns, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	cronJob, err := m.kclient.CronJobs(cast.ToString(usr[m.namespaceClaim])).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) DeleteTask(ctx context.Context, ref *meshpaaspb.Ref) error {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	if err := m.kclient.CronJobs(cast.ToString(usr[m.namespaceClaim])).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *Manager) ListTasks(ctx context.Context) (*meshpaaspb.Tasks, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	var kapps = &meshpaaspb.Tasks{}
	ns, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cronJobs, err := m.kclient.CronJobs(cast.ToString(usr[m.namespaceClaim])).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	for _, cronJob := range cronJobs.Items {
		kapp := &k8sTask{
			namespace: ns,
			cronJob:   &cronJob,
		}
		kapps.Tasks = append(kapps.Tasks, kapp.toTask())
	}
	return kapps, nil
}
