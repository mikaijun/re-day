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
	UserService   *UserService
}

func NewServices(db *gorm.DB, ctx context.Context) *Services {
	actionService := &ActionService{db}
	taskService := &TaskService{db, ctx}
	userService := &UserService{db, ctx}
	services := &Services{
		ActionService: actionService,
		TaskService:   taskService,
		UserService:   userService,
	}
	return services
}

func GetServices(ctx context.Context) *Services {
	return ctx.Value(ServicesKey).(*Services)
}
