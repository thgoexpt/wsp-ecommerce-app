package dbmodel

import (
	"testing"
)

func TestMakeMeat(t *testing.T) {
	name := "Kurobuta"
	meattype := "Pig"
	grade := "A"
	des := "Meat of black pig"
	price := 9999.00

	meat, _ := MakeMeat(name, meattype, grade, des, price)

	if name != meat.Name {
		t.Errorf("expected name: %s, but get: %s", name, meat.Name)
	}
	if meattype != meat.Type {
		t.Errorf("expected full type: %s, but get: %s", meattype, meat.Type)
	}
	if grade != meat.Grade {
		t.Errorf("expected grade: %s, but get: %s", grade, meat.Grade)
	}
	if des != meat.Description {
		t.Errorf("expected description: %s, but get: %s", des, meat.Description)
	}
	if price != meat.Price {
		t.Errorf("expected price: %f, but get: %f", price, meat.Price)
	}
}
