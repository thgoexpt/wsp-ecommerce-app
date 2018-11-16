package pagemodel

import (
	"fmt"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"log"
	"time"
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
	Price          string
	TrackingNumber string
}

func ToSalesHistoryPageModel(sh []dbmodel.SalesHistory, menu Menu) (SalesHistoryPageModel, error) {
	page := SalesHistoryPageModel{Menu: menu}

	shms := new([]SalesHistoryModel)
	for _, v := range sh {
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

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Println(err)
		return SalesHistoryModel{}, nil
	}
	year, month, day := sh.Time.In(loc).Date()
	shm.Date = fmt.Sprintf("%02d/%02d/%04d", day, month, year)
	shm.Time = fmt.Sprintf("%02d:%02d", sh.Time.In(loc).Hour(), sh.Time.In(loc).Minute())

	shm.Price = fmt.Sprintf("%.2f", sh.Price)
	shm.TrackingNumber = sh.TrackingNumber

	return shm, nil
}
