package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"github.com/mikaijun/gqlgen-tasks/service"
)

// Task is the resolver for the task field.
func (r *actionResolver) Task(ctx context.Context, obj *model.Action) (*model.Task, error) {
	action, err := loader.LoadAction(ctx, obj.TaskId)
	if err != nil {
		return nil, err
	}
	return action, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *actionResolver) CreatedAt(ctx context.Context, obj *model.Action) (string, error) {
	return obj.CreatedAt.Format("2006-01-02 15:04:05"), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *actionResolver) UpdatedAt(ctx context.Context, obj *model.Action) (string, error) {
	return obj.UpdatedAt.Format("2006-01-02 15:04:05"), nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	return service.LoginFunc(r.DB, input.ID)
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	userId := ctx.Value(model.AuthKey).(string)
	task := model.Task{
		Content:   input.Content,
		ID:        uuid.New().String(),
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.DB.Create(&task)
	return &task, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		ID:   uuid.New().String(),
		Name: input.Name,
	}
	r.DB.Create(&user)
	return &user, nil
}

// CreateAction is the resolver for the createAction field.
func (r *mutationResolver) CreateAction(ctx context.Context, input model.NewAction) (*model.Action, error) {
	actions := model.Action{
		ID:        uuid.New().String(),
		TaskId:    input.TaskId,
		Score:     input.Score,
		Comment:   input.Comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.DB.Create(&actions)
	return &actions, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	tasks := []*model.Task{}
	r.DB.Find(&tasks)
	return tasks, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := []*model.User{}
	r.DB.Find(&user)
	return user, nil
}

// Actions is the resolver for the actions field.
func (r *queryResolver) Actions(ctx context.Context) ([]*model.Action, error) {
	action := []*model.Action{}
	r.DB.Find(&action)
	return action, nil
}

// User is the resolver for the user field.
func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	user, err := loader.LoadUser(ctx, obj.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *taskResolver) CreatedAt(ctx context.Context, obj *model.Task) (string, error) {
	return obj.CreatedAt.Format("2006-01-02 15:04:05"), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *taskResolver) UpdatedAt(ctx context.Context, obj *model.Task) (string, error) {
	return obj.UpdatedAt.Format("2006-01-02 15:04:05"), nil
}

// Tasks is the resolver for the tasks field.
func (r *userResolver) Tasks(ctx context.Context, obj *model.User) ([]*model.Task, error) {
	task, err := loader.LoadTask(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// CreatedAt is the resolver for the created_at field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.CreatedAt.Format("2006-01-02 15:04:05"), nil
}

// UpdatedAt is the resolver for the updated_at field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (string, error) {
	return obj.UpdatedAt.Format("2006-01-02 15:04:05"), nil
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
