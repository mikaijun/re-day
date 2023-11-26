package loader

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/dataloader"
	"github.com/mikaijun/gqlgen-tasks/graph/model"

	"gorm.io/gorm"
)

type ActionLoader struct {
	DB *gorm.DB
}

// BatchGetUsers は、ID によって多くのユーザーを取得できるバッチ関数を実装します。
func (u *ActionLoader) BatchGetActions(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	TaskIDs := make([]string, len(keys))
	for ix, key := range keys {
		TaskIDs[ix] = key.String()
	}

	taskTemp := []*model.Task{}

	if err := u.DB.Where("id IN ?", TaskIDs).Find(&taskTemp).Error; err != nil {
		err := fmt.Errorf("fail get task, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	tasksByTaskId := map[string]*model.Task{}
	for _, task := range taskTemp {
		tasksByTaskId[task.ID] = task
	}

	tasks := make([]*model.Task, len(TaskIDs))

	for i, id := range TaskIDs {
		tasks[i] = tasksByTaskId[id]
	}

	output := make([]*dataloader.Result, len(keys))
	for index := range TaskIDs {
		task := tasks[index]
		output[index] = &dataloader.Result{Data: task, Error: nil}
	}
	return output
}

func LoadActionByTaskId(ctx context.Context, TaskID string) (*model.Task, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.ActionLoader.Load(ctx, dataloader.StringKey(TaskID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*model.Task), nil
}
