package repository

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"strconv"
)

type UserRepository interface {
	AddUser(user dbmodel.User) error
	EditUser(user dbmodel.User) error
	GetUserByID(id string) (dbmodel.User, error)
	GetUserByUsername(username string) (dbmodel.User, error)
	CheckDuplicate(username string) error
}

type MockUserRepository struct {
	next int
	users []dbmodel.User
}

func(mur MockUserRepository) AddUser(user dbmodel.User) error {
	err := mur.CheckDuplicate(user)
	if err != nil {
		return errors.New("add user: "+err.Error())
	}

	if user.ID == "" {
		user.ID = bson.ObjectId(strconv.Itoa(mur.next))
		mur.next++
	}

	mur.users = append(mur.users, user)
	return nil
}

func (mur MockUserRepository) EditUser(user dbmodel.User) error {
	return errors.New("edit user: not implemented yet")
}

func (mur MockUserRepository) GetUserByID(id string) (dbmodel.User, error) {
	user := dbmodel.User{ID:bson.ObjectId(id)}
	for _,v := range mur.users {
		if v.IsSame(user) {
			return v, nil
		}
	}

	return dbmodel.User{}, errors.New("get user by id: not found")
}

func (mur MockUserRepository) GetUserByUsername(username string) (dbmodel.User, error) {
	for _,v := range mur.users {
		if v.Username == username {
			return v, nil
		}
	}

	return dbmodel.User{}, errors.New("get user by username: not found")
}

func (mur MockUserRepository) CheckDuplicate(user dbmodel.User) error {
	for _,v := range mur.users {
		if v.ID == user.ID {
			return errors.New("duplicated id")
		} else if v.Username == user.Username {
			return errors.New("username is in used")
		} else if v.Email == user.Email {
			return errors.New("email is in used")
		}
	}

	return nil
}

func MakeMockUserRepository() MockUserRepository {
	mur := MockUserRepository{
		next: 1,
		users: []dbmodel.User{},
	}

	u1, _ := dbmodel.MakeUser("test","test","Tester","test@example.com","Kasetsart Uni. TH", dbmodel.TypeUser)
	u2, _ := dbmodel.MakeUser("test2","test2","Tester2","test2@example.com","Kasetsart Uni. TH", dbmodel.TypeUser)

	mur.AddUser(u1)
	mur.AddUser(u2)

	return mur
}