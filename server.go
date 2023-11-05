package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/mikaijun/gqlgen-tasks/db"
	"github.com/mikaijun/gqlgen-tasks/graph"
	"github.com/mikaijun/gqlgen-tasks/loader"
	"github.com/mikaijun/gqlgen-tasks/middleware"
	"github.com/mikaijun/gqlgen-tasks/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := db.ConnectGORM()
	// loaderの初期化
	ldrs := loader.NewLoaders(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// resolver.goで宣言した構造体にデータベースの値を受け渡し
		DB: db, // 追加
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Middleware(ldrs, db, srv))
	http.HandleFunc("/login", service.Login)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func loadEnv() {
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
