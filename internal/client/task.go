package client

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateTask(ctx context.Context, task *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	kapp := &k8sTask{}
	namespace, err := m.kclient.Namespaces().Get(ctx, task.Project, v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toTaskNamespace(task), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	tsk, err := toTask(task)
	if err != nil {
		return nil, err
	}
	cronJob, err := m.kclient.CronJobs(task.Project).Create(ctx, tsk, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) UpdateTask(ctx context.Context, task *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	kapp := &k8sTask{}
	namespace, err := m.kclient.Namespaces().Get(ctx, task.Project, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	cronJob, err := m.kclient.CronJobs(task.Project).Get(ctx, task.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cronJob, err = overwriteTask(cronJob, task)
	if err != nil {
		return nil, err
	}
	cronJob, err = m.kclient.CronJobs(task.Project).Update(ctx, cronJob, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) GetTask(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Task, error) {
	kapp := &k8sTask{}

	ns, err := m.kclient.Namespaces().Get(ctx, ref.Project, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	cronJob, err := m.kclient.CronJobs(ref.Project).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.cronJob = cronJob
	return kapp.toTask(), nil
}

func (m *Manager) DeleteTask(ctx context.Context, ref *meshpaaspb.Ref) error {
	if err := m.kclient.CronJobs(ref.Project).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *Manager) ListTasks(ctx context.Context, project *meshpaaspb.ProjectRef) (*meshpaaspb.Tasks, error) {
	var kapps = &meshpaaspb.Tasks{}
	ns, err := m.kclient.Namespaces().Get(ctx, project.GetName(), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	cronJobs, err := m.kclient.CronJobs(project.GetName()).List(ctx, v1.ListOptions{
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
