package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var t = template.Must(template.ParseGlob("template/*"))

func main() {
	r := mux.NewRouter()
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	r.PathPrefix("/static/").Handler(fs)
	r.HandleFunc("/", hello)
	http.ListenAndServe(":8000", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "hello.html", nil)
}
