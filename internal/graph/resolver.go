package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"test/internal/ports"
)

type Resolver struct {
	repo ports.TestRepo
}

func NewGraphQlResolver(repo ports.TestRepo) *Resolver {
	return &Resolver{
		repo: repo,
	}
}
