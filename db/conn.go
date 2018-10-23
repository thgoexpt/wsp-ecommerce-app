package db

import (
	"github.com/globalsign/mgo"
)

var session *mgo.Session = nil

func GetDB() (*mgo.Database, error) {

	host := "localhost"
	port := "27017"
	db := "solid"

	if session == nil {
		var err error
		session, err = mgo.Dial(host + ":" + port)
		if err != nil {
			return nil, err
		}
	}

	return session.Copy().DB(db), nil

}
