package dao

import (
	"context"
	"errors"

	"github.com/knuls/bennu/auth"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenDao struct {
	validator *validator.Validator
	tokens    *mongo.Collection
}

func (d *TokenDao) Find(ctx context.Context, filter Where) ([]*auth.Token, error) {
	return nil, errors.New("no impl")
}

func (d *TokenDao) FindOne(ctx context.Context, filter Where) (*auth.Token, error) {
	return nil, errors.New("no impl")
}

func (d *TokenDao) Create(ctx context.Context, token *auth.Token) (string, error) {
	return "", errors.New("no impl")
}

func (d *TokenDao) Update(ctx context.Context, token *auth.Token) (*auth.Token, error) {
	return nil, errors.New("no impl")
}

func NewTokenDao(db *mongo.Database, validator *validator.Validator) *TokenDao {
	return &TokenDao{
		validator: validator,
		tokens:    db.Collection(tokensCollectionName),
	}
}
