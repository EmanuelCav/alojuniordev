package models

import (
	"fmt"

	"github.com/rahul-sinha1908/go-mongoose/mongoose"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string             `json:"username,omitempty" validate:"required" bson:"username,omitempty" mson:"cunique"`
	Password   string             `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
	Email      string             `json:"email,omitempty" validate:"required" bson:"email,omitempty" mson:"cunique"`
	Role       primitive.ObjectID `json:"role" bson:"role" mson:"collection=RoleModel"`
	Photo      primitive.ObjectID `json:"photo" bson:"photo" mson:"collection=ImageModel"`
	Phone      string             `json:"phone" bson:"phone" mson:"cunique"`
	Reputation int                `json:"reputation,omitempty" bson:"reputation,omitempty"`
	Status     bool               `json:"status,omitempty" bson:"status,omitempty"`
	Social     []string           `json:"social" bson:"social"`
	CreatedAt  primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserRegisterModel struct {
	Username string `json:"username,omitempty" validate:"required" bson:"username,omitempty" mson:"cunique"`
	Password string `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
	Confirm  string `json:"confirm,omitempty" validate:"required" bson:"confirm,omitempty"`
	Email    string `json:"email,omitempty" validate:"required" bson:"email,omitempty" mson:"cunique"`
	Role     string `json:"role" bson:"role" mson:"collection=RoleModel"`
}

type UserLoginModel struct {
	Email    string `json:"email,omitempty" validate:"required" bson:"email,omitempty" mson:"cunique"`
	Password string `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
}

// 	Teams         primitive.A        `json:"teams" bson:"teams" mson:"collection=UserModel"`

// PopulateTest This will populate the test field and give to you
func (c *UserModel) PopulateTest() *UserModel {
	u := UserModel{}
	mongoose.PopulateObject(c, "Test", &u)
	return &u
}

// PopulateTeams This will populate the teams and give it to you
func (c *UserModel) PopulateTeams() *[]UserModel {
	teams := make([]UserModel, 0)
	err := mongoose.PopulateObjectArray(c, "Teams", &teams)
	if err != nil {
		fmt.Println("Error While Populate ", err)
	}
	return &teams
}
