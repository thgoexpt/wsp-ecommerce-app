package dbmodel

import (
	"github.com/globalsign/mgo/bson"
)

type MeatState struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Meats bson.ObjectId `bson:"meat"`
	Views int           `bson:"views"`
	Solds int           `bson:"solds"`
}

func CreateMeatState(meat bson.ObjectId) MeatState {
	return MeatState{
		Meats: meat,
		Views: 0,
		Solds: 0,
	}
}
