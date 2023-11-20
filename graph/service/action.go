package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

type ActionService struct {
	db *gorm.DB
}

func (s *ActionService) CreateAction(input model.NewAction) (*model.Action, error) {
	actions := model.Action{
		ID:        uuid.New().String(),
		TaskId:    input.TaskId,
		Score:     input.Score,
		Comment:   input.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.db.Create(&actions).Error; err != nil {
		return nil, errors.New("行動を生成できませんでした")
	}
	return &actions, nil
}

func (s *ActionService) FindActions() ([]*model.Action, error) {
	action := []*model.Action{}
	s.db.Find(&action)
	return action, nil
}
