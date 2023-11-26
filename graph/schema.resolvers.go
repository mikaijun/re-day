package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/graph/service"
)

// Task is the resolver for the task field.
func (r *actionResolver) Task(ctx context.Context, obj *model.Action) (*model.Task, error) {
	service := service.GetServices(ctx)
	return service.TaskService.FindTaskByAction(obj)
}

// CreatedAt is the resolver for the created_at field.
func (r *actionResolver) CreatedAt(ctx context.Context, obj *model.Action) (string, error) {
	return service.FormatStringToDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *actionResolver) UpdatedAt(ctx context.Context, obj *model.Action) (string, error) {
	return service.FormatStringToDate(obj.UpdatedAt), nil
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	service := service.GetServices(ctx)
	return service.TaskService.CreateTask(input)
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTask) (*model.Task, error) {
	service := service.GetServices(ctx)
	return service.TaskService.UpdateTask(input)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	service := service.GetServices(ctx)
	return service.UserService.CreateUser(input.Name)
}

// CreateAction is the resolver for the createAction field.
func (r *mutationResolver) CreateAction(ctx context.Context, input model.NewAction) (*model.Action, error) {
	service := service.GetServices(ctx)
	return service.ActionService.CreateAction(input)
}

// UpdateAction is the resolver for the updateAction field.
func (r *mutationResolver) UpdateAction(ctx context.Context, input model.UpdateAction) (*model.Action, error) {
	service := service.GetServices(ctx)
	return service.ActionService.UpdateAction(input)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	service := service.GetServices(ctx)
	return service.TaskService.FindTasks()
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	service := service.GetServices(ctx)
	return service.UserService.FindUsers()
}

// Actions is the resolver for the actions field.
func (r *queryResolver) Actions(ctx context.Context) ([]*model.Action, error) {
	service := service.GetServices(ctx)
	return service.ActionService.FindActions()
}

// User is the resolver for the user field.
func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	service := service.GetServices(ctx)
	return service.UserService.FindUserByTask(obj)
}

// CreatedAt is the resolver for the created_at field.
func (r *taskResolver) CreatedAt(ctx context.Context, obj *model.Task) (string, error) {
	return service.FormatStringToDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *taskResolver) UpdatedAt(ctx context.Context, obj *model.Task) (string, error) {
	return service.FormatStringToDate(obj.UpdatedAt), nil
}

// Tasks is the resolver for the tasks field.
func (r *userResolver) Tasks(ctx context.Context, obj *model.User) ([]*model.Task, error) {
	service := service.GetServices(ctx)
	return service.TaskService.FindTasksByUser(obj)
}

// CreatedAt is the resolver for the created_at field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	return service.FormatStringToDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (string, error) {
	return service.FormatStringToDate(obj.UpdatedAt), nil
}

// Action returns ActionResolver implementation.
func (r *Resolver) Action() ActionResolver { return &actionResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type actionResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
