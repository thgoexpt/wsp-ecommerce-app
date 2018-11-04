package dbmodel

import (
	"github.com/globalsign/mgo/bson"
)

type Cart struct {
	ID     bson.ObjectId  `bson:"_id,omitempty"`
	UserID bson.ObjectId         `bson:"userId"`
	Meats  map[string]int `bson:"meats"`
}

func (c Cart) RemoveMeat(meat bson.ObjectId) {
	delete(c.Meats, meat.Hex())
}

func (c Cart) SetMeat(meat bson.ObjectId, quantity int) {
	c.Meats[meat.Hex()] = quantity
}

func InitialCart(userId bson.ObjectId) Cart {
	cart := Cart{
		UserID: userId,
		Meats:  make(map[string]int),
	}

	return cart
}
