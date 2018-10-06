package db

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model"
)

func RegisUser(user model.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	count, err := db.C("Users").Find(model.User{Username: user.Username}).Count()
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Username already exists")
	}

	count, err = db.C("Users").Find(model.User{Email: user.Email}).Count()
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

func AuthenticateUser(username, password string) (model.User, error) {
	db, err := GetDB()
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}
	err = db.C("Users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return model.User{}, err
	}

	ok := user.VerifyHash(password)
	if !ok {
		return model.User{}, errors.New("Invalid username/password")
	}

	return user, nil
}
