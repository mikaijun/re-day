package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"gorm.io/gorm"
)

func FindTaskByAction(ctx context.Context, action *model.Action) (*model.Task, error) {
	task, err := loader.LoadActionByTaskId(ctx, action.TaskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func CreateTask(ctx context.Context, content string, db *gorm.DB) (*model.Task, error) {
	userId := ctx.Value(model.AuthKey).(string)

	task := model.Task{
		Content:   content,
		ID:        uuid.New().String(),
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&task).Error; err != nil {
		return nil, errors.New("タスクを生成できませんでした")
	}
	return &task, nil
}

func FindTasks(db *gorm.DB) ([]*model.Task, error) {
	tasks := []*model.Task{}
	db.Find(&tasks)
	return tasks, nil
}

func FindTasksByUser(ctx context.Context, user *model.User) ([]*model.Task, error) {
	task, err := loader.LoadTask(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
