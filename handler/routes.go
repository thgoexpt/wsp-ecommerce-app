package handler

import (
	"html/template"
	"net/http"

	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model"
	"golang.org/x/crypto/bcrypt"
)

var t = template.Must(template.ParseGlob("template/*"))

func Hello(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users []model.User
	err = db.C("Users").Find(nil).All(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "hello.html", users)
}

func Regis(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(r.PostFormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := model.User{
		Username: r.PostFormValue("username"),
		Hash:     string(hash),
		Fullname: r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Phone:    r.PostFormValue("phone"),
		Address:  r.PostFormValue("address"),
		Type:     model.USER,
	}

	err = db.C("Users").Insert(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Hello(w, r)
}
