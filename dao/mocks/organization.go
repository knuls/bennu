package mocks

import (
	"context"
	"errors"
	"time"

	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MockOrgs = []*models.Organization{
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		Name:      "Org1",
		UserID:    MockUsers[0].ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now().Add(5 * time.Minute)),
		Name:      "Org2",
		UserID:    MockUsers[0].ID,
		CreatedAt: time.Now().Add(5 * time.Minute),
		UpdatedAt: time.Now().Add(5 * time.Minute),
	},
	{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now().Add(10 * time.Minute)),
		Name:      "Org3",
		UserID:    MockUsers[0].ID,
		CreatedAt: time.Now().Add(10 * time.Minute),
		UpdatedAt: time.Now().Add(10 * time.Minute),
	},
}

type OrganizationDao struct {
}

func (m *OrganizationDao) Find(ctx context.Context, filter dao.Where) ([]*models.Organization, error) {
	return MockOrgs, nil
}
func (m *OrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*models.Organization, error) {
	return MockOrgs[0], nil
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
	return nil, errors.New("some mock error")
}
func (m *ErrOrganizationDao) FindOne(ctx context.Context, filter dao.Where) (*models.Organization, error) {
	return nil, errors.New("some mock error")
}
func (m *ErrOrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	return "", errors.New("some mock error")
}
func (m *ErrOrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, nil
}
