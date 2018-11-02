package handler

import (
	"encoding/gob"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/gorilla/sessions"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
	"golang.org/x/crypto/bcrypt"
)

var s = sessions.NewCookieStore([]byte("NOT FOR PRODUCTION"))

var defaultHeader = pagemodel.Menu{
	Warning: "Something went wrong {defualt}",
}

func init() {
	gob.Register(dbmodel.User{})
	gob.Register(dbmodel.Meat{})
}

func CheckSession(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	session, err := s.Get(r, "user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, ok := session.Values["user"].(dbmodel.User)
	if !ok {
		user = dbmodel.User{}
		session.Values["user"] = user
	}

	v.Set("user", user)
	v.Set("next", true)
}

func BuildHeader(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header := pagemodel.Menu{}

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		header.User = ""
	} else {
		header.User = user.Username
		header.UserType = user.Type
	}

	warning, ok := v.Get("warning").(string)
	if !ok {
		header.Warning = ""
	} else {
		header.Warning = warning
	}

	success, ok := v.Get("success").(string)
	if !ok {
		header.Success = ""
	} else {
		header.Success = success
	}

	v.Set("header", header)
	v.Set("next", true)
}

func Home(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Home{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "home.html", model)
}

func About(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.About{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "about.html", model)
}

func Contact(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Contact{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "contact.html", model)
}

func Cart(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Card{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "cart.html", model)
}

func Product(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", false)
	model := pagemodel.Product{
		Menu:  header,
		Meats: []pagemodel.MeatModel{},
	}

	meats, err := db.GetAllMeats()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	t.ExecuteTemplate(w, "product_real.html", model)
}

func ProductSortType(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Product{
		Menu:  header,
		Meats: []pagemodel.MeatModel{},
	}

	vars := mux.Vars(r)

	meats, err := db.SortType(vars["type"], vars["price_sort"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "product_sort.html", model)
}

func ProductSearch(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", false)
	model := pagemodel.Product{
		Menu:  header,
		Meats: []pagemodel.MeatModel{},
	}

	vars := mux.Vars(r)

	startPrice, err := strconv.ParseFloat(vars["startPrice"], 64)
	if err != nil {
		v.Set("warning", "startPrice is not a number.")
		v.Set("next", true)
		return
	}
	endPrice, err := strconv.ParseFloat(vars["endPrice"], 64)
	if err != nil {
		v.Set("warning", "startPrice is not a number.")
		v.Set("next", true)
		return
	}

	meats, err := db.Search(vars["name"], startPrice, endPrice, vars["price_sort"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	t.ExecuteTemplate(w, "product_search.html", model)
}

func GetMeatModel(meat dbmodel.Meat) pagemodel.MeatModel {
	return pagemodel.MeatModel{
		ID:          meat.ID.Hex(),
		Pic:         "/image/meat_" + meat.ID.Hex() + meat.ImageExtension,
		ProName:     meat.Name,
		Type:        meat.Type,
		Grade:       meat.Grade,
		Description: meat.Description,
		Price:       meat.Price,
		Expire:      meat.Expire.Format(dbmodel.TimeFormat),
		Quantity:    meat.Quantity,
		Total:       meat.Price * float64(meat.Quantity),
	}
}

func ProductDetail(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", false)
	vars := mux.Vars(r)
	meat, err := db.GetMeat(vars["meatId"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	meatModel := GetMeatModel(meat)

	model := pagemodel.ProductDetail{
		Menu:      header,
		MeatModel: meatModel,
	}

	t.ExecuteTemplate(w, "product-detail.html", model)
}

func ComingSoon(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "comingsoon.html", model)
}

func AddProduct(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "add-product.html", model)
}

func ProductStock(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", false)
	model := pagemodel.Stock{
		Menu:  header,
		Meats: []pagemodel.MeatModel{},
	}

	meats, err := db.GetAllMeats()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	t.ExecuteTemplate(w, "product-stock.html", model)
}

func Regis(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := dbmodel.MakeUser(r.PostFormValue("username"), r.PostFormValue("password"),
		r.PostFormValue("name"), r.PostFormValue("email"), r.PostFormValue("address"), dbmodel.TypeUser)
	if err != nil {
		v.Set("warning", "Something went wrong")
	} else {
		err = db.RegisUser(user)
		if err != nil {
			v.Set("warning", err.Error())
		} else {
			v.Set("success", "User created successful, please login.")
		}
	}
	v.Set("next", true)
}

func AddMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}
	t.ExecuteTemplate(w, "add-product.html", model)
	v.Set("next", false)
}

func RegisMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	priceStr := r.PostFormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		v.Set("warning", "Price is not a number.")
		v.Set("next", true)
		return
	}
	quantityStr := r.PostFormValue("quantity")
	quantity64, err := strconv.ParseInt(quantityStr, 10, 64)
	if err != nil {
		v.Set("warning", "Quantity is not an integer.")
		v.Set("next", true)
		return
	}
	quantity := int(quantity64)
	expireStr := r.PostFormValue("expire")
	//expireSplit := strings.Split(expireStr, "//")
	//
	//expireToParse := expireSplit[2] + "-" + expireSplit[1] + "-" + expireSplit[0] + "T00:00:00+07:00"

	// GMTPlus7 := int((7 * time.Hour).Seconds())
	// bangkok := time.FixedZone("Bangkok Time", GMTPlus7)
	expire, err := time.Parse("2/1/2006", expireStr)
	if err != nil {
		v.Set("warning", "Expire is invalid.")
		v.Set("next", true)
		return
	}
	file, h, err := r.FormFile("uploadfile")
	if err != nil {
		v.Set("warning", err)
		v.Set("next", true)
		return
	}
	ext := filepath.Ext(h.Filename)

	meat, err := dbmodel.MakeMeat(r.PostFormValue("name"), r.PostFormValue("type"),
		r.PostFormValue("grade"), r.PostFormValue("des"), price, quantity, expire, ext)
	if err != nil {
		v.Set("warning", err.Error())
	} else {
		err = db.RegisMeat(meat)
		if err != nil {
			v.Set("warning", err.Error())
		} else {
			v.Set("success", "You have successfully add meat into the store.")
		}
	}
	id, err := db.GetMeatId(meat)
	if err != nil {
		v.Set("warning", err)
		v.Set("next", true)
		return
	}
	fname := "meat_" + id + ext
	err = db.EditFile(fname, file)
	if err != nil {
		v.Set("warning", "cannot upload image")
		v.Set("next", true)
		return
	}
	v.Set("next", true)
}

