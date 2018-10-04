package model

import "github.com/globalsign/mgo/bson"
import "golang.org/x/crypto/bcrypt"

const TypeUser = 0
const TypeEmployee = 4
const TypeOwner = 8

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Hash     string        `bson:"hash"`
	Fullname string        `bson:"fullname"`
	Email    string        `bson:"email"`
	Address  string        `bson:"address"`
	Type     int           `bson:"type"`
}

func (u User) VerifyHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return (err == nil)
}

func MakeUser(username, hash, fullname, email, address string, usertype int) (User, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(hash), bcrypt.DefaultCost)
	if err != nil {
		return User{}, nil
	}
	user := User{
		Username: username,
		Hash:     string(h),
		Fullname: fullname,
		Email:    email,
		Address:  address,
		Type:     usertype,
	}

	return user, nil
}
