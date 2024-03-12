package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TabModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"tab,omitempty" bson:"tab,omitempty" validate:"required" mson:"cunique"`
	Codes     primitive.A        `json:"codes,omitempty" bson:"codes,omitempty" mson:"collection=CodeModel"`
	Notes     primitive.A        `json:"notes,omitempty" bson:"notes,omitempty" mson:"collection=NoteModel"`
	Shows     primitive.A        `json:"shows,omitempty" bson:"shows,omitempty" mson:"collection=ShowModel"`
	Texts     primitive.A        `json:"texts,omitempty" bson:"texts,omitempty" mson:"collection=TextModel"`
	Subtitles primitive.A        `json:"subtitles,omitempty" bson:"subtitles,omitempty" mson:"collection=SubtitleModel"`
	Commands  primitive.A        `json:"commands,omitempty" bson:"commands,omitempty" mson:"collection=CommandModel"`
	Lists     primitive.A        `json:"lists,omitempty" bson:"lists,omitempty" mson:"collection=ListModel"`
	CreatedAt primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
