package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mikaijun/gqlgen-tasks/route"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	route.Route()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
