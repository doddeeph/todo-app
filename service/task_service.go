package service

import (
	"context"
	"github.com/doddeeph/todo-app/domain"
	"time"
)

type taskService struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskService(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskService {
	return &taskService{taskRepository: taskRepository, contextTimeout: timeout}
}

func (ts *taskService) Create(ctx context.Context, task *domain.Task) error {
	newCtx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()
	return ts.taskRepository.Create(newCtx, task)
}

func (ts *taskService) FetchByUserId(ctx context.Context, userId string) ([]domain.Task, error) {
	newCtx, cancel := context.WithTimeout(ctx, ts.contextTimeout)
	defer cancel()
	return ts.taskRepository.FetchByUserId(newCtx, userId)
}
