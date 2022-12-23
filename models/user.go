package models

import (
	"net/http"
	"time"
)

type User struct {
	BaseModel
	Email       string    `json:"email" bson:"email" validate:"required,email"`
	FirstName   string    `json:"firstName" bson:"firstName" validate:"required"`
	LastName    string    `json:"lastName" bson:"lastName" validate:"required"`
	Password    string    `json:"password" bson:"password" validate:"required"`
	Verified    bool      `json:"verified" bson:"verified" validate:"required"`
	LastLoginAt time.Time `json:"lastLoginAt,omitempty" bson:"lastLoginAt,omitempty"`
}

func (m *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUser() *User {
	return &User{}
}
