package pagemodel

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

type Menu struct {
	User        string
	UserID      bson.ObjectId
	UserAddress string
	UserType    int
	Warning     string
	Success     string
}

func (m Menu) IsPermissable() bool {
	if m.UserType == dbmodel.TypeEmployee || m.UserType == dbmodel.TypeOwner {
		return true
	}
	return false
}
