package db

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var mockUserID = bson.ObjectIdHex("c8eccf23ec86cef2dc20b97e")
var mockMeatID = bson.ObjectIdHex("73cad0cd912f21387873132b")
var mockMeatID2 = bson.ObjectIdHex("d7fffd5cea0895fdb1e19208")

func RemoveMockCart() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	db.C("Carts").Remove(bson.M{"userID": mockUserID})
}

func TestRegisCart(t *testing.T) {
	cart := dbmodel.InitialCart(mockUserID)
	meatQty := 3

	cart.SetMeat(mockMeatID, meatQty)
	meatVal := cart.GetQuantity(mockMeatID)
	if meatQty != meatVal {
		t.Errorf("expected meat qty: %d, but get: %d", meatQty, meatVal)
	}

	RegisCart(cart)
	dbCart, _ := GetCartID(mockUserID)
	if cart.UserID != dbCart.UserID {
		t.Errorf("expected cart: %s, but get: %s", cart.UserID.Hex(), dbCart.UserID.Hex())
	}

	dbVal := dbCart.GetQuantity(mockMeatID)
	if meatQty != dbVal {
		t.Errorf("expected meat qty: %d, but get: %d", meatQty, dbVal)
	}

	RemoveMockCart()
}

func TestUpdateCart(t *testing.T) {
	// cart := dbmodel.InitialCart(mockUserID)
	qty := 5
	qty2 := 6
	qty3 := 7

	UpdateCart(mockUserID, mockMeatID, qty)
	dbCart, _ := GetCartID(mockUserID)
	dbVal := dbCart.GetQuantity(mockMeatID)
	if qty != dbVal {
		t.Errorf("expected meat qty: %d, but get: %d", qty, dbVal)
	}

	UpdateCart(mockUserID, mockMeatID, qty2)
	dbCart, _ = GetCartID(mockUserID)
	dbVal = dbCart.GetQuantity(mockMeatID)
	if qty2 != dbVal {
		t.Errorf("expected meat qty2: %d, but get: %d", qty2, dbVal)
	}

	UpdateCart(mockUserID, mockMeatID2, qty3)
	dbCart, _ = GetCartID(mockUserID)
	val1 := dbCart.GetQuantity(mockMeatID)
	if qty2 != val1 {
		t.Errorf("expected 2nd tried meat qty2: %d, but get: %d", qty2, val1)
	}
	val2 := dbCart.GetQuantity(mockMeatID2)
	if qty3 != val2 {
		t.Errorf("expected meat qty3: %d, but get: %d", qty3, val2)
	}
	numMeats := len(dbCart.Meats)
	if numMeats != 2 {
		t.Errorf("expected total of meats in cart: %d, but get: %d", 2, numMeats)
	}

	RemoveMockCart()
}

func TestUpdateAfterRegisCart(t *testing.T) {
	cart := dbmodel.InitialCart(mockUserID)
	meatQty := 4
	meatQty2 := 10

	cart.SetMeat(mockMeatID, meatQty)
	RegisCart(cart)

	UpdateCart(mockUserID, mockMeatID, meatQty2)

	dbCart, _ := GetCartID(mockUserID)
	dbVal := dbCart.GetQuantity(mockMeatID)
	if meatQty2 != dbVal {
		t.Errorf("expected meat qty: %d, but get: %d", meatQty2, dbVal)
	}

	RemoveMockCart()
}
