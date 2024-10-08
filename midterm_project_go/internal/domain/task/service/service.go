package service

import (
	"midterm_cloud_project_2024/internal/domain/task/entity"
	"midterm_cloud_project_2024/internal/domain/task/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() ([]entity.Task, error) {
	return s.repo.GetTasks()
}

func (s *TaskService) AddTask(task entity.Task) error {
	return s.repo.AddTask(task)
}
