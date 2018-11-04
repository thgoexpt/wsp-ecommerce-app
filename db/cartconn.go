package db

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var NonCart = errors.New("Carts Not Exist")

var mockUserID1 = bson.ObjectIdHex("ba2946f27d9d403ce895633b")
var mockUserID2 = bson.ObjectIdHex("f8f0b5922a47fef34a30327b")

func MockCart() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	cart1 := dbmodel.InitialCart(mockUserID1)
	cart2 := dbmodel.InitialCart(mockUserID2)

	db.C("Carts").Remove(bson.M{"userID": cart1.UserID})
	db.C("Carts").Remove(bson.M{"userID": cart2.UserID})

	RegisCart(cart1)
	RegisCart(cart2)
}

func RegisCart(cart dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	// Check case

	_, err = db.C("Carts").Upsert(bson.M{"userID": cart.UserID}, cart)
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

	user, err := GetUserFromName(userName)
	if err != nil {
		return dbmodel.Cart{}, errors.New("Unable to find user")
	}
	return GetCartID(user.ID)
}

func GetCartID(id bson.ObjectId) (dbmodel.Cart, error) {
	db, err := GetDB()
	if err != nil {
		return dbmodel.Cart{}, err
	}
	defer db.Session.Close()

	var cart dbmodel.Cart
	err = db.C("Carts").Find(bson.M{"userID": id}).One(&cart)
	if err != nil {
		return dbmodel.Cart{}, NonCart
	}
	return cart, nil
}

func UpdateCart(userID, meat bson.ObjectId, quantity int) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	cartMeat := dbmodel.CartMeats{
		ID:       meat,
		Quantity: quantity,
	}

	//FIXME: First Time using this will not work
	err = db.C("Carts").Update(bson.M{
		"userID": userID,
		// "meats": bson.M{
		// 	"meat": meat,
		// },
	},
		bson.M{
			"$pull": bson.M{
				"meats": bson.M{
					"meat": meat,
				},
			},
		})
	if err != nil {
		return err
	}

	_, err = db.C("Carts").Upsert(
		bson.M{
			"userID": userID,
		},
		bson.M{
			// "$pull": bson.M{"meats": bson.M{"meat": meat}},
			"$push": bson.M{"meats": cartMeat},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
