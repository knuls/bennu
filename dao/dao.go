package dao

import (
	"context"

	"github.com/knuls/bennu/auth"
	"github.com/knuls/bennu/organizations"
	"github.com/knuls/bennu/users"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	usersCollectionName         = "users"
	organizationsCollectionName = "organizations"
	tokensCollectionName        = "tokens"
)

type Where bson.D

type Dao[T users.User | organizations.Organization | auth.Token] interface {
	finder[T]
	creator[T]
	updater[T]
}

type finder[T users.User | organizations.Organization | auth.Token] interface {
	Find(ctx context.Context, filter Where) ([]*T, error)
	FindOne(ctx context.Context, filter Where) (*T, error)
}

type creator[T users.User | organizations.Organization | auth.Token] interface {
	Create(ctx context.Context, t *T) (string, error)
}

type updater[T users.User | organizations.Organization | auth.Token] interface {
	Update(ctx context.Context, t *T) (*T, error)
}
