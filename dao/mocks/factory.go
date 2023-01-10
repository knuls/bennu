package mocks

import (
	"github.com/knuls/bennu/auth"
	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/organizations"
	"github.com/knuls/bennu/users"
)

type Factory struct {
}

func (f *Factory) GetUserDao() dao.Dao[users.User] {
	return &UserDao{}
}
func (f *Factory) GetOrganizationDao() dao.Dao[organizations.Organization] {
	return &OrganizationDao{}
}
func (f *Factory) GetTokenDao() dao.Dao[auth.Token] {
	return &TokenDao{}
}

type ErrFactory struct {
}

func (f *ErrFactory) GetUserDao() dao.Dao[users.User] {
	return &ErrUserDao{}
}
func (f *ErrFactory) GetOrganizationDao() dao.Dao[organizations.Organization] {
	return &ErrOrganizationDao{}
}
func (f *ErrFactory) GetTokenDao() dao.Dao[auth.Token] {
	return &ErrTockenDao{}
}
