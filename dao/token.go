package dao

import (
	"context"
	"errors"

	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenDao struct {
	tokens    *mongo.Collection
	validator *validator.Validator
}

func (d *TokenDao) Find(ctx context.Context, filter Where) ([]*models.Token, error) {
	return nil, errors.New("no impl")
}

func (d *TokenDao) FindOne(ctx context.Context, filter Where) (*models.Token, error) {
	return nil, errors.New("no impl")
}

func (d *TokenDao) Create(ctx context.Context, token *models.Token) (string, error) {
	return "", errors.New("no impl")
}

func (d *TokenDao) Update(ctx context.Context, token *models.Token) (*models.Token, error) {
	return nil, errors.New("no impl")
}

func NewTokenDao(tokens *mongo.Collection, validator *validator.Validator) *TokenDao {
	return &TokenDao{
		tokens:    tokens,
		validator: validator,
	}
}
