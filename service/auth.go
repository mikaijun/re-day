package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type ctxKey string

const (
	AuthKey = ctxKey("auth")
)

func GenerateToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id": 1,
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

func GetUserId(token string) float64 {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("my-secret-key"), nil
	})

	if err != nil {
		fmt.Print(err)
		return 0
	}

	claims := tokenObj.Claims.(jwt.MapClaims)
	id := claims["id"]
	return id.(float64)
}

// ContextからLoadersを取得する
func GetLoaders2(ctx context.Context) string {
	return ctx.Value(AuthKey).(string)
}
