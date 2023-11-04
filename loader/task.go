package loader

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/dataloader"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

// TaskLoader はデータベースからtaskを読み取ります
type TaskLoader struct {
	DB *gorm.DB
}

// BatchGetTasks は、ID によって多くのtaskを取得できるバッチ関数を実装します。
func (u *TaskLoader) BatchGetTasks(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// 単一のクエリで要求されたすべてのtaskを読み取ります
	userIDs := make([]string, len(keys))

	for ix, key := range keys {
		userIDs[ix] = key.String()
	}

	tasksTemp := []*model.Task{}
	if err := u.DB.Where("user_id IN ?", userIDs).Find(&tasksTemp).Error; err != nil {
		err := fmt.Errorf("fail get tasks, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	taskByUserId := map[string][]*model.Task{}
	for _, task := range tasksTemp {
		taskByUserId[task.UserId] = append(taskByUserId[task.UserId], task)
	}

	tasks := make([][]*model.Task, len(userIDs))

	for i, id := range userIDs {
		tasks[i] = taskByUserId[id]
	}

	output := make([]*dataloader.Result, len(tasks))
	for index := range tasks {
		task := tasks[index]
		output[index] = &dataloader.Result{Data: task, Error: nil}
	}
	return output
}

// dataloader.Loadをwrapして型づけした実装
func LoadTask(ctx context.Context, taskID string) ([]*model.Task, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.TaskLoader.Load(ctx, dataloader.StringKey(taskID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.([]*model.Task), nil
}
