package handler

import (
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
	"net/http"
	"net/http/httptest"
	"testing"
)

func pageTest(df middleware.DoableFunc, userType int, t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	v := middleware.ValueMap{"header":pagemodel.Menu{UserType:userType}}
	func(df middleware.DoableFunc, w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
		defer func() { recover() }()
		df(w, r, v)
	}(df, w, r, &v)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected status code: %d, but get: %d", http.StatusOK, w.Result().StatusCode)
	}
	if v.Get("next").(bool) {
		t.Errorf("expected next to be: %T, but get: %T", false, v.Get("next").(bool))
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
		t.Errorf("expected next to be: %T, but get: %T", true, v.Get("next").(bool))
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
		t.Errorf("expected next to be: %T, but get: %T", true, v.Get("next").(bool))
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
		t.Errorf("expected next to be: %T, but get: %T", true, v.Get("next").(bool))
	}
}

func TestHome(t *testing.T) {
	pageTest(Home, 0, t)
}

func TestAbout(t *testing.T) {
	pageTest(About, 0, t)
}

func TestContact(t *testing.T) {
	pageTest(Contact, 0, t)
}

func TestCart(t *testing.T) {
	pageTest(Cart, 0, t)
}

func TestProduct(t *testing.T) {
	pageTest(Product, 0, t)
}

func TestProductDetail(t *testing.T) {
	pageTest(ProductDetail, 0, t)
}

func TestComingSoon(t *testing.T) {
	pageTest(ComingSoon, 0, t)
}

func TestProfile(t *testing.T) {
	pageTest(Profile, 0, t)
}

func TestProfileEdit(t *testing.T) {
	pageTest(ProfileEdit, 0, t)
}

func TestSale(t *testing.T) {
	pageTest(Sale, 0, t)
}

func TestAddProduct(t *testing.T) {
	pageTest(AddProduct, 0, t)
}

func TestCheckout(t *testing.T) {
	pageTest(Checkout, dbmodel.TypeUser, t)
}

func TestProductSortType(t *testing.T) {
	pageTest(ProductSortType, 0, t)
}

func TestProductSortTypePaging(t *testing.T) {
	pageTest(ProductSortTypePaging, 0, t)
}

func TestProductPaging(t *testing.T) {
	pageTest(ProductPaging, 0, t)
}

func TestProductSearchPaging(t *testing.T) {
	pageTest(ProductSearchPaging, 0, t)
}

func TestProductStock(t *testing.T) {
	pageTest(ProductStock, dbmodel.TypeEmployee, t)
}

func TestOwner(t *testing.T) {
	pageTest(Owner, dbmodel.TypeOwner, t)
}