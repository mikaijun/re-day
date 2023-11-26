package service

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"gorm.io/gorm"
)

func _getToken(tokenString string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, nil
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNED_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func _getUserId(token *jwt.Token, db *gorm.DB) (string, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	return userId.(string), nil
}

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	loaders := loader.NewLoaders(db)
	loaders.UserLoader.ClearAll()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := _getToken(tokenString)
		nextCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)

		if err != nil {
			http.Error(w, `{"errors":[{"message":"ヘッダーが無効です"}]}`, http.StatusBadRequest)
			return
		}

		if token != nil {
			userId, err := _getUserId(token, db)
			if err != nil {
				http.Error(w, `{"errors":[{"message":"認証が無効です"}]}`, http.StatusBadRequest)
				return
			}
			nextCtx = context.WithValue(nextCtx, model.AuthKey, userId)
		}

		services := NewServices(db, nextCtx)
		nextCtx = context.WithValue(nextCtx, ServicesKey, services)

		next.ServeHTTP(w, r.WithContext(nextCtx))
	})
}
