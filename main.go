package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	solidenv "github.com/guitarpawat/wsp-ecommerce/env"
	"github.com/guitarpawat/wsp-ecommerce/handler"
)

var env = solidenv.GetEnv()

func main() {

	r := mux.NewRouter()
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	r.PathPrefix("/static/").Handler(fs)

	r.HandleFunc("/image/{name}", handler.Images)

	r.Handle("/", handlePage(handler.Home))

	r.Handle("/about/", handlePage(handler.About))

	r.Handle("/cart/", handlePage(handler.Cart))

	r.Handle("/contact/", handlePage(handler.Contact))

	r.Handle("/product/", handlePage(handler.Product))
	r.Handle("/product/sort/type={meattype}&priceSort={price_sort}/", handlePage(handler.ProductSortType))
	r.Handle("/product/search/name={name}&startPrice={startPrice}&endPrice={endPrice}&priceSort={price_sort}/", handlePage(handler.ProductSearch))

	r.Handle("/product-detail/{meatId}/", handlePage(handler.ProductDetail))

	r.Handle("/profile/", handlePage(handler.Profile))

	r.Handle("/profile-edit/", handlePage(handler.ProfileEdit))

	r.Handle("/add-product/", handlePage(handler.AddProduct))

	r.Handle("/product-stock/", handlePage(handler.ProductStock))

	r.Handle("/sale-history/", handlePage(handler.SaleHistory))

	r.Handle("/owner/", handlePage(handler.Owner))

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

	r.Handle("/add_meat/", handlePage(handler.AddMeat))

	r.Handle("/regis_meat/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.RegisMeat),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.Home))).
		Methods("POST")

	r.Handle("/edit-profile/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.EditProfile),
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.Profile))).
		Methods("POST")

	r.Handle("/product/add_cart:{meatId}&quantity={quantity}/", middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(handler.AddCart),
		middleware.DoableFunc(handler.Cart))).
		Methods("GET")

	r.Handle("/sales_history/", handlePage(handler.SaleHistory))

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

	db.MockUser()
}

func handlePage(df middleware.DoableFunc) http.Handler {
	return middleware.MakeMiddleware(nil,
		middleware.DoableFunc(handler.CheckSession),
		middleware.DoableFunc(handler.BuildHeader),
		middleware.DoableFunc(df))
}

func forceSslHeroku(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-forwarded-proto") != "https" {
			sslUrl := "https://" + r.Host + r.RequestURI
			http.Redirect(w, r, sslUrl, http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
