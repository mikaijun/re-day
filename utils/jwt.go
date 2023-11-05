package utils

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
	// ユーザートークンを生成する
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id": 1,
	}
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		// トークンの生成に失敗した場合
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// トークンをレスポンスとして返す
	w.Header().Set("Authorization", "Bearer "+tokenString)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}
