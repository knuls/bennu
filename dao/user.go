package dao

import (
	"context"
	"errors"
	"time"

	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao struct {
	users     *mongo.Collection
	validator *validator.Validator
}

func (d *UserDao) Find(ctx context.Context, filter Where) ([]*models.User, error) {
	var users []*models.User
	cursor, err := d.users.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return users, nil
		}
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserDao) FindOne(ctx context.Context, filter Where) (*models.User, error) {
	result := d.users.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("no user found")
		}
		return nil, err
	}
	var user *models.User
	if err = result.Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) Create(ctx context.Context, user *models.User) (string, error) {
	exists, err := d.Find(ctx, Where{{Key: "email", Value: user.Email}})
	if err != nil {
		return "", err
	}
	if len(exists) > 0 {
		return "", errors.New("email exists")
	}
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	now := time.Now()
	user.Verified = false
	user.CreatedAt = now
	user.UpdatedAt = now
	if err := d.validator.ValidateStruct(user); err != nil {
		return "", err
	}
	result, err := d.users.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *UserDao) Update(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, errors.New("no impl")
}

func NewUserDao(users *mongo.Collection, validator *validator.Validator) *UserDao {
	return &UserDao{
		users:     users,
		validator: validator,
	}
}
