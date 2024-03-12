package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShowModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Show      primitive.ObjectID `json:"image,omitempty" bson:"image,omitempty" mson:"collection=ImageModel"`
	Position  int                `json:"position" bson:"position"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
