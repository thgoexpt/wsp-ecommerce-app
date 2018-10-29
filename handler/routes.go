package handler

import (
	"encoding/gob"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
)

var s = sessions.NewCookieStore([]byte("NOT FOR PRODUCTION"))

var defaultHeader = pagemodel.Menu{
	Warning: "Something went wrong {defualt}",
}

func init() {
	gob.Register(dbmodel.User{})
	gob.Register(dbmodel.Meat{})
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

func BuildHeader(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header := pagemodel.Menu{}

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		header.User = ""
	} else {
		header.User = user.Username
	}

	warning, ok := v.Get("warning").(string)
	if !ok {
		header.Warning = ""
	} else {
		header.Warning = warning
	}

	success, ok := v.Get("success").(string)
	if !ok {
		header.Success = ""
	} else {
		header.Success = success
	}

	v.Set("header", header)
	v.Set("next", true)
}

func UserDetail(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.UserDetail{
		Menu: header,
	}

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		model.Fullname = ""
		model.Email = ""
		model.Address = ""
	} else {
		model.Fullname = user.Fullname
		model.Email = user.Email
		model.Address = user.Address
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "home.html", model)
}

func Home(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Home{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "home.html", model)
}

func About(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.About{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "about.html", model)
}

func Contact(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Contact{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "contact.html", model)
}

func Cart(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Card{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "cart.html", model)
}

func Product(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Product{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "product.html", model)
}

func ProductDetail(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "product-detail.html", model)
}

func ComingSoon(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "comingsoon.html", model)
}

func Regis(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := dbmodel.MakeUser(r.PostFormValue("username"), r.PostFormValue("password"),
		r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("address"), dbmodel.TypeUser)
	if err != nil {
		v.Set("warning", "Something went wrong")
	} else {
		err = db.RegisUser(user)
		if err != nil {
			v.Set("warning", err.Error())
		} else {
			v.Set("success", "User created successful, please login.")
		}
	}
	v.Set("next", true)
}

func MeatTestPage(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	// db, err := db.GetDB()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// var meat []dbmodel.Meat
	// err = db.C("Users").Find(nil).All(&meats)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	v.Set("next", false)
	t.ExecuteTemplate(w, "add_meat_test.html", nil)
}

func RegisMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	priceStr := r.PostFormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		v.Set("warning", "Price is not a number.")
		v.Set("next", true)
		return
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
	dayStr, monthStr, yearStr := str.Split(expireStr, "//")

	expireToParse := yearStr + "-" + monthStr + "-" + dayStr + "T00:00:00+07:00"

	// GMTPlus7 := int((7 * time.Hour).Seconds())
	// bangkok := time.FixedZone("Bangkok Time", GMTPlus7)
	expire, err := time.Parse(time.RFC3339, expireToParse)
	if err != nil {
		v.Set("warning", "Expire is invalid.")
		v.Set("next", true)
		return
	}

	meat, err := dbmodel.MakeMeat(r.PostFormValue("name"), r.PostFormValue("type"),
		r.PostFormValue("grade"), r.PostFormValue("des"), price, quantity, exprie)
	if err != nil {
		v.Set("warning", err.Error())
	} else {
		err = db.RegisMeat(meat)
		if err != nil {
			v.Set("warning", err.Error())
		} else {
			v.Set("success", "You have successfully add meat into the store.")
		}
	}
	v.Set("next", true)
}

func Login(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model := dbmodel.User{}
	if user != model {
		v.Set("warning", "You are already logged in")
	} else {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := db.AuthenticateUser(r.PostFormValue("username"), r.PostFormValue("password"))
		if err != nil {
			v.Set("warning", "Invalid username/password")
		} else {
			session, err := s.Get(r, "user")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(3)
				return
			}

			if r.PostFormValue("remember") != "true" {
				session.Options.MaxAge = 0
			}
			session.Values["user"] = user
			session.Save(r, w)
			v.Set("success", "Login successful")
		}
	}

	v.Set("next", true)
}

func Logout(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	session, err := s.Get(r, "user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Values["user"] = dbmodel.User{}
	session.Save(r, w)

	v.Set("success", "Logout successful")
	v.Set("next", true)
}
