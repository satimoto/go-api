package user

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	ListUsers(ctx context.Context) ([]db.User, error)
}
type UserResolver struct {
	Repository UserRepository
}

func NewResolver(repositoryService *db.RepositoryService) *UserResolver {
	repo := UserRepository(repositoryService)
	return &UserResolver{repo}
}
