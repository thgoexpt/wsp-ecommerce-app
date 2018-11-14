package pagemodel

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

type Menu struct {
	User      string
	UserID    bson.ObjectId
	UserType  int
	Warning   string
	Success   string
	Cart      []CartMeatModel
	CartTotal float64
}

func (m Menu) IsPermissable() bool {
	if m.UserType == dbmodel.TypeEmployee || m.UserType == dbmodel.TypeOwner {
		return true
	}
	return false
}

func (m Menu) CountCart() int {
	return len(m.Cart)
}
