package models

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	BaseModel
	Name          string               `json:"name" bson:"name" validate:"required,alphanum"`
	Slug          string               `json:"slug" bson:"slug" validate:"required"`
	Profile       OrganizationProfile  `json:"profile" bson:"profile"`
	UserID        primitive.ObjectID   `json:"userId" bson:"userId" validate:"required,oid"`
	Collaborators []primitive.ObjectID `json:"collaborators,omitempty" bson:"collaborators,omitempty" validate:"dive,oid"`
}

type OrganizationProfile struct {
	Website string `json:"website" bson:"website" validate:"url"`
	Address string `json:"address" bson:"address"`
	Phone   string `json:"phone" bson:"phone" validate:"e164"`
}

func (m *Organization) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewOrganization() *Organization {
	return &Organization{}
}
