package service

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/loader"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

type DisableToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func _getToken(tokenString string, db *gorm.DB) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNED_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func _tokenCheck(tokenString string, db *gorm.DB) bool {
	disableToken := DisableToken{}

	if err := db.Where("token = ?", tokenString).First(&disableToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		} else {
			return false
		}
	}

	if disableToken.Token == tokenString {
		return false
	}

	return true
}

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := _getToken(tokenString, db)

		if err != nil {
			http.Error(w, `{"errors":[{"message":"トークンが取得できませんでした"}]}`, http.StatusUnauthorized)
			return
		}

		isEnableToken := _tokenCheck(tokenString, db)

		if !isEnableToken {
			http.Error(w, `{"errors":[{"message":"無効化されたトークンです"}]}`, http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"]

		loaders := loader.NewLoaders(db)
		loaders.UserLoader.ClearAll()
		nextCtx := context.WithValue(r.Context(), loader.LoadersKey, loaders)
		nextCtx = context.WithValue(nextCtx, model.AuthKey, userId)
		services := NewServices(db, nextCtx)
		nextCtx = context.WithValue(nextCtx, ServicesKey, services)
		next.ServeHTTP(w, r.WithContext(nextCtx))
	})
}
