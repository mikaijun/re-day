package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/loader"
	"github.com/mikaijun/gqlgen-tasks/graph/model"

	"gorm.io/gorm"
)

type TaskService struct {
	db  *gorm.DB
	ctx context.Context
}

func (s *TaskService) FindTaskByAction(action *model.Action) (*model.Task, error) {
	task, err := loader.LoadActionByTaskId(s.ctx, action.TaskId)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(newTask model.NewTask) (*model.Task, error) {
	userId := s.ctx.Value(model.AuthKey).(string)

	task := model.Task{
		Content:   newTask.Content,
		ID:        uuid.New().String(),
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(&task).Error; err != nil {
		return nil, errors.New("タスクを生成できませんでした")
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(updateTask model.UpdateTask) (*model.Task, error) {
	task := model.Task{}

	if err := s.db.Where("id = ?", updateTask.TaskId).First(&task).Error; err != nil {
		return nil, errors.New("指定したタスクが存在しません")
	}

	task.Content = updateTask.Content

	if err := s.db.Save(&task).Error; err != nil {
		return nil, errors.New("タスクを更新できませんでした")
	}

	return &task, nil
}

func (s *TaskService) FindTasks() ([]*model.Task, error) {
	tasks := []*model.Task{}
	s.db.Find(&tasks)
	return tasks, nil
}

func (s *TaskService) FindTasksByUser(user *model.User) ([]*model.Task, error) {
	task, err := loader.LoadTask(s.ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
