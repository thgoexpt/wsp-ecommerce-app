package handler

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectToHTTPS(t *testing.T) {
	var env string
	flag.StringVar(&env, "env", "TESTING", "Indicates the program runs on the TESTING or PRODUCTION environment")
	flag.Parse()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	RedirectToHTTPS(w, r)
	res := w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://example.com:4433" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://example.com:4433", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://example.com/", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://example.com:4433/" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://example.com:4433/", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://example.com/static/", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://example.com:4433/static/" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://example.com:4433/static/", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://example.com/static", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://example.com:4433/static" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://example.com:4433/static", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://www.example.com", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://www.example.com:4433" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://www.example.com:4433", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://www.example.com/", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://www.example.com:4433/" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://www.example.com:4433/", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://www.example.com/static/", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://www.example.com:4433/static/" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://www.example.com:4433/static/", res.Header.Get("Location"))
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://www.example.com/static", nil)
	RedirectToHTTPS(w, r)
	res = w.Result()
	if http.StatusPermanentRedirect != res.StatusCode {
		t.Errorf("expected status code: %d, but get: %d", http.StatusPermanentRedirect, res.StatusCode)
	}
	if "https://www.example.com:4433/static" != res.Header.Get("Location") {
		t.Errorf("expected redirect location to: %s, but get: %s", "https://www.example.com:4433/static", res.Header.Get("Location"))
	}
}
