package repository

import (
	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"testing"
)

func TestNewMockUserRepository(t *testing.T) {
	ur := NewMockUserRepository()

	_, ok := ur.(*MockUserRepository)

	if !ok {
		t.Errorf("cannot cast UserRepository to MockUserRepository")
	}
}

func TestMockUserRepository_AddUser(t *testing.T) {
	mur := NewMockUserRepository()

	u1 := dbmodel.User{Username:"gp",Email:"gp@example.com"}
	u2 := dbmodel.User{Username:"gp2",Email:"gp@example.com2"}
	u3 := dbmodel.User{Username:"gp",Email:"gg@example.com"}
	u4 := dbmodel.User{Username:"gg",Email:"gp@example.com"}
	u5 := dbmodel.User{ID:bson.ObjectId("1")}

	err := mur.AddUser(u1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	err = mur.AddUser(u2)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	err = mur.AddUser(u3)
	if err == nil {
		t.Errorf("expected error")
	}

	err = mur.AddUser(u4)
	if err == nil {
		t.Errorf("expected error")
	}

	err = mur.AddUser(u5)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestMockUserRepository_GetUserByID(t *testing.T) {
	mur := NewMockUserRepository()
	u := dbmodel.User{ID:bson.ObjectId("99"), Username:"gp", Email:"gp@example.com"}

	_, err := mur.GetUserByID("99")
	if err == nil {
		t.Errorf("expected error")
	}

	mur.AddUser(u)
	u2, err := mur.GetUserByID("99")
	if err != nil {
		t.Errorf("unexpected error")
	}

	if u2.ID != u.ID {
		t.Errorf("expected: %s, but get: %s", u.ID, u2.ID)
	}

	if u2.Username != u.Username {
		t.Errorf("expected: %s, but get: %s", u.Username, u2.Username)
	}

	if u2.Email != u.Email {
		t.Errorf("expected: %s, but get: %s", u.Email, u2.Email)
	}
}

func TestMockUserRepository_GetUserByUsername(t *testing.T) {
	mur := NewMockUserRepository()
	u := dbmodel.User{ID:bson.ObjectId("99"), Username:"gp", Email:"gp@example.com"}

	_, err := mur.GetUserByUsername("gp")
	if err == nil {
		t.Errorf("expected error")
	}

	mur.AddUser(u)
	u2, err := mur.GetUserByUsername("gp")
	if err != nil {
		t.Errorf("unexpected error")
	}

	if u2.ID != u.ID {
		t.Errorf("expected: %s, but get: %s", u.ID, u2.ID)
	}

	if u2.Username != u.Username {
		t.Errorf("expected: %s, but get: %s", u.Username, u2.Username)
	}

	if u2.Email != u.Email {
		t.Errorf("expected: %s, but get: %s", u.Email, u2.Email)
	}
}

func TestMockUserRepository_CheckDuplicate(t *testing.T) {
	mur := NewMockUserRepository()

	u1 := dbmodel.User{Username:"gp",Email:"gp@example.com"}
	u2 := dbmodel.User{Username:"gp2",Email:"gp@example.com2"}
	u3 := dbmodel.User{Username:"gp",Email:"gg@example.com"}
	u4 := dbmodel.User{Username:"gg",Email:"gp@example.com"}
	u5 := dbmodel.User{ID:bson.ObjectId("1")}

	err := mur.CheckDuplicate(u1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	mur.AddUser(u1)

	err = mur.CheckDuplicate(u1)
	if err == nil {
		t.Errorf("expected error")
	}

	err = mur.AddUser(u2)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	mur.AddUser(u2)

	err = mur.CheckDuplicate(u2)
	if err == nil {
		t.Errorf("expected error")
	}

	err = mur.CheckDuplicate(u3)
	if err == nil {
		t.Errorf("expected error")
	}

	err = mur.CheckDuplicate(u4)
	if err == nil {
		t.Errorf("expected error")
	}

	err = mur.CheckDuplicate(u5)
	if err == nil {
		t.Errorf("expected error")
	}
}