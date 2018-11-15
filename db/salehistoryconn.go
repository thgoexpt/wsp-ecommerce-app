package db

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"math/rand"
	"time"
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

	meatobj2id, err := GetMeatId(CupidWing)
	if err != nil {
		panic(err)
	}
	meatobj2, err := GetMeat(meatobj2id)
	if err != nil {
		panic(err)
	}

	meatobj3id, err := GetMeatId(ChickWing)
	if err != nil {
		panic(err)
	}
	meatobj3, err := GetMeat(meatobj3id)
	if err != nil {
		panic(err)
	}

	MockSalesHistory1, err = dbmodel.MakeSalesHistory(TestTime, userobj, map[dbmodel.Meat]int{meatobj1: 3, meatobj2: 1}, 12.50, "EA123456788TH")
	if err != nil {
		panic(err)
	}
	MockSalesHistory2, err = dbmodel.MakeSalesHistory(TestTime, userobj, map[dbmodel.Meat]int{meatobj3: 5}, 15.00, "EA123456789TH")
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

func GetUserSalesHistory(userID bson.ObjectId) ([]dbmodel.SalesHistory, error) {
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

func CommitSalesHistory(c dbmodel.Cart) error {
	history, err := MakeHistory(c)
	if err != nil {
		return err
	}
	err = RegisSalesHistory(history)
	if err != nil {
		return err
	}
	return nil
}

func MakeHistory(c dbmodel.Cart) (dbmodel.SalesHistory, error) {
	history := dbmodel.SalesHistory{
		Time: time.Now(),
		User: c.UserID,
		TrackingNumber: fmt.Sprintf("EA%09d%TH", rand.Int63n(999999999)),
		Meats: []dbmodel.Meats{},
		Price: 0.00,
	}

	for _, v := range c.Meats {
		meatObj := dbmodel.Meats{
			Meat:v.ID,
			Quatity:v.Quantity,
		}
		history.Meats = append(history.Meats, meatObj)
		meat, err := GetMeat(string(v.ID))
		if err != nil {
			return dbmodel.SalesHistory{}, err
		}
		history.Price += meat.Price
	}

	return history, nil
}