package main

import (
	"github.com/guitarpawat/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guitarpawat/wsp-ecommerce/handler"
)

func main() {
	r := mux.NewRouter()
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	r.PathPrefix("/static/").Handler(fs)

	r.Handle("/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.Home))).
		Methods("GET")

	r.HandleFunc("/about/", handler.About).Methods("GET")
	r.HandleFunc("/cart/", handler.Cart).Methods("GET")
	r.HandleFunc("/contact/", handler.Contact).Methods("GET")
	r.HandleFunc("/product/", handler.Product).Methods("GET")
	r.HandleFunc("/product-detail/", handler.ProductDetail).Methods("GET")

	r.HandleFunc("/regis/", handler.Regis).Methods("POST")

	r.Handle("/login/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.Login),
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.Home))).
		Methods("POST")

	r.Handle("/logout/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.Logout),
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.Home)))

	log.Fatalln(http.ListenAndServe(":8000", r))
}
