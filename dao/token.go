package dao

import (
	"context"

	"github.com/knuls/bennu/models"
)

type TokenDao struct {
}

func (d *TokenDao) Find(ctx context.Context, filter Where) ([]*models.Token, error) {
	return nil, nil
}

func (d *TokenDao) FindOne(ctx context.Context, filter Where) (*models.Token, error) {
	return nil, nil
}

func (d *TokenDao) Create(ctx context.Context, token *models.Token) (string, error) {
	return "", nil
}

func (d *TokenDao) Update(ctx context.Context, token *models.Token) (*models.Token, error) {
	return nil, nil
}

func NewTokenDao() *TokenDao {
	return &TokenDao{}
}
