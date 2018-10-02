package model

import "github.com/globalsign/mgo/bson"

const USER = 0
const ADMIN = 9

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Hash     string        `bson:"hash"`
	Fullname string        `bson:"fullname"`
	Email    string        `bson:"email"`
	Phone    string        `bson:"phone"`
	Address  string        `bson:"address"`
	Type     int           `bson:"type"`
}
