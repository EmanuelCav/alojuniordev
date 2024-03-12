package models

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImageModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	ImageId   string             `json:"imageId,omitempty" bson:"imageId,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}
