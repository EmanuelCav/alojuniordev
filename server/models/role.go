package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoleModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty" validate:"required"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
