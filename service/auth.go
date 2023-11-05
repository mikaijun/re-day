package service

import (
	"context"
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

// ContextからLoadersを取得する
func GetLoaders2(ctx context.Context) string {
	return ctx.Value(AuthKey).(string)
}
