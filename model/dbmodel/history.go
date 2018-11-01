package dbmodel

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type SalesHistoryModel struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Time time.Time `bson:"time"`
	Meats map[bson.ObjectId]int `bson:"meats"`
	Price float64 `bson:"price"`
	TrackingNumber string `bson:"tracking_number"`
}
