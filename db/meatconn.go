package db

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/env"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var TestTime, _ = time.Parse(dbmodel.TimeFormat, "15/04/2019")
var TestMeat, _ = dbmodel.MakeMeat("Kurobuta", "Pig", "C", "Black Pig's Meat", 300.0, 50, TestTime, ".jpg")

func init() {
	if env.GetEnv() != env.Production {
		MockMeat()
	}
}

func MockMeat() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	db.C("Users").Remove(bson.M{"name": TestMeat.Name})

	RegisMeat(TestMeat)
}

func RegisMeat(meat dbmodel.Meat) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	err = db.C("Meats").Insert(meat)
	if err != nil {
		return err
	}
	return nil
}

func GetMeatId(meat dbmodel.Meat) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	defer db.Session.Close()

	var result dbmodel.Meat
	err = db.C("Meats").Find(bson.M{"name": meat.Name, "type": meat.Type, "grade": meat.Grade}).One(&result)
	if err != nil {
		return "", err
	}

	return result.ID.Hex(), nil
}

func GetMeat(id string) (dbmodel.Meat, error) {
	db, err := GetDB()
	if err != nil {
		return dbmodel.Meat{}, err
	}
	defer db.Session.Close()

	var meat dbmodel.Meat
	hexId := bson.ObjectIdHex(id)

	err = db.C("Meats").Find(bson.M(bson.M{"_id": hexId})).One(&meat)
	if err != nil {
		return meat, err
	}
	return meat, nil
}

func GetAllMeats() ([]dbmodel.Meat, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Session.Close()

	var meats []dbmodel.Meat
	err = db.C("Meats").Find(nil).All(&meats)
	if err != nil {
		return nil, err
	}
	return meats, nil
}
