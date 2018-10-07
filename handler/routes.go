package handler

import (
	"encoding/gob"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
	"html/template"
	"log"
	"net/http"
)

var t = template.Must(template.ParseGlob("template/*"))

var s= sessions.NewCookieStore([]byte("NOT FOR PRODUCTION"))

func init() {
	gob.Register(dbmodel.User{})
}

func CheckSession(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	session, err := s.Get(r, "user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, ok := session.Values["user"].(dbmodel.User)
	if !ok {
		user = dbmodel.User{}
		session.Values["user"] = user
	}

	v.Set("user", user)
	v.Set("next", true)
}

func Home(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model := pagemodel.Home{
		Menu: pagemodel.Menu{
			User: user.Email,
		},
	}
	t.ExecuteTemplate(w, "home.html", model)

	v.Set("next", false)
}

func About(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "about.html", nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "contact.html", nil)
}

func Cart(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "cart.html", nil)
}

func Product(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "product.html", nil)
}

func ProductDetail(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "product-detail.html", nil)
}

func Regis(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := dbmodel.MakeUser(r.PostFormValue("username"), r.PostFormValue("password"),
		r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("address"), dbmodel.TypeUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.RegisUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	vm := middleware.ValueMap{}
	vm.Set("next", true)

	Home(w, r, &vm)
}

func Login(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Email != "" {
		v.Set("warning", "You are already logged in")
	} else {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := db.AuthenticateUser(r.PostFormValue("username"), r.PostFormValue("password"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Invalid username/password")
			return
		}

		session, err := s.Get(r, "user")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(3)
			return
		}

		session.Values["user"] = user
		session.Save(r,w)
	}

	v.Set("next", true)
}

func Logout(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap)  {
	session, err := s.Get(r, "user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Values["user"] = dbmodel.User{}
	session.Save(r,w)

	v.Set("next", true)
}
