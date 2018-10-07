package handler

import (
	"net"
	"net/http"
)

func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	host,_, err :=net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}
	target := "https://" + host + ":4433" + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}

	http.Redirect(w, r, target, http.StatusPermanentRedirect)
}
