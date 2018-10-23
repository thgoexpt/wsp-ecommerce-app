package handler

import (
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
	"net/http"
	"net/http/httptest"
	"testing"
)

func pageTest(df middleware.DoableFunc, t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	func(df middleware.DoableFunc, w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
		defer func() { recover() }()
		df(w, r, v)
	}(df, w, r, &v)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected status code: %d, but get: %d", http.StatusOK, w.Result().StatusCode)
	}
	if v.Get("next").(bool) {
		t.Errorf("expected next to be: %t, but get: %t", false, v.Get("next").(bool))
	}
}

func TestCheckSession(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	CheckSession(w, r, &v)
	session, err := s.Get(r, "user")
	if err != nil {
		t.Fatalf("failed to create new session")
	}

	_, ok := session.Values["user"].(dbmodel.User)
	if !ok {
		t.Fatalf("failed to create new blank user data")
	}

	_, ok = v.Get("user").(dbmodel.User)
	if !ok {
		t.Fatalf("failed to get user data")
	}

	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %t, but get: %t", true, v.Get("next").(bool))
	}
}

func TestBuildHeader(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	BuildHeader(w, r, &v)
	user := v.Get("header").(pagemodel.Menu).User
	if user != "" {
		t.Errorf("expected blank user, but get: %s", user)
	}
	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %t, but get: %t", true, v.Get("next").(bool))
	}

	username := "sample"
	success := "hi"
	warning := "bye"
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/", nil)
	v = middleware.ValueMap{"user": username, "success": success, "warning": warning}
	BuildHeader(w, r, &v)
	if v.Get("header").(pagemodel.Menu).Success != success {
		t.Errorf("expected success: %s, but get: %s", success, v.Get("header").(pagemodel.Menu).Success)
	}
	if v.Get("header").(pagemodel.Menu).Warning != warning {
		t.Errorf("expected warning: %s, but get: %s", warning, v.Get("header").(pagemodel.Menu).Warning)
	}
	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %t, but get: %t", true, v.Get("next").(bool))
	}
}

func TestHome(t *testing.T) {
	pageTest(Home, t)
}

func TestAbout(t *testing.T) {
	pageTest(About, t)
}

func TestContact(t *testing.T) {
	pageTest(Contact, t)
}

func TestCart(t *testing.T) {
	pageTest(Cart, t)
}

func TestProduct(t *testing.T) {
	pageTest(Product, t)
}

func TestProductDetail(t *testing.T) {
	pageTest(ProductDetail, t)
}

func TestComingSoon(t *testing.T) {
	pageTest(ComingSoon, t)
}

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
		t.Errorf("expected next to be: %t, but get: %t", true, v.Get("next").(bool))
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
		t.Errorf("expected next to be: %t, but get: %t", true, v.Get("next").(bool))
	}
}

func TestMock(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{}
	Mock(w, r, &v)

	if !v.Get("next").(bool) {
		t.Errorf("expected next to be: %t, but get: %t", true, v.Get("next").(bool))
	}
}