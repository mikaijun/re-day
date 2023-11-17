package service

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

type ctxKey string

const (
	AuthKey = ctxKey("auth")
)

func GetUserByToken(db *gorm.DB, token string) (*model.User, error) {
	tokenObj, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("my-secret-key"), nil
	})

	claims := tokenObj.Claims.(jwt.MapClaims)
	id := claims["id"]
	user := &model.User{}
	db.Where("id = ?", id).First(user)
	return user, nil
}

func GetUserByContext(ctx context.Context) *model.User {
	return ctx.Value(AuthKey).(*model.User)
}
