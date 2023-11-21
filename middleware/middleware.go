package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/graph/service"
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

func _getUserId(token *jwt.Token, db *gorm.DB, w http.ResponseWriter) (string, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	authExpirie := &model.AuthExpirie{}
	db.Where("user_id = ?", userId).First(authExpirie)

	if authExpirie.ExpiresAt.Before(time.Now()) {
		return "", errors.New("認証が無効です")
	}
	return userId.(string), nil
}

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		token, err := _getToken(tokenString)

		if token == nil {
			return
		}

		if err != nil {
			http.Error(w, `{"errors":[{"message":"ヘッダーが無効です"}]}`, http.StatusBadRequest)
			return
		}

		userId, err := _getUserId(token, db, w)
		if err != nil {
			http.Error(w, `{"errors":[{"message":"認証が無効です"}]}`, http.StatusBadRequest)
			return
		}

		loaders := loader.NewLoaders(db)
		services := service.NewServices(db)
		loaders.UserLoader.ClearAll()

		nextLoaderCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)
		nextServicesCtx := context.WithValue(nextLoaderCtx, service.ServicesKey, services)
		nextUserCtx := context.WithValue(nextServicesCtx, model.AuthKey, userId)

		next.ServeHTTP(w, r.WithContext(nextUserCtx))
	})
}
