package handler

import (
	"html/template"
	"net/http"

	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model"
)

var t = template.Must(template.ParseGlob("template/*"))

func Home(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "home.html", nil)
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

	user, err := model.MakeUser(r.PostFormValue("username"), r.PostFormValue("password"),
		r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("address"), model.TypeUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.RegisUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	Home(w, r)
}

func Login(w http.ResponseWriter, r *http.Request) {

}
