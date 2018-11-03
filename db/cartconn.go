package db

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

func RegisCart(cart dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	// Check case

	err = db.C("Carts").Insert(cart)
	if err != nil {
		return err
	}
	return nil
}

func SetCart(id bson.ObjectId, cart dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	err = db.C("Carts").Update(bson.M{"_id": id}, cart)
	if err != nil {
		return err
	}
	return nil
}
