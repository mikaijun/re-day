package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mikaijun/gqlgen-tasks/loader"
	"github.com/mikaijun/gqlgen-tasks/service"
)

func Middleware(loaders *loader.Loaders, next http.Handler) http.Handler {
	loaders.UserLoader.ClearAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextLoaderCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)
		r = r.WithContext(nextLoaderCtx)

		token := r.Header.Get("Authorization")
		id := service.GetUserId(token)
		fmt.Print(id)
		nextUserCtx := context.WithValue(r.Context(), service.AuthKey, "1")
		r = r.WithContext(nextUserCtx)
		next.ServeHTTP(w, r)
	})
}
