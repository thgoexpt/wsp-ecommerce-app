package handler

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseGlob("template/*"))

func Hello(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "hello.html", "Welcome to SOLID company")
}
