package pagemodel

import (
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"testing"
)

func TestMenu_IsPermissable_NonUser(t *testing.T) {
	menu := Menu{}
	expected := false
	if expected != menu.IsPermissable() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsPermissable())
	}
}

func TestMenu_IsPermissable_User(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeUser}
	expected := false
	if expected != menu.IsPermissable() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsPermissable())
	}
}

func TestMenu_IsPermissable_Employee(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeEmployee}
	expected := true
	if expected != menu.IsPermissable() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsPermissable())
	}
}

func TestMenu_IsPermissable_Owner(t *testing.T) {
	menu := Menu{UserType:dbmodel.TypeOwner}
	expected := true
	if expected != menu.IsPermissable() {
		t.Errorf("expected: %t, but get: %t", expected, menu.IsPermissable())
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