package handler

import (
	"net/http"

	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
)

func Owner(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}
	v.Set("next", false)

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.Owner{
		Menu:     header,
		SoldMeat: []pagemodel.CartMeatModel{},
	}

	soldMeat, err := db.GetSoldMeats()
	if err != nil {
		v.Set("warning", "Owner: unable to get data >> "+err.Error())
		t.ExecuteTemplate(w, "owner.html", model)
		return
	}
	for _, meatState := range soldMeat {
		meat, err := db.GetMeat(meatState.Meat.Hex())
		if err != nil {
			v.Set("warning", "Owner: unable to get meat >> "+err.Error())
			t.ExecuteTemplate(w, "owner.html", model)
			return
		}
		soldMeatModel := GetCartMeatModel(meat, meatState.Sold)
		model.SoldMeat = append(model.SoldMeat, soldMeatModel)
	}

	t.ExecuteTemplate(w, "owner.html", model)
}
