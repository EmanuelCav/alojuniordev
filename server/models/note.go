package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NoteModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Note      string             `json:"note,omitempty" bson:"note,omitempty" validate:"required" mson:"cunique"`
	Position  int                `json:"position" bson:"position"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
