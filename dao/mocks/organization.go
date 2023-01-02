package mocks

import (
	"context"
	"errors"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
)

type OrganizationDao struct {
}

func (m *OrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*models.Organization, error) {
	orgs := []*models.Organization{
		models.NewOrganization(),
		models.NewOrganization(),
		models.NewOrganization(),
	}
	return orgs, nil
}
func (m *OrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*models.Organization, error) {
	return models.NewOrganization(), nil
}
func (m *OrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	return "", nil
}
func (m *OrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, nil
}

type ErrOrganizationDao struct {
}

func (m *ErrOrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*models.Organization, error) {
	return nil, errors.New("some error")
}
func (m *ErrOrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*models.Organization, error) {
	return nil, errors.New("some error")
}
func (m *ErrOrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	return "", errors.New("some error")
}
func (m *ErrOrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, nil
}
