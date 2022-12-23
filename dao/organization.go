package dao

import (
	"context"

	"github.com/knuls/bennu/models"
)

type OrganizationDao struct {
}

func (d *OrganizationDao) Find(ctx context.Context, filter Where) ([]*models.Organization, error) {
	return nil, nil
}

func (d *OrganizationDao) FindOne(ctx context.Context, filter Where) (*models.Organization, error) {
	return nil, nil
}

func (d *OrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	return "", nil
}

func (d *OrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, nil
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{}
}
