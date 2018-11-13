package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
)

func GetMeatModel(meat dbmodel.Meat) pagemodel.MeatModel {
	return pagemodel.MeatModel{
		ID:          meat.ID.Hex(),
		Pic:         "/image/meat_" + meat.ID.Hex() + meat.ImageExtension,
		ProName:     meat.Name,
		Type:        meat.Type,
		Grade:       meat.Grade,
		Description: meat.Description,
		Price:       meat.Price,
		Discount:    meat.Discount,
		Expire:      meat.Expire.Format(dbmodel.TimeFormat),
		Quantity:    meat.Quantity,
		Total:       meat.Price * float64(meat.Quantity),
	}
}

func MeatDetailEdit(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	html := "~~meat_edit~~.html"
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.MeatEdit{
		Menu: header,
	}

	vars := mux.Vars(r)
	meat, err := db.GetMeat(vars["meatId"])
	if err != nil {
		v.Set("warning", "MeatEdit: >> "+err.Error())
		t.ExecuteTemplate(w, html, model)
		return
	}
	model.MeatModel = GetMeatModel(meat)

	v.Set("next", false)
	t.ExecuteTemplate(w, html, model)
}

func EditMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//TODO: Complete edit meat 'post'

	v.Set("next", true)
}
