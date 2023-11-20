package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

func CreateAction(input model.NewAction, db *gorm.DB) (*model.Action, error) {
	actions := model.Action{
		ID:        uuid.New().String(),
		TaskId:    input.TaskId,
		Score:     input.Score,
		Comment:   input.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(&actions).Error; err != nil {
		return nil, errors.New("行動を生成できませんでした")
	}
	return &actions, nil
}

func FindActions(db *gorm.DB) ([]*model.Action, error) {
	action := []*model.Action{}
	db.Find(&action)
	return action, nil
}
