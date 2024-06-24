package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Title  string             `json:"title" bson:"title" form:"title" binding:"required`
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
}

type TaskService interface {
	Create(ctx context.Context, task *Task) error
	FetchByUserId(ctx context.Context, userId string) ([]Task, error)
}

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	FetchByUserId(ctx context.Context, userId string) ([]Task, error)
}
