package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"github.com/mikaijun/gqlgen-tasks/service"
)

func Middleware(loaders *loader.Loaders, next http.Handler) http.Handler {
	loaders.UserLoader.ClearAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)
		r = r.WithContext(nextCtx)
		token := r.Header.Get("Authorization")
		tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("my-secret-key"), nil
		})
		if err != nil {
			fmt.Print(err)
			return
		}
		claims := tokenObj.Claims.(jwt.MapClaims)
		id := claims["id"]
		fmt.Print(id)
		nextCtx2 := context.WithValue(r.Context(), service.AuthKey, "1")
		r = r.WithContext(nextCtx2)
		next.ServeHTTP(w, r)
	})
}
