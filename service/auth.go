package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

type ctxKey string

const (
	AuthKey = ctxKey("auth")
)

func Login(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id": "1550a82d-00f8-410c-9e82-e5e3cfe64e42",
	}
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+tokenString)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func GetUserByToken(db *gorm.DB, token string) (*model.User, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("my-secret-key"), nil
	})

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	claims := tokenObj.Claims.(jwt.MapClaims)
	id := claims["id"]
	user := &model.User{}
	db.Where("id = ?", id).First(user)
	return user, nil
}

func GetUserByContext(ctx context.Context) *model.User {
	return ctx.Value(AuthKey).(*model.User)
}
