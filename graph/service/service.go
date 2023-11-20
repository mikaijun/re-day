package service

import (
	"context"

	"gorm.io/gorm"
)

type ctxKey string

const (
	ServicesKey = ctxKey("services")
)

type Services struct {
	ActionService *ActionService
	TaskService   *TaskService
}

func NewServices(db *gorm.DB) *Services {

	actionService := &ActionService{db}
	taskService := &TaskService{db}
	services := &Services{
		ActionService: actionService,
		TaskService:   taskService,
	}
	return services
}

func GetServices(ctx context.Context) *Services {
	return ctx.Value(ServicesKey).(*Services)
}
