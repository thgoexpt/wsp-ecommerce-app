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

	r.Handle("/", handlePage(handler.Home))

	r.Handle("/about/", handlePage(handler.About))

	r.Handle("/cart/", handlePage(handler.Cart))

	r.Handle("/contact/", handlePage(handler.Contact))

	r.Handle("/product/", handlePage(handler.Product))

	r.Handle("/product-detail/", handlePage(handler.ProductDetail))

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

	httpr := mux.NewRouter()
	httpr.PathPrefix("/").HandlerFunc(handler.RedirectToHTTPS)

	go func() {
		log.Fatalln(http.ListenAndServeTLS(":4433","ssl/server.crt","ssl/server.key",r))
	}()
	log.Fatalln(http.ListenAndServe(":8000",httpr))
}

func handlePage(df middleware.DoableFunc) http.Handler {
	return middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(df))
}
