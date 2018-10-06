package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guitarpawat/wsp-ecommerce/handler"
)

func main() {
	r := mux.NewRouter()
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	r.PathPrefix("/static/").Handler(fs)
	// r.HandleFunc("/", handler.Hello).Methods("GET")
	r.HandleFunc("/", handler.Home).Methods("GET")
	r.HandleFunc("/about/", handler.About).Methods("GET")
	r.HandleFunc("/cart/", handler.Cart).Methods("GET")
	r.HandleFunc("/contact/", handler.Contact).Methods("GET")
	r.HandleFunc("/product/", handler.Product).Methods("GET")
	r.HandleFunc("/product-detail/", handler.ProductDetail).Methods("GET")

	r.HandleFunc("/regis/", handler.Regis).Methods("POST")
	r.HandleFunc("/login/", handler.Login).Methods("POST")
	http.ListenAndServe(":8000", r)
}
