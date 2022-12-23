package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	BaseModel
	Scope  string             `json:"scope" bson:"scope" validate:"required"`
	Token  string             `json:"token" bson:"token" validate:"required"`
	Active bool               `json:"active" bson:"active" validate:"required"`
	UserID primitive.ObjectID `json:"userId" bson:"userId" validate:"required,oid"`
}

func NewToken() *Token {
	return &Token{}
}
