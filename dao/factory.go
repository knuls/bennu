package dao

import (
	"github.com/knuls/bennu/auth"
	"github.com/knuls/bennu/organizations"
	"github.com/knuls/bennu/users"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type Factory struct {
	userDao         *users.Dao
	organizationDao *organizations.Dao
	tokenDao        *auth.Dao
}

func (f *Factory) GetUserDao() *users.Dao {
	return f.userDao
}

func (f *Factory) GetOrganizationDao() *organizations.Dao {
	return f.organizationDao
}

func (f *Factory) GetTokenDao() *auth.Dao {
	return f.tokenDao
}

func NewFactory(validator *validator.Validator, db *mongo.Database) *Factory {
	return &Factory{
		userDao:         users.NewDao(validator, db.Collection("users")),
		organizationDao: organizations.NewDao(validator, db.Collection("organizations")),
		tokenDao:        auth.NewDao(validator, db.Collection("tokens")),
	}
}
