package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectModel struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" validate:"required" mson:"cunique"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" validate:"required"`
	Image       primitive.ObjectID `json:"image,omitempty" bson:"image,omitempty" mson:"collection=ImageModel"`
	Repository  string             `json:"repository,omitempty" bson:"repository,omitempty"`
	Tabs        primitive.A        `json:"tabs,omitempty" bson:"tabs,omitempty" mson:"collection=TabModel"`
	Favorites   primitive.A        `json:"favorites,omitempty" bson:"favorites,omitempty" mson:"collection=UserModel"`
	Tools       primitive.A        `json:"tools,omitempty" bson:"tools,omitempty" mson:"collection=ToolModel"`
	CreatedAt   primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
