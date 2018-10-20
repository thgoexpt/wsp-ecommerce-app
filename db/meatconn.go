package db

import (
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

func RegisMeat(meat dbmodel.Meat) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	// TODO
	// Check case

	err = db.C("Meats").Insert(meat)
	if err != nil {
		return err
	}

	return nil
}
