package dao

import (
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type Factory struct {
	userDao         *UserDao
	organizationDao *OrganizationDao
	tokenDao        *TokenDao
}

func (f *Factory) GetUserDao() *UserDao {
	return f.userDao
}

func (f *Factory) GetOrganizationDao() *OrganizationDao {
	return f.organizationDao
}

func (f *Factory) GetTokenDao() *TokenDao {
	return f.tokenDao
}

func NewFactory(db *mongo.Database, validator *validator.Validator) *Factory {
	return &Factory{
		userDao:         NewUserDao(db.Collection("users"), validator),
		organizationDao: NewOrganizationDao(db.Collection("organizations"), validator),
		tokenDao:        NewTokenDao(db.Collection("tokens"), validator),
	}
}