func Login(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model := dbmodel.User{}
	if user != model {
		v.Set("warning", "You are already logged in")
	} else {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := db.AuthenticateUser(r.PostFormValue("username"), r.PostFormValue("password"))
		if err != nil {
			v.Set("warning", "Invalid username/password")
		} else {
			session, err := s.Get(r, "user")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(3)
				return
			}

			if r.PostFormValue("remember") != "true" {
				session.Options.MaxAge = 0
			}
			session.Values["user"] = user
			session.Save(r, w)
			v.Set("success", "Login successful")
		}
	}

	v.Set("next", true)
}

func Logout(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	session, err := s.Get(r, "user")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.Values["user"] = dbmodel.User{}
	session.Save(r, w)

	v.Set("success", "Logout successful")
	v.Set("next", true)
}

func Profile(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.UserDetail{
		Menu: header,
	}

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		model.Fullname = ""
		model.Email = ""
		model.Address = ""
	} else {
		model.Fullname = user.Fullname
		model.Email = user.Email
		model.Address = user.Address
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "profile.html", model)
}

func ProfileEdit(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.UserDetail{
		Menu: header,
	}

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		model.Fullname = ""
		model.Email = ""
		model.Address = ""
	} else {
		model.Fullname = user.Fullname
		model.Email = user.Email
		model.Address = user.Address
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "profile-edit.html", model)
}

func EditProfile(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, ok := v.Get("user").(dbmodel.User)
	if !ok {
		v.Set("warning", "Unable to get user.")
	} else {
		err = db.UpdateUser(user.ID, r.PostFormValue("fullname"), r.PostFormValue("email"), r.PostFormValue("address"))
		if err != nil {
			v.Set("warning", "Edit profile fail.")
		} else {
			session, err := s.Get(r, "user")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			newUser, err := db.GetUser(user.ID)
			if err != nil {
				v.Set("warning", "Unable to get new user.")
			} else {
				session.Values["user"] = newUser
				session.Save(r, w)
				v.Set("success", "You have successfully edit your profile.")
			}
		}

		if r.PostFormValue("pass-old") != "" {
			if r.PostFormValue("pass") == r.PostFormValue("pass-repeat") {
				newHash, err := bcrypt.GenerateFromPassword([]byte(r.PostFormValue("pass")), bcrypt.DefaultCost)
				if err != nil {
					v.Set("warning", "Error// Unable to update password.")
				} else {
					err = db.UpdatePass(user.ID, r.PostFormValue("pass-old"), string(newHash))
					if err != nil {
						v.Set("warning", err.Error())
					} else {
						v.Set("success", "Update password successfully.")
					}
				}
			} else {
				v.Set("warning", "Password is not the same.")
			}
		}
	}
	v.Set("next", true)
}

func Mock(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	db.Mock()
	v.Set("next", true)
}

func Images(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, err := db.GetFile(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(b)
}

func SaleHistory(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "sale-history.html", model)
}
