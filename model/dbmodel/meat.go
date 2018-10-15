package dbmodel

import "github.com/globalsign/mgo/bson"

type Meat struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Type        string        `bson:"type"`
	Grade       string        `bson:"grade"`
	Description string        `bson:"des"`
	Price       float64       `bson:"price"`
}

func MakeMeat(name, meattype, grade, des string, price float64) (Meat, error) {
	meat := Meat{
		Name:        name,
		Type:        meattype,
		Grade:       grade,
		Description: des,
		Price:       price,
	}

	return meat, nil
}
