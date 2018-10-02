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
	r.HandleFunc("/", handler.Hello).Methods("GET")
	r.HandleFunc("/", handler.Regis).Methods("POST")
	http.ListenAndServe(":8000", r)
}
