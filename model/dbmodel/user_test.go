package dbmodel

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestMakeUser(t *testing.T) {
	username := "solid"
	password := "ske14 ku76"
	name := "test person"
	email := "test@example.com"
	address := "Bangkok Thailand"
	usertype := TypeOwner

	user, _ := MakeUser(username, password, name, email, address, usertype)

	if username != user.Username {
		t.Errorf("expected username: %s, but get: %s", username, user.Username)
	}
	if name != user.Fullname {
		t.Errorf("expected full name: %s, but get: %s", name, user.Fullname)
	}
	if email != user.Email {
		t.Errorf("expected email: %s, but get: %s", email, user.Email)
	}
	if address != user.Address {
		t.Errorf("expected address: %s, but get: %s", address, user.Address)
	}
	if usertype != user.Type {
		t.Errorf("expected user type: %d, but get: %d", usertype, user.Type)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		t.Errorf("error while trying to compare password with hash: %s", err)
	}
}

func TestUser_VerifyHash(t *testing.T) {
	username := "solid"
	password := "ske14 ku76"
	name := "test person"
	email := "test@example.com"
	address := "Bangkok Thailand"
	usertype := TypeOwner

	user, _ := MakeUser(username, password, name, email, address, usertype)

	ok := user.VerifyHash(password)
	if !ok {
		t.Errorf("error while verify hash")
	}
}

func TestUser_IsSame(t *testing.T) {
	u1 := User{
		ID: "1",
		Username: "u1",
	}

	u2 := User{
		ID: "1",
		Username: "u2",
	}

	u3 := User{
		ID: "2",
		Username: "u1",
	}

	if !u1.IsSame(u1) {
		t.Errorf("Expected: u1 is same as u1")
	}

	if !u1.IsSame(u2) {
		t.Errorf("Expected: u1 is same as u2")
	}

	if !u2.IsSame(u1) {
		t.Errorf("Expected: u2 is same as u1")
	}

	if u1.IsSame(u3) {
		t.Errorf("Expected: u1 is not same as u3")
	}

	if u3.IsSame(u1) {
		t.Errorf("Expected: u3 is not same as u1")
	}
}