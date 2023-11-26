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

func (s *ActionService) UpdateAction(input model.UpdateAction) (*model.Action, error) {
	action := model.Action{}

	if err := s.db.Where("id = ?", input.ActionId).First(&action).Error; err != nil {
		return nil, errors.New("指定した行動が存在しません")
	}

	action.Comment = input.Comment
	action.Score = input.Score

	if err := s.db.Save(&action).Error; err != nil {
		return nil, errors.New("行動を更新できませんでした")
	}

	return &action, nil
}

func (s *ActionService) FindActions() ([]*model.Action, error) {
	action := []*model.Action{}
	s.db.Find(&action)
	return action, nil
}
