package mocks

import (
	"context"
	"errors"
	"time"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MockUsers = []*users.User{
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		FirstName: "first",
		LastName:  "knuls",
		Password:  "super-secret",
		Verified:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now().Add(5 * time.Minute)),
		FirstName: "second",
		LastName:  "knuls",
		Password:  "super-secret",
		Verified:  true,
		CreatedAt: time.Now().Add(5 * time.Minute),
		UpdatedAt: time.Now().Add(5 * time.Minute),
	},
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now().Add(10 * time.Minute)),
		FirstName: "third",
		LastName:  "knusecols",
		Password:  "super-secret",
		Verified:  false,
		CreatedAt: time.Now().Add(10 * time.Minute),
		UpdatedAt: time.Now().Add(10 * time.Minute),
	},
}

type UserDao struct {
}

func (m *UserDao) Find(ctx context.Context, filter dao.Where) ([]*users.User, error) {
	return MockUsers, nil
}
func (m *UserDao) FindOne(ctx context.Context, filter dao.Where) (*users.User, error) {
	return MockUsers[0], nil
}
func (m *UserDao) Create(ctx context.Context, user *users.User) (string, error) {
	return "", nil
}
func (m *UserDao) Update(ctx context.Context, user *users.User) (*users.User, error) {
	return nil, nil
}

type ErrUserDao struct {
}

func (m *ErrUserDao) Find(ctx context.Context, filter dao.Where) ([]*users.User, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrUserDao) FindOne(ctx context.Context, filter dao.Where) (*users.User, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrUserDao) Create(ctx context.Context, user *users.User) (string, error) {
	return "", nil
}
func (m *ErrUserDao) Update(ctx context.Context, user *users.User) (*users.User, error) {
	return nil, nil
}
