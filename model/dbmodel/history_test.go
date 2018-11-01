package dbmodel

import (
	"testing"
	"time"
)

func TestMakeSalesHistory_NoMeat(t *testing.T) {
	expectedTime := time.Now()
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	sh, _ := MakeSalesHistory(expectedTime, map[Meat]int{}, expectedPrice, expectedTrackingNumber)

	if expectedTime != sh.Time {
		t.Errorf("expected time: %s, but get: %s", expectedTime, sh.Time)
	}

	for _ = range sh.Meats {
		t.Fatalf("not expected any meats to be here")
	}

	if expectedPrice != sh.Price {
		t.Errorf("expected price: %f, but get: %f", expectedPrice, sh.Price)
	}

	if expectedTrackingNumber != sh.TrackingNumber {
		t.Errorf("expected time: %s, but get: %s", expectedTrackingNumber, sh.TrackingNumber)
	}
}

func TestMakeSalesHistory_SomeMeat(t *testing.T) {
	mockMeat1 := Meat{ID: "1"}
	mockMeat2 := Meat{ID: "2"}
	expectedTime := time.Now()
	expectedMeat := map[Meat]int{mockMeat1: 3, mockMeat2: 1}
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	sh, _ := MakeSalesHistory(expectedTime, expectedMeat, expectedPrice, expectedTrackingNumber)

	if expectedTime != sh.Time {
		t.Errorf("expected time: %s, but get: %s", expectedTime, sh.Time)
	}

	val1, ok := sh.Meats[mockMeat1.ID]
	if !ok {
		t.Errorf("expected to have meat id %s", mockMeat1.ID)
	}

	if val1 != expectedMeat[mockMeat1] {
		t.Errorf("expected quantity of meat id %s: %d, but get: %d", mockMeat1.ID, expectedMeat[mockMeat1], val1)
	}

	val2, ok := sh.Meats[mockMeat2.ID]
	if !ok {
		t.Errorf("expected to have meat id %s", mockMeat2.ID)
	}

	if val2 != expectedMeat[mockMeat2] {
		t.Errorf("expected quantity of meat id %s: %d, but get: %d", mockMeat2.ID, expectedMeat[mockMeat2], val2)
	}

	if expectedPrice != sh.Price {
		t.Errorf("expected price: %f, but get: %f", expectedPrice, sh.Price)
	}

	if expectedTrackingNumber != sh.TrackingNumber {
		t.Errorf("expected time: %s, but get: %s", expectedTrackingNumber, sh.TrackingNumber)
	}
}

func TestMakeSalesHistory_ErrorNoID(t *testing.T) {
	mockMeat := Meat{}
	expectedTime := time.Now()
	expectedMeat := map[Meat]int{mockMeat: 3}
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	_, err := MakeSalesHistory(expectedTime, expectedMeat, expectedPrice, expectedTrackingNumber)

	if err == nil {
		t.Errorf("expected error")
	}
}
