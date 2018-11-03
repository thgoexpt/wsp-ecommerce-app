package dbmodel

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

const TimeFormat = "02/01/2006"

type Meat struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	Name           string        `bson:"name"`
	Type           string        `bson:"type"`
	Grade          string        `bson:"grade"`
	Description    string        `bson:"des"`
	Price          float64       `bson:"price"`
	Quantity       int           `bson:"quantity"`
	Expire         time.Time     `bson:"expire"`
	ImageExtension string        `bson:"imageExtension"`
}

func MakeMeat(name, meattype, grade, des string, price float64, quantity int, expire time.Time, imageExt string) (Meat, error) {
	meat := Meat{
		Name:           name,
		Type:           meattype,
		Grade:          grade,
		Description:    des,
		Price:          price,
		Quantity:       quantity,
		Expire:         expire,
		ImageExtension: imageExt,
	}

	return meat, nil
}
