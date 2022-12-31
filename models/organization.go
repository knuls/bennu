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
	ID            primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string               `json:"name" bson:"name" validate:"required,alphanum"`
	UserID        primitive.ObjectID   `json:"userId" bson:"userId" validate:"required,oid"`
	Collaborators []primitive.ObjectID `json:"collaborators,omitempty" bson:"collaborators,omitempty" validate:"dive,oid"`
	CreatedAt     time.Time            `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt     time.Time            `json:"updatedAt" bson:"updatedAt" validate:"required"`
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

func NewOrganization() *Organization {
	return &Organization{}
}
