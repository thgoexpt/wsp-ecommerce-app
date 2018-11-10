package db

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var TestTime, _ = time.Parse(dbmodel.TimeFormat, "15/04/2019")
var TestTime2, _ = time.Parse(dbmodel.TimeFormat, "14/02/2019")
var ChickWingTime, _ = time.Parse(dbmodel.TimeFormat, "20/01/2019")

var TestMeat, _ = dbmodel.MakeMeat("Kurobuta", "Chicken", "C", "Black Pig's Meat", 300.1, 50, TestTime, ".jpg")
var CupidWing, _ = dbmodel.MakeMeat("Cupid's Wing", "Angle", "R", "Juicy wing meat of an angelic creature!", 400.0, 69, TestTime2, ".jpg")
var ChickWing, _ = dbmodel.MakeMeat("Chick's Wing", "Chicken", "D", "Chick's Meat", 500.0, 100, ChickWingTime, ".jpg")

const SortPrice = "price"
const SortPriceReverse = "-price"

func MockMeat() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	db.C("Meats").Remove(bson.M{"name": TestMeat.Name})
	db.C("Meats").Remove(bson.M{"name": CupidWing.Name})
	db.C("Meats").Remove(bson.M{"name": ChickWing.Name})

	RegisMeat(TestMeat)
	RegisMeat(CupidWing)
	RegisMeat(ChickWing)
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
	return GetMeatsPaging(0, 1)
}

func GetMeatsPaging(limit, page int) ([]dbmodel.Meat, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Session.Close()

	var meats []dbmodel.Meat
	if limit > 0 {
		err = db.C("Meats").Find(nil).Limit(limit).Skip((page - 1) * limit).Iter().All(&meats)
	} else {
		err = db.C("Meats").Find(nil).Iter().All(&meats)
	}
	if err != nil {
		return nil, err
	}
	return meats, nil
}

func SortType(meattype, sorting string) ([]dbmodel.Meat, error) {
	return SearchSort("", meattype, 0, -1, sorting, 10, 1)
}

func Search(name string, startPrice, endPrice float64, sorting string) ([]dbmodel.Meat, error) {
	return SearchSort(name, "", startPrice, endPrice, sorting, 10, 1)
}

func SearchSort(name, meattype string, startPrice, endPrice float64, sorting string, limit, page int) ([]dbmodel.Meat, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Session.Close()

	if name == "" || name == "all" {
		name = "."
	}
	if meattype == "" || meattype == "all" {
		meattype = "."
	}
	if startPrice < 0 {
		startPrice = 0
	}
	if sorting != SortPrice && sorting != SortPriceReverse {
		sorting = SortPrice
	}

	var meats []dbmodel.Meat
	var query *mgo.Query
	if endPrice != -1 {
		query = db.C("Meats").Find(bson.M{
			"name": bson.RegEx{
				Pattern: "(" + name + ")",
				Options: "i", //insensitive
			},
			"type": bson.RegEx{
				Pattern: "(" + meattype + ")",
				Options: "i", //insensitive
			},
			"price": bson.M{
				"$gte": startPrice,
				"$lte": endPrice,
			},
		})
	} else {
		query = db.C("Meats").Find(bson.M{
			"name": bson.RegEx{
				Pattern: "(" + name + ")",
				Options: "i", //insensitive
			},
			"type": bson.RegEx{
				Pattern: "(" + meattype + ")",
				Options: "i", //insensitive
			},
			"price": bson.M{
				"$gte": startPrice,
			},
		})
	}

	if limit > 0 {
		err = query.Limit(limit).Sort(sorting).Skip((page - 1) * limit).Iter().All(&meats)
	} else {
		err = query.Sort(sorting).Iter().All(&meats)
	}

	if err != nil {
		return nil, err
	}
	return meats, nil
}
