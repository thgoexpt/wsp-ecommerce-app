package db

import (
	"github.com/globalsign/mgo"
	"github.com/guitarpawat/wsp-ecommerce/env"
)

var session *mgo.Session = nil

func GetDB() (*mgo.Database, error) {
	db := "solid"

	if session == nil {
		host := "localhost:27017"

		if env.GetEnv() == env.Production {
			host = env.GetMongoURI()
		}

		var err error
		session, err = mgo.Dial(host)
		if err != nil {
			return nil, err
		}
	}

	return session.Copy().DB(db), nil

}
