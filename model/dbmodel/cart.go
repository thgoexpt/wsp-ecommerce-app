package dbmodel

import (
	"github.com/globalsign/mgo/bson"
)

// TypeUnpaid is cart that yet to be paid.
const TypeUnpaid = 0

// TypePaid is cart that have already been paid
// It should be unable to modify
const TypePaid = 17

type Cart struct {
	ID     bson.ObjectId         `bson:"_id,omitempty"`
	UserID bson.ObjectId         `bson:"userId"`
	Type   int                   `bson:"type"`
	Meats  map[bson.ObjectId]int `bson:"meats"`
}

func (c Cart) RemoveMeat(meat bson.ObjectId) {
	if c.Type != TypePaid {
		delete(c.Meats, meat)
	}
}

func (c Cart) SetMeat(meat bson.ObjectId, quantity int) {
	if c.Type != TypePaid {
		c.Meats[meat] = quantity
	}
}

func InitialCart(userId bson.ObjectId) Cart {
	cart := Cart{
		UserID: userId,
		Type:   TypeUnpaid,
		Meats:  make(map[bson.ObjectId]int),
	}

	return cart
}
