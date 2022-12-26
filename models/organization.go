package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	ID            primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string               `json:"name" bson:"name" validate:"required,alphanum"`
	Slug          string               `json:"slug" bson:"slug" validate:"required"`
	Profile       organizationProfile  `json:"profile" bson:"profile"`
	UserID        primitive.ObjectID   `json:"userId" bson:"userId" validate:"required,oid"`
	Collaborators []primitive.ObjectID `json:"collaborators,omitempty" bson:"collaborators,omitempty" validate:"dive,oid"`
	CreatedAt     time.Time            `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt     time.Time            `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

type organizationProfile struct {
	Website string `json:"website" bson:"website" validate:"url"`
	Address string `json:"address" bson:"address"`
	Phone   string `json:"phone" bson:"phone" validate:"e164"`
}

func (m *Organization) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (m *Organization) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&m)
	if errors.Is(err, io.EOF) {
		return err
	}
	return err
}
