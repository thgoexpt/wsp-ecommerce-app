package dbmodel

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"time"
)

type SalesHistoryModel struct {
	ID             bson.ObjectId         `bson:"_id,omitempty"`
	Time           time.Time             `bson:"time"`
	User           bson.ObjectId         `bson:"user"`
	Meats          map[bson.ObjectId]int `bson:"meats"`
	Price          float64               `bson:"price"`
	TrackingNumber string                `bson:"tracking_number"`
}

func MakeSalesHistory(time time.Time, user User, meats map[Meat]int, price float64, trackingNum string) (SalesHistoryModel, error) {
	meatsID := map[bson.ObjectId]int{}

	for k, v := range meats {
		if k.ID == "" {
			return SalesHistoryModel{}, errors.New("no meat id: please get id from database")
		}

		meatsID[k.ID] = v
	}

	if user.ID == "" {
		return SalesHistoryModel{}, errors.New("no user id: please get id from database")
	}

	sh := SalesHistoryModel{
		Time:           time,
		User:           user.ID,
		Meats:          meatsID,
		Price:          price,
		TrackingNumber: trackingNum,
	}

	return sh, nil
}
