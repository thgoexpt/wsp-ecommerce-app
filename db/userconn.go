package db

import (
	"errors"

	"github.com/guitarpawat/wsp-ecommerce/env"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var TestUser, _ = dbmodel.MakeUser("test", "test", "Test User", "test@example.com", "Kasetsart, TH", dbmodel.TypeUser)
var TestRegis, _ = dbmodel.MakeUser("regis", "regis", "Regis User", "regis@example.com", "Kasetsart, TH", dbmodel.TypeUser)
var TestEmployee, _ = dbmodel.MakeUser("emp", "emp", "Happy Employee", "emp@example.com", "Kasetsart, TH", dbmodel.TypeEmployee)
var TestOwner, _ = dbmodel.MakeUser("owner", "owner", "Rich Owner", "owner@example.com", "Kasetsart, TH", dbmodel.TypeOwner)
var TestLoginFailUserName = "fail"

func init() {
	if env.GetEnv() != env.Production {
		Mock()
	}
}

func Mock() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	db.C("Users").Remove(bson.M{"username": TestUser.Username})
	db.C("Users").Remove(bson.M{"username": TestRegis.Username})
	db.C("Users").Remove(bson.M{"username": TestEmployee.Username})
	db.C("Users").Remove(bson.M{"username": TestOwner.Username})

	RegisUser(TestUser)
	RegisUser(TestEmployee)
	RegisUser(TestOwner)
}

func RegisUser(user dbmodel.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	if user.Username == TestRegis.Username && (user.Hash != TestRegis.Hash ||
		user.Fullname != TestRegis.Fullname || user.Email != TestRegis.Email ||
		user.Address != TestRegis.Address || user.Type != TestRegis.Type) {
		return errors.New("Username already exists")
	}

	if user.Username == TestLoginFailUserName {
		return errors.New("Username already exists")
	}

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

func UpdateUser(id bson.ObjectId, fullname, email, address string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	err = db.C("Users").Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"fullname": fullname, "email": email, "address": address}})
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

	if username == TestLoginFailUserName {
		return dbmodel.User{}, errors.New("Invalid username/password")
	}

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
