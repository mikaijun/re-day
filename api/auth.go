package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type req struct {
	Id string `json:"id"`
}

func Auth(db *gorm.DB) {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		req := &req{}
		dec.Decode(req)
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"id": req.Id,
		}
		tokenString, _ := token.SignedString([]byte("my-secret-key"))
		w.Write([]byte(tokenString))
	})
}
