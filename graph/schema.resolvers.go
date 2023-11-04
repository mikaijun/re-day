package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	//ランダムな数字の生成
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	task := model.Task{
		Text:      input.Text,
		ID:        fmt.Sprintf("T%d", rand),
		UserId:    input.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.DB.Create(&task)
	return &task, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	//ランダムな値の生成
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	user := model.User{
		ID:   fmt.Sprintf("U%d", rand),
		Name: input.Name,
	}
	r.DB.Create(&user)
	return &user, nil
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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Task returns TaskResolver implementation.
func (r *Resolver) Task() TaskResolver { return &taskResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
