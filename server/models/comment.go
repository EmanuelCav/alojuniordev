package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Comment   string             `json:"comment,omitempty" bson:"comment,omitempty" validate:"required"`
	User      primitive.ObjectID `json:"user" bson:"user" mson:"collection=UserModel"`
	Project   primitive.ObjectID `json:"project" bson:"project" mson:"collection=ProjectModel"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
