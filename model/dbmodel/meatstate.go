package dbmodel

import (
	"github.com/globalsign/mgo/bson"
)

type MeatState struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Meat  bson.ObjectId `bson:"meat"`
	Views int           `bson:"views"`
	Sold  int           `bson:"sold"`
}

func CreateMeatState(meat bson.ObjectId) MeatState {
	return MeatState{
		Meat:  meat,
		Views: 0,
		Sold:  0,
	}
}
