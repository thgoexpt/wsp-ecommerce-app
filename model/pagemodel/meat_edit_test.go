package pagemodel

import "testing"

func TestMeatEdit_IsAddMeat_AddMeat(t *testing.T) {
	me := MeatEdit{State:AddMeatTxt}
	expected := true
	if expected != me.IsAddMeat() {
		t.Errorf("expected: %t, but get: %t", expected, me.IsAddMeat())
	}
}

func TestMeatEdit_IsAddMeat_EditMeat(t *testing.T) {
	me := MeatEdit{State:EditMeatTxt}
	expected := false
	if expected != me.IsAddMeat() {
		t.Errorf("expected: %t, but get: %t", expected, me.IsAddMeat())
	}
}