package pagemodel

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

type Menu struct {
	User           string
	UserID         bson.ObjectId
	UserType       int
	UserAddress    string
	Warning        string
	Success        string
	MeatInCartCart []CartMeatModel
	CartTotal      float64
}

func (m Menu) IsAdmin() bool {
	return m.UserType == dbmodel.TypeEmployee || m.IsOwner()
}

func (m Menu) IsOwner() bool {
	return m.UserType == dbmodel.TypeOwner
}

func (m Menu) CountCart() int {
	return len(m.MeatInCartCart)
}
