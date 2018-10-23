package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	solidenv "github.com/guitarpawat/wsp-ecommerce/env"
	"github.com/guitarpawat/wsp-ecommerce/handler"
	"log"
	"net/http"
)

var env = solidenv.GetEnv()

func main() {

	r := mux.NewRouter()
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	r.PathPrefix("/static/").Handler(fs)

	r.Handle("/", handlePage(handler.Home))

	r.Handle("/about/", handlePage(handler.About))

	r.Handle("/cart/", handlePage(handler.ComingSoon))

	r.Handle("/contact/", handlePage(handler.ComingSoon))

	r.Handle("/product/", handlePage(handler.Product))

	r.Handle("/product-detail/", handlePage(handler.ComingSoon))

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

	handler.Validate()

	if env == solidenv.Production {
		log.Fatalln(http.ListenAndServe(":"+solidenv.GetPort(), forceSslHeroku(r)))
	} else if env == solidenv.CI {
		r.Handle("/mock/", middleware.MakeMiddleware(nil,
			middleware.DoableFunc(handler.Mock),
			middleware.DoableFunc(handler.CheckSession),
			middleware.DoableFunc(handler.BuildHeader),
			middleware.DoableFunc(handler.Home)))

		fmt.Println("Running on port 8000")
		log.Fatalln(http.ListenAndServe(":8000", r))
	} else {
		r.Handle("/mock/", middleware.MakeMiddleware(nil,
			middleware.DoableFunc(handler.Mock),
			middleware.DoableFunc(handler.CheckSession),
			middleware.DoableFunc(handler.BuildHeader),
			middleware.DoableFunc(handler.Home)))

		fmt.Println("Running on port 8000 and 4433")
		go func() {
			log.Fatalln(http.ListenAndServeTLS(":4433", "ssl/server.crt", "ssl/server.key", r))
		}()
		log.Fatalln(http.ListenAndServe(":8000", httpr))
	}

	db.Mock()
}

func handlePage(df middleware.DoableFunc) http.Handler {
	return middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(df))
}

func forceSslHeroku(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil {
			sslUrl := "https://" + r.Host + r.RequestURI
			http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})

}