package dao

type DaoFactory struct {
	UserDao         *UserDao
	OrganizationDao *OrganizationDao
	TokenDao        *TokenDao
}

func (f *DaoFactory) GetUserDao() *UserDao {
	return f.UserDao
}

func (f *DaoFactory) GetOrganizationDao() *OrganizationDao {
	return f.OrganizationDao
}

func (f *DaoFactory) GetTokenDao() *TokenDao {
	return f.TokenDao
}

func NewDaoFactory() *DaoFactory {
	return &DaoFactory{
		UserDao:         NewUserDao(),
		OrganizationDao: NewOrganizationDao(),
		TokenDao:        NewTokenDao(),
	}
}
