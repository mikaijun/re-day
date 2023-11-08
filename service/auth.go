package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"gorm.io/gorm"
)

type ctxKey string
type req struct {
	Id string `json:"id"`
}

const (
	AuthKey = ctxKey("auth")
)

func Login(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	req := &req{}
	dec.Decode(req)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"id": req.Id,
	}
	tokenString, _ := token.SignedString([]byte("my-secret-key"))
	w.Write([]byte(tokenString))
}

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
