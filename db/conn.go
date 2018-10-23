package db

import (
	"github.com/globalsign/mgo"
	"github.com/guitarpawat/wsp-ecommerce/env"
	"strings"
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

	if env.GetEnv() == env.Production {
		db = strings.Split(db, "/")[2]
		return session.Copy().DB(db), nil
	}

	return session.Copy().DB(db), nil

}
