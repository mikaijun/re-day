package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/loader"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	ctx context.Context
}

func (s *UserService) CreateUser(name string) (*model.User, error) {
	user := model.User{
		ID:   uuid.New().String(),
		Name: name,
	}
	s.db.Create(&user)
	return &user, nil
}

func (s *UserService) FindUsers() ([]*model.User, error) {
	user := []*model.User{}
	s.db.Find(&user)
	return user, nil
}

func (s *UserService) FindUserByTask(task *model.Task) (*model.User, error) {
	user, err := loader.LoadUser(s.ctx, task.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
