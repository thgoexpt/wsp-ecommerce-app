package db

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var MockSalesHistory1 dbmodel.SalesHistory
var MockSalesHistory2 dbmodel.SalesHistory

func MockSalesHistory() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	db.C("SalesHistory").Remove(bson.M{"tracking_number": "EA123456788TH"})
	db.C("SalesHistory").Remove(bson.M{"tracking_number": "EA123456789TH"})
	userobj, err := AuthenticateUser("test", "test")
	if err != nil {
		panic(err)
	}

	meatobj1id, err := GetMeatId(TestMeat)
	if err != nil {
		panic(err)
	}
	meatobj1, err := GetMeat(meatobj1id)
	if err != nil {
		panic(err)
	}

	meatobj2id, err := GetMeatId(TestMeat2)
	if err != nil {
		panic(err)
	}
	meatobj2, err := GetMeat(meatobj2id)
	if err != nil {
		panic(err)
	}

	MockSalesHistory1, err = dbmodel.MakeSalesHistory(TestTime, userobj, map[dbmodel.Meat]int{meatobj1:3, meatobj2:1},12.50,"EA123456788TH")
	if err != nil {
		panic(err)
	}
	MockSalesHistory2, err = dbmodel.MakeSalesHistory(TestTime, userobj, map[dbmodel.Meat]int{meatobj2:5},12.50,"EA123456789TH")
	if err != nil {
		panic(err)
	}

	RegisSalesHistory(MockSalesHistory1)
	RegisSalesHistory(MockSalesHistory2)
}

func RegisSalesHistory(sh dbmodel.SalesHistory) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	err = db.C("SalesHistory").Insert(sh)
	if err != nil {
		return err
	}
	return nil
}

func GetUserSalesHistory(userID bson.ObjectId) ([]dbmodel.SalesHistory, error)  {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Session.Close()

	sales := new([]dbmodel.SalesHistory)

	err = db.C("SalesHistory").Find(bson.M(bson.M{"user": userID})).Sort("time").All(sales)
	if err != nil {
		return nil, err
	}
	return *sales, nil
}