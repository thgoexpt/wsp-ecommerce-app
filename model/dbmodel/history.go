package dbmodel

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"time"
)

type SalesHistory struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	Time           time.Time     `bson:"time"`
	User           bson.ObjectId `bson:"user"`
	Meats          []Meats         `bson:"meats"`
	Price          float64       `bson:"price"`
	TrackingNumber string        `bson:"tracking_number"`
}

type Meats struct {
	Meat    bson.ObjectId `bson:"meat"`
	Quatity int           `bson:"quality"`
}

func MakeSalesHistory(time time.Time, user User, meats map[Meat]int, price float64, trackingNum string) (SalesHistory, error) {
	meatsID := []Meats{}

	for k, v := range meats {
		if k.ID == "" {
			return SalesHistory{}, errors.New("no meat id: please get id from database")
		}

		meatsID = append(meatsID, Meats{k.ID, v})
	}

	if user.ID == "" {
		return SalesHistory{}, errors.New("no user id: please get id from database")
	}

	sh := SalesHistory{
		Time:           time,
		User:           user.ID,
		Meats:          meatsID,
		Price:          price,
		TrackingNumber: trackingNum,
	}

	return sh, nil
}
