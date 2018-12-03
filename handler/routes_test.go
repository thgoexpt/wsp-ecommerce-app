package handler

import (
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	username := "sample"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	session, _ := s.Get(r, "user")
	session.Values["user"] = dbmodel.User{Username: username}
	session.Save(r, w)
	CheckSession(w, r, &v)
	Login(w, r, &v)

	expected := "You are already logged in"
	if v.Get("warning") != expected {
		t.Errorf("expected success message: %s, but get: %s", expected, v.Get("success"))
	}

	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %T, but get: %T", true, v.Get("next").(bool))
	}
}

func TestLogout(t *testing.T) {
	username := "sample"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	session, _ := s.Get(r, "user")
	session.Values["user"] = dbmodel.User{Username: username}
	session.Save(r, w)
	Logout(w, r, &v)
	model := dbmodel.User{}

	if session.Values["user"] != model {
		t.Errorf("expected blank user model")
	}
	expected := "Logout successful"
	if v.Get("success") != expected {
		t.Errorf("expected success message: %s, but get: %s", expected, v.Get("success"))
	}

	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %T, but get: %T", true, v.Get("next").(bool))
	}
}

func TestMock(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	Mock(w, r, &v)

	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %T, but get: %T", true, v.Get("next").(bool))
	}
}