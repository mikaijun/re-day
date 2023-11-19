package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"gorm.io/gorm"
)

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	loaders := loader.NewLoaders(db)
	loaders.UserLoader.ClearAll()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextLoaderCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)
		r = r.WithContext(nextLoaderCtx)
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		tokenObj, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			var secretKey = "secret"
			return []byte(secretKey), nil
		})

		if err != nil {
			http.Error(w, `{"errors":[{"message":"ヘッダーが無効です"}]}`, http.StatusBadRequest)
			return
		}

		claims := tokenObj.Claims.(jwt.MapClaims)
		userId := claims["user_id"]
		authExpirie := &model.AuthExpirie{}
		db.Where("user_id = ?", userId).First(authExpirie)

		if authExpirie.ExpiresAt.Before(time.Now()) {
			http.Error(w, `{"errors":[{"message":"認証が無効です"}]}`, http.StatusBadRequest)
			return
		}

		nextUserCtx := context.WithValue(r.Context(), model.AuthKey, userId)
		r = r.WithContext(nextUserCtx)
		next.ServeHTTP(w, r)
	})
}
