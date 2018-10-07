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
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.Home)))

	r.Handle("/about/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.About)))

	r.HandleFunc("/cart/", handler.Cart)
	r.HandleFunc("/contact/", handler.Contact)
	r.HandleFunc("/product/", handler.Product)
	r.HandleFunc("/product-detail/", handler.ProductDetail)

	r.Handle("/regis/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.Regis),
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.Home))).
		Methods("POST")

	r.Handle("/login/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.Login),
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.Home))).
		Methods("POST")

	r.Handle("/logout/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.Logout),
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.Home)))

	log.Fatalln(http.ListenAndServe(":8000", r))
}
