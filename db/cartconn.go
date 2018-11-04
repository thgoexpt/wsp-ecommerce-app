package db

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/env"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var NonCart = errors.New("Carts Not Exist")

func init() {
	if env.GetEnv() != env.Production {
		MockCart()
	}

	db, err := GetDB()
	if err != nil {
		return
	}
	defer db.Session.Close()

	var users []dbmodel.User
	err = db.C("User").Find(nil).Iter().All(&users)
	if err != nil {
		return
	}
	for i := 0; i < len(users); i++ {
		cart := dbmodel.InitialCart(users[i].ID)
		RegisCart(cart)
	}

}

func MockCart() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

}

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

func GetCart(userName string) (dbmodel.Cart, error) {
	db, err := GetDB()
	if err != nil {
		return dbmodel.Cart{}, err
	}
	defer db.Session.Close()

	user, err := GetUserFromName("userName")
	if err != nil {
		return dbmodel.Cart{}, err
	}
	var cart dbmodel.Cart
	err = db.C("Carts").Find(bson.M{"UserID": user.ID}).One(&cart)
	if err != nil {
		return dbmodel.Cart{}, NonCart
	}
	return cart, nil
}

func UpdateCart(c dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	err = db.C("Carts").Update(bson.M{"_id": c.ID}, c)
	if err != nil {
		return err
	}
	return nil
}
