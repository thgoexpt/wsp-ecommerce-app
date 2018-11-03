package pagemodel

import "github.com/globalsign/mgo/bson"

type Menu struct {
	User     string
	UserID   bson.ObjectId
	UserType int
	Warning  string
	Success  string
}
