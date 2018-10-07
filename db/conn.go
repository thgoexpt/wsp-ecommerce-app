package db

import "github.com/globalsign/mgo"

var url = "localhost:27017"

var session *mgo.Session = nil

func GetDB() (*mgo.Database, error) {

	if session == nil {
		var err error
		session, err = mgo.Dial(url)
		if err != nil {
			return nil, err
		}
	}

	return session.Copy().DB("solid"), nil

}
