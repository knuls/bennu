package mocks

import (
	"context"
	"errors"

	"github.com/knuls/bennu/auth"
	"github.com/knuls/bennu/dao"
)

type TokenDao struct {
}

func (m *TokenDao) Find(ctx context.Context, filter dao.Where) ([]*auth.Token, error) {
	tokens := []*auth.Token{
		auth.NewToken(),
		auth.NewToken(),
		auth.NewToken(),
	}
	return tokens, nil
}
func (m *TokenDao) FindOne(ctx context.Context, filter dao.Where) (*auth.Token, error) {
	return auth.NewToken(), nil
}
func (m *TokenDao) Create(ctx context.Context, token *auth.Token) (string, error) {
	return "", nil
}
func (m *TokenDao) Update(ctx context.Context, token *auth.Token) (*auth.Token, error) {
	return nil, nil
}

type ErrTockenDao struct {
}

func (m *ErrTockenDao) Find(ctx context.Context, filter dao.Where) ([]*auth.Token, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrTockenDao) FindOne(ctx context.Context, filter dao.Where) (*auth.Token, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrTockenDao) Create(ctx context.Context, token *auth.Token) (string, error) {
	return "", errors.New("some mock error")
}
func (m *ErrTockenDao) Update(ctx context.Context, token *auth.Token) (*auth.Token, error) {
	return nil, nil
}
