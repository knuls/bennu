package models

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	Verified  bool               `json:"verified" bson:"verified"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

func (m *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (m *User) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&m)
	if errors.Is(err, io.EOF) {
		return err
	}
	return err
}
