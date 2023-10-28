package graph

//go:generate go run github.com/99designs/gqlgen generate
//今後schema.graphqlsを変更した際に go generate ./...で更新することができるようになる。

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}
