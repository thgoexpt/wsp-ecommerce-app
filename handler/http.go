package handler

import (
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/env"
	"net"
	"net/http"
)

func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}
	port := ":443"
	if env.GetEnv() != "PRODUCTION" {
		port = ":4433"
	}
	target := "https://" + host + port + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}

	http.Redirect(w, r, target, http.StatusPermanentRedirect)
}

func RedirectToHTTPSMiddleware(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		host = r.Host
	}
	port := ":443"
	if env.GetEnv() != "PRODUCTION" {
		port = ":4433"
	}
	target := "https://" + host + port + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}

	http.Redirect(w, r, target, http.StatusPermanentRedirect)

	v.Set("next", true)
}