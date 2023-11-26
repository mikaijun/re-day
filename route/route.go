package route

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgrijalva/jwt-go"
	"github.com/mikaijun/gqlgen-tasks/db"
	"github.com/mikaijun/gqlgen-tasks/graph"
	"github.com/mikaijun/gqlgen-tasks/graph/model"
	"github.com/mikaijun/gqlgen-tasks/middleware"
)

type DisableToken struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type req struct {
	Id string `json:"id"`
}

func Route() {
	db := db.ConnectGORM()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	http.Handle("/query", middleware.Middleware(db, srv))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		req := &req{}
		dec.Decode(req)

		user := &model.User{}
		if err := db.Where("id = ?", req.Id).First(user).Error; err != nil {
			http.Error(w, "IDが間違ってます", http.StatusBadRequest)
		}

		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims = jwt.MapClaims{
			"user_id": req.Id,
			"exp":     time.Now().AddDate(0, 0, 1).Unix(),
		}
		tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNED_KEY")))
		w.Write([]byte(tokenString))
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNED_KEY")), nil
		})

		if err != nil {
			http.Error(w, "トークンが無効です", http.StatusUnauthorized)
			return
		}

		disableToken := DisableToken{
			Token:     tokenString,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&disableToken).Error; err != nil {
			http.Error(w, "ログアウトできませんでした", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("ログアウトしました"))
	})
}
