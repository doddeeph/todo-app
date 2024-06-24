package repository

import (
	"context"
	"github.com/doddeeph/todo-app/domain"
	"github.com/doddeeph/todo-app/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	db         mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{db: db, collection: collection}
}

func (tr *taskRepository) Create(ctx context.Context, task *domain.Task) error {
	collection := tr.db.Collection(tr.collection)
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) FetchByUserId(ctx context.Context, userId string) ([]domain.Task, error) {
	collection := tr.db.Collection(tr.collection)
	var tasks []domain.Task
	idHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return tasks, err
	}
	cursor, err := collection.Find(ctx, bson.M{"userId": idHex})
	if err != nil {
		return tasks, err
	}
	err = cursor.All(ctx, &tasks)
	if tasks == nil {
		return []domain.Task{}, err
	}
	return tasks, err
}
