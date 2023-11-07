package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mikaijun/gqlgen-tasks/db"
	"github.com/mikaijun/gqlgen-tasks/graph"
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
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Middleware(db, srv))
	http.HandleFunc("/login", service.Login)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
