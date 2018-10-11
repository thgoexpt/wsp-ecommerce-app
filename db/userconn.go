package db

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

func RegisUser(user dbmodel.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	count, err := db.C("Users").Find(bson.M{"username": user.Username}).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Username already exists")
	}

	count, err = db.C("Users").Find(bson.M{"email": user.Email}).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Email already in use")
	}

	err = db.C("Users").Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(username, password string) (dbmodel.User, error) {
	db, err := GetDB()
	if err != nil {
		return dbmodel.User{}, err
	}
	defer db.Session.Close()

	user := dbmodel.User{}
	err = db.C("Users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return dbmodel.User{}, err
	}

	ok := user.VerifyHash(password)
	if !ok {
		return dbmodel.User{}, errors.New("Invalid username/password")
	}

	return user, nil
}