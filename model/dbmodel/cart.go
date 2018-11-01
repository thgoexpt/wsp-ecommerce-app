package dbmodel

import (
	"github.com/globalsign/mgo/bson"
)

type Cart struct {
	ID     bson.ObjectId         `bson:"_id,omitempty"`
	UserID bson.ObjectId         `bson:"userId"`
	Meats  map[bson.ObjectId]int `bson:"meats"`
}

func (c Cart) RemoveMeat(meat bson.ObjectId) {
	delete(c.Meats, meat)
}

func (c Cart) SetMeat(meat bson.ObjectId, quantity int) {
	c.Meats[meat] = quantity
}

func InitialCart(userId bson.ObjectId) Cart {
	cart := Cart{
		UserID: userId,
		Meats:  make(map[bson.ObjectId]int),
	}

	return cart
}
