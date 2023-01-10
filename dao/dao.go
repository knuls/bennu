package dao

import (
	"context"

	"github.com/knuls/bennu/models"
	"github.com/knuls/bennu/users"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	usersCollectionName         = "users"
	organizationsCollectionName = "organizations"
	tokensCollectionName        = "tokens"
)

type Where bson.D

type Dao[T users.User | models.Organization | models.Token] interface {
	finder[T]
	creator[T]
	updater[T]
}

type finder[T users.User | models.Organization | models.Token] interface {
	Find(ctx context.Context, filter Where) ([]*T, error)
	FindOne(ctx context.Context, filter Where) (*T, error)
}

type creator[T users.User | models.Organization | models.Token] interface {
	Create(ctx context.Context, t *T) (string, error)
}

type updater[T users.User | models.Organization | models.Token] interface {
	Update(ctx context.Context, t *T) (*T, error)
}
