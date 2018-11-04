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

func (c Cart) RemoveMeat(meat bson.ObjectId) {
	for i := 0; i < len(c.Meats); i++ {
		if c.Meats[i].ID == meat {
			c.Meats = append(c.Meats[:i], c.Meats[i+1:]...)
			return
		}
	}
}

func (c Cart) SetMeat(meat bson.ObjectId, quantity int) {
	for _, cartMeat := range c.Meats {
		if cartMeat.ID == meat {
			cartMeat.Quantity = quantity
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
