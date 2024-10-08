package repository

import (
	"midterm_cloud_project_2024/internal/domain/task/entity"
	"sync"
)

type TaskRepository struct {
	mu    sync.Mutex
	tasks map[string]entity.Task
}

func NewRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]entity.Task),
	}
}

func (r *TaskRepository) GetTasks() ([]entity.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]entity.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) AddTask(task entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tasks[task.ID] = task
	return nil
}
