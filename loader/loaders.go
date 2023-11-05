package loader

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/graph-gophers/dataloader"
	"gorm.io/gorm"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
	authKey    = ctxKey("auth")
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

// ミドルウェアはデータ ローダーをコンテキストに挿入します
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	loaders.UserLoader.ClearAll()
	// ローダーをリクエストコンテキストに挿入するミドルウェアを返す
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		token := r.Header.Get("Authorization")
		// JWT オブジェクトを作成します。
		tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("my-secret-key"), nil
		})
		if err != nil {
			fmt.Print(err)
			return
		}

		// トークンのクレームからユーザー ID を取得します。
		claims := tokenObj.Claims.(jwt.MapClaims)
		id := claims["id"]
		fmt.Print(id)
		nextCtx2 := context.WithValue(r.Context(), authKey, "1")
		r = r.WithContext(nextCtx2)
		next.ServeHTTP(w, r)
	})
}

// ContextからLoadersを取得する
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// ContextからLoadersを取得する
func GetLoaders2(ctx context.Context) string {
	return ctx.Value(authKey).(string)
}
