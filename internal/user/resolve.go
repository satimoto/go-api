package user

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, id int64) (db.User, error)
	GetUserByDeviceToken(ctx context.Context, deviceToken string) (db.User, error)
	GetUserByLinkingPubkey(ctx context.Context, linkingPubkey string) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
}

type UserResolver struct {
	Repository UserRepository
}

func NewResolver(repositoryService *db.RepositoryService) *UserResolver {
	repo := UserRepository(repositoryService)
	return &UserResolver{repo}
}
