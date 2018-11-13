package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
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

func EditMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.MeatEdit{
		Menu:  header,
		State: pagemodel.EditMeatTxt,
	}

	vars := mux.Vars(r)
	meat, err := db.GetMeat(vars["meatID"])
	if err != nil {
		v.Set("warning", "Product: unable to get all meats >> "+err.Error())
		t.ExecuteTemplate(w, "add-product.html", model)
		return
	}
	model.MeatModel = GetMeatModel(meat)

	v.Set("next", false)
	t.ExecuteTemplate(w, "add-product.html", model)
}

func UpdateMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceStr := r.PostFormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	discountStr := r.PostFormValue("discount")
	var discount float64
	if discountStr == "" {
		discount = -1.0
	} else {
		discount, err = strconv.ParseFloat(discountStr, 64)
		if err != nil {
			v.Set("warning", "Price is not a number.")
			v.Set("next", true)
			return
		}
	}
	quantityStr := r.PostFormValue("quantity")
	quantity64, err := strconv.ParseInt(quantityStr, 10, 64)
	if err != nil {
		v.Set("warning", "Quantity is not an integer.")
		v.Set("next", true)
		return
	}
	quantity := int(quantity64)
	expireStr := r.PostFormValue("expire")
	expire, err := time.Parse("2/1/2006", expireStr)
	if err != nil {
		v.Set("warning", "Expire is invalid.")
		v.Set("next", true)
		return
	}
	id := r.PostFormValue("pro_id")
	file, h, err := r.FormFile("uploadfile")
	var ext string
	if err != nil {
		if err == http.ErrMissingFile {
			meat, err := db.GetMeat(id)
			if err != nil {
				v.Set("warning", err)
				v.Set("next", true)
				return
			}
			ext = meat.ImageExtension
		} else {
			v.Set("warning", err)
			v.Set("next", true)
			return
		}
	} else {
		ext = filepath.Ext(h.Filename)
		fname := "meat_" + id + ext
		err = db.EditFile(fname, file)
		if err != nil {
			v.Set("warning", "cannot upload image")
			v.Set("next", true)
			return
		}
	}

	meat, err := dbmodel.MakeMeat(r.PostFormValue("name"), r.PostFormValue("type"),
		r.PostFormValue("grade"), r.PostFormValue("des"), price, discount, quantity, expire, ext)
	if err != nil {
		v.Set("warning", err.Error())
	} else {
		err = db.UpdateMeat(bson.ObjectIdHex(id), meat)
		if err != nil {
			v.Set("warning", err.Error())
		} else {
			v.Set("success", "You have successfully edit meat in the store.")
		}
	}
	v.Set("next", true)
}
