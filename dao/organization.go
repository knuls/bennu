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

type OrganizationDao struct {
	validator     *validator.Validator
	organizations *mongo.Collection
}

func (d *OrganizationDao) Find(ctx context.Context, filter Where) ([]*models.Organization, error) {
	var orgs []*models.Organization
	cursor, err := d.organizations.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return orgs, nil
		}
		return nil, err
	}
	if err = cursor.All(ctx, &orgs); err != nil {
		return nil, err
	}
	return orgs, nil
}

func (d *OrganizationDao) FindOne(ctx context.Context, filter Where) (*models.Organization, error) {
	result := d.organizations.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("no org found")
		}
		return nil, err
	}
	var org *models.Organization
	if err = result.Decode(&org); err != nil {
		return nil, err
	}
	return org, nil
}

func (d *OrganizationDao) Create(ctx context.Context, org *models.Organization) (string, error) {
	exists, err := d.Find(ctx, Where{{Key: "name", Value: org.Name}})
	if err != nil {
		return "", err
	}
	if len(exists) > 0 {
		return "", errors.New("name exists")
	}
	now := time.Now()
	org.CreatedAt = now
	org.UpdatedAt = now
	if err := d.validator.ValidateStruct(org); err != nil {
		return "", err
	}
	result, err := d.organizations.InsertOne(ctx, org)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (d *OrganizationDao) Update(ctx context.Context, org *models.Organization) (*models.Organization, error) {
	return nil, errors.New("no impl")
}

func NewOrganizationDao(db *mongo.Database, validator *validator.Validator) *OrganizationDao {
	return &OrganizationDao{
		validator:     validator,
		organizations: db.Collection(organizationsCollectionName),
	}
}
