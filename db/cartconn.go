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
	count, err := db.C("Carts").Find(bson.M{"userId": cart.UserID, "type": dbmodel.TypeUnpaid}).Count()
	if count != 0 {
		return AppendCart(cart)
	}

	err = db.C("Carts").Insert(cart)
	if err != nil {
		return err
	}
	return nil
}

func AppendCart(cart dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	var exCart dbmodel.Cart
	err = db.C("Carts").Find(bson.M{"userId": cart.UserID, "type": dbmodel.TypeUnpaid}).One(&exCart)
	if err != nil {
		return err
	}
	for meat, quantity := range cart.Meats {
		cart.Meats[meat] = quantity
	}
	err = db.C("Carts").Update(bson.M{"_id": exCart.ID}, bson.M{"$set": bson.M{"meats": cart.Meats}})
	if err != nil {
		return err
	}

	return nil
}
