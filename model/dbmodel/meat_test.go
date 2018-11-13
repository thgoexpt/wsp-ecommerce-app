package dbmodel

import (
	"testing"
	"time"
)

func TestMakeMeat(t *testing.T) {
	name := "Kurobuta"
	meattype := "Pig"
	grade := "A"
	des := "Meat of black pig"
	price := 9999.00
	discount := 9.00
	quantity := 666
	expire, _ := time.Parse(time.RFC3339, "1998-01-20T06:30:15+07:00")
	ext := ".JPG"

	meat, _ := MakeMeat(name, meattype, grade, des, price, discount, quantity, expire, ext)

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
	if discount != meat.Discount {
		t.Errorf("expected discount: %f, but get: %f", discount, meat.Discount)
	}
	if quantity != meat.Quantity {
		t.Errorf("expected quantity: %d, but get: %d", quantity, meat.Quantity)
	}
	if expire != meat.Expire {
		t.Errorf("expected expire: %s, but get: %s", expire, meat.Expire)
	}

	if ext != meat.ImageExtension {
		t.Errorf("expected image extension: %s, but get: %s", ext, meat.ImageExtension)
	}
}
