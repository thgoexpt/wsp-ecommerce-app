package db

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

func RegisSalesHistory(sh dbmodel.SalesHistoryModel) error {
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

func GetUserSalesHistory(userID string) ([]dbmodel.SalesHistoryModel, error)  {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Session.Close()

	sales := new([]dbmodel.SalesHistoryModel)
	hexId := bson.ObjectIdHex(userID)

	err = db.C("Meats").Find(bson.M(bson.M{"user": hexId})).Sort("time").All(sales)
	if err != nil {
		return nil, err
	}
	return *sales, nil
}