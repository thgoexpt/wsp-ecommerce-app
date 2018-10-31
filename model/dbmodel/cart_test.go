package dbmodel

import (
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
)

func TestMakeCart(t *testing.T) {
	user := MakeExUser()
	meatA := MakeMeatA()
	aQuantity := 19
	meatB := MakeMeatB()
	bQuantity := 28

	cart := InitialCart(user.ID)

	if user.ID != cart.UserID {
		t.Errorf("expected id: %s, but get: %s", user.ID, cart.UserID)
	}

	i, _ := cart.Meats[meatA.ID]
	if i != 0 {
		t.Errorf("expected non exist meat quantity: %d, but get: %d", 0, i)
	}
	if meatA.ID == meatB.ID {
		t.Errorf("expected meatA and meatB to be different")
	}

	cart.SetMeat(meatA.ID, aQuantity)
	cart.SetMeat(meatB.ID, bQuantity)
	cartA, ok := cart.Meats[meatA.ID]
	if !ok {
		t.Errorf("expected meatA exist in cart.Meats: %t, but get: %t", true, ok)
	}
	if aQuantity != cartA {
		t.Errorf("expected meatA quantity: %d, but get: %d", aQuantity, cartA)
	}
	cartB, ok := cart.Meats[meatB.ID]
	if !ok {
		t.Errorf("expected meatB exist in cart.Meats: %t, but get: %t", true, ok)
	}
	if bQuantity != cartB {
		t.Errorf("expected meatB quantity: %d, but get: %d", bQuantity, cartB)
	}
}

func MakeExUser() User {
	user, _ := MakeUser("Alpha", "password", "Alpha Tester", "alpha@tester.com", "Digital Simulation", TypeUser)
	user.ID = bson.NewObjectId()
	return user
}

func MakeMeatA() Meat {
	name := "Scaled Fish"
	meattype := "Fish"
	grade := "D"
	des := "Fish that covered in scale"
	price := 846.5
	quantity := 5555
	expire, _ := time.Parse(time.RFC3339, "1998-01-20T06:30:15+07:00")
	ext := ".JPG"

	meat, _ := MakeMeat(name, meattype, grade, des, price, quantity, expire, ext)
	meat.ID = bson.NewObjectId()
	return meat
}

func MakeMeatB() Meat {
	name := "Rainbow Jelly"
	meattype := "Medusa"
	grade := "C"
	des := "Jellyfish jelly-like substance"
	price := 1354.6
	quantity := 999
	expire, _ := time.Parse("2/1/2006", "15/04/2018")
	ext := ".JPG"

	meat, _ := MakeMeat(name, meattype, grade, des, price, quantity, expire, ext)
	meat.ID = bson.NewObjectId()
	return meat
}
