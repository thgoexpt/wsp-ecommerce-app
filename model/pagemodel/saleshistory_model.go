package pagemodel

import (
	"fmt"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

type SalesHistoryPageModel struct {
	SalesHistory []SalesHistoryModel
	Menu
}

type SalesHistoryModel struct {
	ID             string
	Date           string
	Time           string
	Meats          map[dbmodel.Meat]int
	Price          float64
	TrackingNumber string
}

func ToSalesHistoryPageModel(sh []dbmodel.SalesHistory, menu Menu) (SalesHistoryPageModel, error) {
	page := SalesHistoryPageModel{Menu: menu}

	shms := new([]SalesHistoryModel)
	for _,v := range sh {
		shm, err := ToSalesHistoryModel(v)
		if err != nil {
			return SalesHistoryPageModel{}, err
		}

		*shms = append(*shms, shm)
	}

	page.SalesHistory = *shms

	return page, nil
}

func ToSalesHistoryModel(sh dbmodel.SalesHistory) (SalesHistoryModel, error) {
	shm := SalesHistoryModel{Meats: map[dbmodel.Meat]int{}}

	for _, v := range sh.Meats {
		meat, err := db.GetMeat(v.Meat.Hex())
		if err != nil {
			return SalesHistoryModel{}, err
		}
		shm.Meats[meat] = v.Quatity
	}

	shm.ID = sh.ID.Hex()

	year, month, day := sh.Time.Date()
	shm.Date = fmt.Sprintf("%02d/%02d/%04d", day, month, year)
	shm.Time = fmt.Sprintf("%02d:%02d", sh.Time.Hour(), sh.Time.Minute())

	shm.Price = sh.Price
	shm.TrackingNumber = sh.TrackingNumber

	return shm, nil
}