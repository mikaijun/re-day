package middleware

import (
	"context"
	"net/http"

	"github.com/mikaijun/gqlgen-tasks/loader"
	"github.com/mikaijun/gqlgen-tasks/service"
	"gorm.io/gorm"
)

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	loaders := loader.NewLoaders(db)
	loaders.UserLoader.ClearAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextLoaderCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)
		r = r.WithContext(nextLoaderCtx)

		token := r.Header.Get("Authorization")
		user, err := service.GetUserByToken(db, token)
		if err != nil {
			panic(err.Error())
		}
		nextUserCtx := context.WithValue(r.Context(), service.AuthKey, user)
		r = r.WithContext(nextUserCtx)
		next.ServeHTTP(w, r)
	})
}
