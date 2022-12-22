package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Scope     string             `json:"scope" bson:"scope" validate:"required"`
	Token     string             `json:"token" bson:"token" validate:"required"`
	Active    bool               `json:"active" bson:"active" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId" validate:"required,oid"`
}
