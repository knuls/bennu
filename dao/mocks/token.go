package mocks

import (
	"context"
	"errors"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
)

type TokenDao struct {
}

func (m *TokenDao) Find(ctx context.Context, filter dao.Where) ([]*models.Token, error) {
	tokens := []*models.Token{
		models.NewToken(),
		models.NewToken(),
		models.NewToken(),
	}
	return tokens, nil
}
func (m *TokenDao) FindOne(ctx context.Context, filter dao.Where) (*models.Token, error) {
	return models.NewToken(), nil
}
func (m *TokenDao) Create(ctx context.Context, token *models.Token) (string, error) {
	return "", nil
}
func (m *TokenDao) Update(ctx context.Context, token *models.Token) (*models.Token, error) {
	return nil, nil
}

type ErrTockenDao struct {
}

func (m *ErrTockenDao) Find(ctx context.Context, filter dao.Where) ([]*models.Token, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrTockenDao) FindOne(ctx context.Context, filter dao.Where) (*models.Token, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrTockenDao) Create(ctx context.Context, token *models.Token) (string, error) {
	return "", errors.New("some mock error")
}
func (m *ErrTockenDao) Update(ctx context.Context, token *models.Token) (*models.Token, error) {
	return nil, nil
}
