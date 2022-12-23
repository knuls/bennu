package dao

import (
	"context"

	"github.com/knuls/bennu/models"
)

type UserDao struct {
}

func (d *UserDao) Find(ctx context.Context, filter Where) ([]*models.User, error) {
	return nil, nil
}

func (d *UserDao) FindOne(ctx context.Context, filter Where) (*models.User, error) {
	return nil, nil
}

func (d *UserDao) Create(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}

func (d *UserDao) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func NewUserDao() *UserDao {
	return &UserDao{}
}
