package dao

import (
	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type Factory interface {
	GetUserDao() Dao[models.User]
	GetOrganizationDao() Dao[models.Organization]
	GetTokenDao() Dao[models.Token]
}

type DaoFactory struct {
	userDao         *UserDao
	organizationDao *OrganizationDao
	tokenDao        *TokenDao
}

func (f *DaoFactory) GetUserDao() Dao[models.User] {
	return f.userDao
}

func (f *DaoFactory) GetOrganizationDao() Dao[models.Organization] {
	return f.organizationDao
}

func (f *DaoFactory) GetTokenDao() Dao[models.Token] {
	return f.tokenDao
}

func NewDaoFactory(db *mongo.Database, validator *validator.Validator) *DaoFactory {
	return &DaoFactory{
		userDao:         NewUserDao(db, validator),
		organizationDao: NewOrganizationDao(db, validator),
		tokenDao:        NewTokenDao(db, validator),
	}
}
