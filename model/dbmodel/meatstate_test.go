package dbmodel

import "testing"

func TestCreateMeatState(t *testing.T) {
	ms := CreateMeatState("1")
	if "1" != ms.Meat {
		t.Errorf("expected meatid: %s, but get: %s", "1", ms.Meat)
	}
}
