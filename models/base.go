package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
}
