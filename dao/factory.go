package dao

import (
	"github.com/knuls/bennu/auth"
	"github.com/knuls/bennu/organizations"
	"github.com/knuls/bennu/users"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type Factory interface {
	GetUserDao() Dao[users.User]
	GetOrganizationDao() Dao[organizations.Organization]
	GetTokenDao() Dao[auth.Token]
}

type DaoFactory struct {
	userDao         *UserDao
	organizationDao *OrganizationDao
	tokenDao        *TokenDao
}

func (f *DaoFactory) GetUserDao() Dao[users.User] {
	return f.userDao
}

func (f *DaoFactory) GetOrganizationDao() Dao[organizations.Organization] {
	return f.organizationDao
}

func (f *DaoFactory) GetTokenDao() Dao[auth.Token] {
	return f.tokenDao
}

func NewDaoFactory(db *mongo.Database, validator *validator.Validator) *DaoFactory {
	return &DaoFactory{
		userDao:         NewUserDao(db, validator),
		organizationDao: NewOrganizationDao(db, validator),
		tokenDao:        NewTokenDao(db, validator),
	}
}
