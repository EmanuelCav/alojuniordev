package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Category  string             `json:"category,omitempty" bson:"category,omitempty" mson:"cunique"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CategoryCreateModel struct {
	Category string `json:"category,omitempty" bson:"category,omitempty" mson:"cunique"`
}
