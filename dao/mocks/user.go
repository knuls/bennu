package mocks

import (
	"context"
	"errors"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
)

type UserDao struct {
}

func (m *UserDao) Find(ctx context.Context, filter dao.Where) ([]*models.User, error) {
	users := []*models.User{
		models.NewUser(),
		models.NewUser(),
		models.NewUser(),
	}
	return users, nil
}
func (m *UserDao) FindOne(ctx context.Context, filter dao.Where) (*models.User, error) {
	return models.NewUser(), nil
}
func (m *UserDao) Create(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}
func (m *UserDao) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

type ErrUserDao struct {
}

func (m *ErrUserDao) Find(ctx context.Context, filter dao.Where) ([]*models.User, error) {
	return nil, errors.New("some error")
}
func (m *ErrUserDao) FindOne(ctx context.Context, filter dao.Where) (*models.User, error) {
	return nil, errors.New("some error")
}
func (m *ErrUserDao) Create(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}
func (m *ErrUserDao) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}
