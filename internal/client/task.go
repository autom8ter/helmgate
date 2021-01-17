package client

import (
	"context"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateTask(ctx context.Context, task *kdeploypb.TaskConstructor) (*kdeploypb.Task, error) {
	kapp := &k8sTask{}
	namespace, err := m.client.Namespaces().Get(ctx, task.Namespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.client.Namespaces().Create(ctx, toTaskNamespace(task), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	tsk, err := toTask(task)
	if err != nil {
		return nil, err
	}
	cronJob, err := m.client.CronJobs(task.Namespace).Create(ctx, tsk, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) UpdateTask(ctx context.Context, task *kdeploypb.TaskUpdate) (*kdeploypb.Task, error) {
	kapp := &k8sTask{}
	namespace, err := m.client.Namespaces().Get(ctx, task.Namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	cronJob, err := m.client.CronJobs(task.Namespace).Get(ctx, task.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cronJob, err = overwriteTask(cronJob, task)
	if err != nil {
		return nil, err
	}
	cronJob, err = m.client.CronJobs(task.Namespace).Update(ctx, cronJob, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) GetTask(ctx context.Context, ref *kdeploypb.Ref) (*kdeploypb.Task, error) {
	kapp := &k8sTask{}

	ns, err := m.client.Namespaces().Get(ctx, ref.Namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	cronJob, err := m.client.CronJobs(ref.Namespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) DeleteTask(ctx context.Context, ref *kdeploypb.Ref) error {
	if err := m.client.CronJobs(ref.Namespace).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *Manager) ListTasks(ctx context.Context, namespace *kdeploypb.Namespace) (*kdeploypb.Tasks, error) {
	var kapps = &kdeploypb.Tasks{}
	ns, err := m.client.Namespaces().Get(ctx, namespace.GetNamespace(), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cronJobs, err := m.client.CronJobs(namespace.GetNamespace()).List(ctx, v1.ListOptions{
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
