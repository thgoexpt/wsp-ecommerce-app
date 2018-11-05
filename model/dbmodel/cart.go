package dbmodel

import (
	"github.com/globalsign/mgo/bson"
)

type Cart struct {
	UserID bson.ObjectId `bson:"userID"`
	Meats  []CartMeats   `bson:"meats"`
}

type CartMeats struct {
	ID       bson.ObjectId `bson:"meat"`
	Quantity int           `bson:"quality"`
}

func (c *Cart) RemoveMeat(meat bson.ObjectId) {
	if len(c.Meats) == 0 {
		return
	}
	for i := 0; i < len(c.Meats); i++ {
		if c.Meats[i].ID == meat {
			c.Meats = append(c.Meats[:i], c.Meats[i+1:]...)
			return
		}
	}
}

func (c *Cart) GetQuantity(meatID bson.ObjectId) int {
	if len(c.Meats) == 0 {
		return 0
	}
	for i := 0; i < len(c.Meats); i++ {
		if c.Meats[i].ID == meatID {
			return c.Meats[i].Quantity
		}
	}
	return 0
}

func (c *Cart) SetMeat(meat bson.ObjectId, quantity int) {
	if len(c.Meats) == 0 {
		return
	}
	for i := 0; i < len(c.Meats); i++ {
		if c.Meats[i].ID == meat {
			c.Meats[i].Quantity = quantity
			return
		}
	}
	newMeat := CartMeats{
		ID:       meat,
		Quantity: quantity,
	}
	c.Meats = append(c.Meats, newMeat)
}

func InitialCart(userId bson.ObjectId) Cart {
	cart := Cart{
		UserID: userId,
		Meats:  []CartMeats{},
	}

	return cart
}
