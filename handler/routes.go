package handler

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/db/mock"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
	"golang.org/x/crypto/bcrypt"
)

var s = sessions.NewCookieStore([]byte("NOT FOR PRODUCTION"))

var defaultHeader = pagemodel.Menu{
	Warning: "Something went wrong {defualt}",
}

func init() {
	gob.Register(bson.ObjectId(""))
	gob.Register(dbmodel.User{})
	gob.Register(dbmodel.Meat{})
	gob.Register(dbmodel.SalesHistory{})
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
	if ok {
		header.UserID = user.ID
		header.User = user.Username
		header.UserType = user.Type
		header.UserAddress = user.Address
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

	header.MeatInCartCart = []pagemodel.CartMeatModel{}
	header.CartTotal = 0.0
	cart, err := db.GetCartID(header.UserID)
	if err == nil {
		for _, cartMeat := range cart.Meats {
			meat, err := db.GetMeat(cartMeat.ID.Hex())
			if err != nil {
				// w.WriteHeader(http.StatusNotFound)
				// return
				// header.Warning = header.Warning + ", header: unable to get meats >> " + err.Error()
				fmt.Println("HEADER: " + meat.Name + err.Error())
			}
			cartMeat := GetCartMeatModel(meat, cartMeat.Quantity)
			header.CartTotal = header.CartTotal + cartMeat.Total
			header.MeatInCartCart = append(header.MeatInCartCart, cartMeat)
		}
	}

	v.Set("header", header)
	v.Set("next", true)
}

func ProceedCheckout(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}
	cart, err := db.GetCartID(header.UserID)
	if err != nil {
		v.Set("warning", "error: "+err.Error())
		v.Set("next", true)
		return
	}

	err = db.CommitSalesHistory(cart)
	for _, cartMeat := range cart.Meats {
		db.SoldMeat(cartMeat.ID, cartMeat.Quantity)
	}

	v.Set("next", false)
	if err != nil {
		v.Set("warning", "error: "+err.Error())
		v.Set("next", true)
		return
	}

	v.Set("success", "Thank you for your purchase, please come again.")
	v.Set("next", true)
}

func AddCart(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", true)
	if header.User != "" {
		vars := mux.Vars(r)
		quantity64, err := strconv.ParseInt(vars["quantity"], 10, 64)
		if err != nil {
			// w.WriteHeader(http.StatusNotFound)
			v.Set("warning", "AddCart: quantity parameter is wrong.")
			return
		}
		quantity := int(quantity64)

		user, err := db.GetUser(header.UserID)
		if err != nil {
			fmt.Println("Get User From ID Error! >> " + err.Error())
			// w.WriteHeader(http.StatusNotFound)
			v.Set("warning", "AddCart: unable to find user >> "+err.Error())
			return
		}

		db.UpdateCart(user.ID, bson.ObjectIdHex(vars["meatId"]), quantity)
		v.Set("success", "successfully add to cart!")
	} else {
		v.Set("warning", "AddCart: login before add cart!")
	}
}

func UpdateCart(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v.Set("next", true)

	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	cart := header.MeatInCartCart
	for _, meat := range cart {
		id := meat.ID
		quantityStr := r.PostFormValue("cartqty" + id)
		quantity64, err := strconv.ParseInt(quantityStr, 10, 64)
		if err != nil {
			v.Set("warning", "UpdateCart: Quantity is not an integer >> "+err.Error())
			return
		}
		quantity := int(quantity64)
		err = db.UpdateCart(header.UserID, bson.ObjectIdHex(id), quantity)
		if err != nil {
			v.Set("warning", "UpdateCart: unable to update cart >> "+err.Error())
			return
		}
		v.Set("success", "successfully update cart!")
	}
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
	discountStr := r.PostFormValue("discount")
	var discount float64
	if discountStr == "" {
		discount = -1.0
	} else {
		discount, err = strconv.ParseFloat(discountStr, 64)
		if err != nil {
			v.Set("warning", "Price is not a number.")
			v.Set("next", true)
			return
		}
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
		r.PostFormValue("grade"), r.PostFormValue("des"), price, discount, quantity, expire, ext)
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
	mock.Mock()
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

func RemoveMeatFromCart(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", true)
	vars := mux.Vars(r)
	meatID := bson.ObjectIdHex(vars["meatID"])
	err := db.RemoveMeat(header.UserID, meatID)
	if err != nil {
		v.Set("warning", "cart_rm: can'T rm meat from cart >> "+err.Error())
		return
	}
	v.Set("success", "Successfully remove meat from cart.")
}
