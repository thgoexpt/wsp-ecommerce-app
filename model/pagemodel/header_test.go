package pagemodel

import (
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"testing"
)

func TestMenu_IsOwner_NonUser(t *testing.T) {
	menu := Menu{}
	expected := false
	if expected != menu.IsOwner() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsOwner())
	}
}

func TestMenu_IsOwner_User(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeUser}
	expected := false
	if expected != menu.IsOwner() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsOwner())
	}
}

func TestMenu_IsOwner_Employee(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeEmployee}
	expected := false
	if expected != menu.IsOwner() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsOwner())
	}
}

func TestMenu_IsOwner_Owner(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeOwner}
	expected := true
	if expected != menu.IsOwner() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsOwner())
	}
}

func TestMenu_IsAdmin_NonUser(t *testing.T) {
	menu := Menu{}
	expected := false
	if expected != menu.IsAdmin() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsAdmin())
	}
}

func TestMenu_IsAdmin_User(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeUser}
	expected := false
	if expected != menu.IsAdmin() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsAdmin())
	}
}

func TestMenu_IsAdmin_Employee(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeEmployee}
	expected := true
	if expected != menu.IsAdmin() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsAdmin())
	}
}

func TestMenu_IsAdmin_Owner(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeOwner}
	expected := true
	if expected != menu.IsAdmin() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsAdmin())
	}
}

func TestMenu_CountCart(t *testing.T) {
	menu := Menu{MeatInCartCart:[]CartMeatModel{CartMeatModel{},CartMeatModel{}}}
	expected := 2
	if expected != menu.CountCart() {
		t.Errorf("expected: %d, but get: %d", expected, menu.CountCart())
	}
}

func TestMenu_CountCart_Nil(t *testing.T) {
	menu := Menu{MeatInCartCart:nil}
	expected := 0
	if expected != menu.CountCart() {
		t.Errorf("expected: %d, but get: %d", expected, menu.CountCart())
	}
}