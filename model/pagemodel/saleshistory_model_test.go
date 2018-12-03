package pagemodel

import (
	"fmt"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"testing"
)

func TestToSalesHistoryPageModel(t *testing.T) {
	sh := dbmodel.SalesHistory{
		ID: "1",
		Meats: []dbmodel.Meats{},
		Price: 56.78,
		TrackingNumber: "9",
	}
	m := Menu{UserID: "2"}
	shpm, _ := ToSalesHistoryPageModel([]dbmodel.SalesHistory{sh}, m)

	if m.UserID != shpm.UserID {
		t.Errorf("expected user id: %s, but get: %s", m.UserID, shpm.UserID)
	}

	if sh.ID.Hex() != shpm.SalesHistory[0].ID {
		t.Errorf("expected history id: %s, but get: %s", sh.ID.Hex(), shpm.SalesHistory[0].ID)
	}

	if len(shpm.SalesHistory[0].Meats) != 0 {
		t.Errorf("expected len meats: %d, but get: %d", 0, len(shpm.SalesHistory[0].Meats))
	}

	if fmt.Sprintf("%.2f", sh.Price) != shpm.SalesHistory[0].Price {
		t.Errorf("expected price: %s, but get: %s",
			fmt.Sprintf("%.2f", sh.Price), shpm.SalesHistory[0].Price)
	}

	if sh.TrackingNumber != shpm.SalesHistory[0].TrackingNumber {
		t.Errorf("expected tracking number: %s, but get: %s",
			sh.TrackingNumber, shpm.SalesHistory[0].TrackingNumber)
	}
}