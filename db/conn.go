package db

import "github.com/globalsign/mgo"

var url = "localhost:27017"

var db *mgo.Database = nil

func GetDB() (*mgo.Database, error) {

	if db == nil {
		session, err := mgo.Dial(url)
		if err != nil {
			return nil, err
		}

		return session.DB("solid"), nil
	}

	return db, nil

}
