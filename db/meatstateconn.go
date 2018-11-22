package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

func RegisMeatState(meat bson.ObjectId) error {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	meatState := dbmodel.CreateMeatState(meat)
	err = db.C("MeatState").Insert(meatState)
	if err != nil {
		return err
	}

	return nil
}

func GetSoldMeats() ([]dbmodel.MeatState, error) {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	meatState := []dbmodel.MeatState{}
	err = db.C("MeatState").Find(bson.M{
		"sold": bson.M{
			"$gt": 0,
		},
	}).Sort("-sold").All(&meatState)
	if err != nil {
		return nil, err
	}
	return meatState, nil
}

func checkMeatStateExist(db *mgo.Database, meat bson.ObjectId) error {
	c, err := db.C("MeatState").Find(bson.M{"meat": meat}).Count()
	if err != nil {
		return err
	}
	if c <= 0 {
		err = RegisMeatState(meat)
		return err
	}
	return nil
}

func ViewMeat(meat bson.ObjectId) error {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	checkMeatStateExist(db, meat)

	err = db.C("MeatState").Update(bson.M{"meat": meat}, bson.M{
		"$inc": bson.M{
			"views": 1,
		},
	})
	return err
}

func SoldMeat(meat bson.ObjectId, sold int) error {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	checkMeatStateExist(db, meat)

	err = db.C("MeatState").Update(bson.M{"meat": meat}, bson.M{
		"$inc": bson.M{
			"sold": sold,
		},
	})
	return err
}
