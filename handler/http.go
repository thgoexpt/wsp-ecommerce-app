package handler

import (
	"fmt"
	"net/http"
)

func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Host
	path := r.URL.Path

	if host == "" {
		host = "127.0.0.1"
	}

	http.Redirect(w, r, fmt.Sprintf("https://%s:4433%s", host, path), http.StatusPermanentRedirect)
}
