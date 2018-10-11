package db

import (
	"github.com/globalsign/mgo"
	"github.com/guitarpawat/wsp-ecommerce/flagvalue"
)

var session *mgo.Session = nil

func GetDB() (*mgo.Database, error) {

	host := flagvalue.GetDBHost()
	port := flagvalue.GetDBPort()
	db := flagvalue.GetDBName()

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "27017"
	}

	if db == "" {
		db = "solid"
	}

	if session == nil {
		var err error
		session, err = mgo.Dial(host + ":" + port)
		if err != nil {
			return nil, err
		}
	}

	return session.Copy().DB(db), nil

}
