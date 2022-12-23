package dao

import (
	"context"

	"github.com/knuls/bennu/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Where bson.D

type Finder[T models.User | models.Organization | models.Token] interface {
	Find(ctx context.Context, filter Where) ([]*T, error)
	FindOne(ctx context.Context, filter Where) (*T, error)
}

type Creater[T models.User | models.Organization | models.Token] interface {
	Create(ctx context.Context, t *T) (string, error)
}

type Updater[T models.User | models.Organization | models.Token] interface {
	Update(ctx context.Context, t *T) (*T, error)
}
