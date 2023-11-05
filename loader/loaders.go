package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

type ctxKey string

const (
	LoadersKey = ctxKey("dataloaders")
)

// 各DataLoaderを取りまとめるstruct
type Loaders struct {
	UserLoader   *dataloader.Loader
	TaskLoader   *dataloader.Loader
	ActionLoader *dataloader.Loader
}

// Loadersの初期化メソッド
func NewLoaders(db *gorm.DB) *Loaders {

	//ローダーの定義
	userLoader := &UserLoader{
		DB: db,
	}
	taskLoader := &TaskLoader{
		DB: db,
	}
	actionLoader := &ActionLoader{
		DB: db,
	}
	loaders := &Loaders{
		UserLoader:   dataloader.NewBatchedLoader(userLoader.BatchGetUsers),
		TaskLoader:   dataloader.NewBatchedLoader(taskLoader.BatchGetTasks),
		ActionLoader: dataloader.NewBatchedLoader(actionLoader.BatchGetActions),
	}
	return loaders
}

// ContextからLoadersを取得する
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(LoadersKey).(*Loaders)
}
