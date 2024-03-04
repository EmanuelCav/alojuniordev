package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToolModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tool      string             `json:"tool,omitempty" bson:"tool,omitempty" validate:"required" mson:"cunique"`
	Categoty  primitive.ObjectID `json:"category" bson:"category" mson:"collection=CategoryModel"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
