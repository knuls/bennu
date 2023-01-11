package auth

import (
	"context"
	"errors"

	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type dao struct {
	validator *validator.Validator
	tokens    *mongo.Collection
}

func (d *dao) Find(ctx context.Context, filter bson.D) ([]*Token, error) {
	return nil, errors.New("no impl")
}

func (d *dao) FindOne(ctx context.Context, filter bson.D) (*Token, error) {
	return nil, errors.New("no impl")
}

func (d *dao) Create(ctx context.Context, token *Token) (string, error) {
	return "", errors.New("no impl")
}

func (d *dao) Update(ctx context.Context, token *Token) (*Token, error) {
	return nil, errors.New("no impl")
}

func NewDao(validator *validator.Validator, tokens *mongo.Collection) *dao {
	return &dao{
		validator: validator,
		tokens:    tokens,
	}
}
