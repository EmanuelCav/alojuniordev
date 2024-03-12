package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToolModel struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tool             string             `json:"tool,omitempty" bson:"tool,omitempty" validate:"required" mson:"cunique"`
	ShortDesciprtion string             `json:"short_description,omitempty" bson:"short_description,omitempty" validate:"required"`
	Desciprtion      string             `json:"description,omitempty" bson:"description,omitempty" validate:"required"`
	Categoty         primitive.ObjectID `json:"category" bson:"category" mson:"collection=CategoryModel"`
	Image            primitive.ObjectID `json:"image" bson:"image" mson:"collection=ImageModel"`
	Tabs             primitive.A        `json:"tabs,omitempty" bson:"tabs,omitempty" mson:"collection=TabModel"`
	CreatedAt        primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt        primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ToolCreateModel struct {
	Tool             string `json:"tool,omitempty" bson:"tool,omitempty" validate:"required" mson:"cunique"`
	ShortDesciprtion string `json:"short_description,omitempty" bson:"short_description,omitempty" validate:"required"`
	Desciprtion      string `json:"description,omitempty" bson:"description,omitempty" validate:"required"`
	Categoty         string `json:"category,omitempty" bson:"category,omitempty" validate:"required"`
}
