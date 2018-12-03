package handler

import (
	"github.com/gorilla/mux"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/handler/template"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Home{
		Menu:     header,
		ShowCase: []pagemodel.MeatModel{},
		Sale:     []pagemodel.MeatModel{},
	}

	meats, err := db.GetMeatsPaging(8, 1)
	if err != nil {
		v.Set("warning", "Home: unable to get showcase meats >> "+err.Error())
	} else {
		for i := 0; i < len(meats); i++ {
			meat := GetMeatModel(meats[i])
			model.ShowCase = append(model.ShowCase, meat)
		}
	}
	saleMeats, err := db.GetSaleMeat(8, 1)
	if err != nil {
		v.Set("warning", "Home: unable to get sale meats >> "+err.Error())
	} else {
		for i := 0; i < len(saleMeats); i++ {
			saleMeats := GetMeatModel(saleMeats[i])
			model.Sale = append(model.Sale, saleMeats)
		}
	}

	v.Set("next", false)
	template.T.ExecuteTemplate(w, "home.html", model)
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
	template.T.ExecuteTemplate(w, "about.html", model)
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
	template.T.ExecuteTemplate(w, "contact.html", model)
}

func Checkout(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Home{
		Menu: header,
	}

	v.Set("next", false)

	cart, err := db.GetCartID(header.UserID)
	if err != nil {
		model.Warning = err.Error()
		template.T.ExecuteTemplate(w, "home.html", model)
		return
	}

	if len(cart.Meats) == 0 {
		model.Warning = "the cart is empty"
		template.T.ExecuteTemplate(w, "home.html", model)
		return
	}

	template.T.ExecuteTemplate(w, "checkout.html", model)
}

func Cart(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.Cart{
		Menu: header,
	}

	v.Set("next", false)
	template.T.ExecuteTemplate(w, "cart.html", model)
}

func Product(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	proCount, _ := db.CountProduct("", "", 0, -1)
	model := PrepareProductPageModel(header,
		"/product/",
		proCount,
		1,
	)

	v.Set("next", false)
	meats, err := db.GetMeatsPaging(db.GetPerProductPage(), model.Page)
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "Product: unable to get all meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
}

func ProductSortType(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	// model := pagemodel.Product{
	// 	Menu:  header,
	// 	Meats: []pagemodel.MeatModel{},
	// }
	vars := mux.Vars(r)
	proCount, _ := db.CountProduct("", vars["meattype"], 0, -1)
	model := PrepareProductPageModel(header,
		"/product/sort/type="+vars["meattype"]+"&priceSort="+vars["price_sort"]+"/",
		proCount,
		1,
	)

	v.Set("next", false)
	meats, err := db.SortType(vars["meattype"], vars["price_sort"])
	// meats, err := db.SortType(vars["meattype"], "price")
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "ProductSortType: unable to get sorted meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
}

func ProductSearch(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	// model := pagemodel.Product{
	// 	Menu:  header,
	// 	Meats: []pagemodel.MeatModel{},
	// }

	vars := mux.Vars(r)

	startPrice, err := strconv.ParseFloat(vars["startPrice"], 64)
	if err != nil {
		v.Set("warning", "ProductSearch: startPrice is not a number.")
		v.Set("next", true)
		return
	}
	endPrice, err := strconv.ParseFloat(vars["endPrice"], 64)
	if err != nil {
		v.Set("warning", "ProductSearch: startPrice is not a number.")
		v.Set("next", true)
		return
	}

	proCount, _ := db.CountProduct(vars["name"], "", startPrice, endPrice)
	model := PrepareProductPageModel(
		header,
		"/product/search/name="+vars["name"]+"&startPrice="+vars["startPrice"]+"&endPrice="+vars["endPrice"]+"&priceSort="+vars["price_sort"]+"/",
		proCount,
		1,
	)

	v.Set("next", false)
	meats, err := db.Search(vars["name"], startPrice, endPrice, vars["price_sort"])
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "ProductSearch: unable to get searched meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
}

func ProductDetail(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	model := pagemodel.ProductDetail{
		Menu: header,
	}

	vars := mux.Vars(r)
	v.Set("next", false)
	meat, err := db.GetMeat(vars["meatId"])
	if err != nil {
		// meat = dbmodel.Meat{}
		v.Set("warning", "ProductDetail: >> "+err.Error())
		template.T.ExecuteTemplate(w, "product-detail.html", model)
		return
	}
	model.MeatModel = GetMeatModel(meat)
	db.ViewMeat(meat.ID)

	template.T.ExecuteTemplate(w, "product-detail.html", model)
}

func Sale(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil || page == 0 {
		page = 1
	}

	model := PrepareProductPageModel(header,
		"/product/",
		0,
		page,
	)

	v.Set("next", false)
	meats, err := db.GetSaleMeat(80, page)
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "Product: unable to get all meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	proCount := len(meats)
	model = PrepareProductPageModel(header,
		"/product/sale/",
		proCount,
		page,
	)

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
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
	template.T.ExecuteTemplate(w, "comingsoon.html", model)
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
	template.T.ExecuteTemplate(w, "add-product.html", model)
}

func ProductStock(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	v.Set("next", false)

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.Stock{
		Menu:  header,
		Meats: []pagemodel.MeatModel{},
	}

	meats, err := db.GetAllMeatsNoFilter()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product-stock.html", model)
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

	model := pagemodel.MeatEdit{
		Menu:  header,
		State: pagemodel.AddMeatTxt,
	}
	template.T.ExecuteTemplate(w, "add-product.html", model)
	v.Set("next", false)
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
	template.T.ExecuteTemplate(w, "profile.html", model)
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
	template.T.ExecuteTemplate(w, "profile-edit.html", model)
}

func SaleHistory(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	sh, err := db.GetUserSalesHistory(header.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(sh) == 0 {
		header.Warning = "No sale record found"
	}

	model, err := pagemodel.ToSalesHistoryPageModel(sh, header)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	v.Set("next", false)
	template.T.ExecuteTemplate(w, "sale-history.html", model)
}

func Owner(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}
	v.Set("next", false)

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.Owner{
		Menu:     header,
		SoldMeat: []pagemodel.CartMeatModel{},
	}

	soldMeat, err := db.GetSoldMeats()
	if err != nil {
		v.Set("warning", "Owner: unable to get data >> "+err.Error())
		template.T.ExecuteTemplate(w, "owner.html", model)
		return
	}
	for _, meatState := range soldMeat {
		meat, err := db.GetMeat(meatState.Meat.Hex())
		if err != nil {
			v.Set("warning", "Owner: unable to get meat >> "+err.Error())
			template.T.ExecuteTemplate(w, "owner.html", model)
			return
		}
		soldMeatModel := GetCartMeatModel(meat, meatState.Sold)
		model.SoldMeat = append(model.SoldMeat, soldMeatModel)
	}

	template.T.ExecuteTemplate(w, "owner.html", model)
}

func ProductPaging(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	vars := mux.Vars(r)

	page := 1
	page64, err := strconv.ParseInt(vars["page"], 10, 64)
	if err != nil {
		v.Set("warning", "ProductSortTypePaging: page is not integer >> "+err.Error())
		v.Set("next", false)
		return
	}
	page = int(page64)

	proCount, _ := db.CountProduct("", "", 0, -1)
	model := PrepareProductPageModel(header,
		"/product/",
		proCount,
		page,
	)

	v.Set("next", false)
	meats, err := db.GetMeatsPaging(db.GetPerProductPage(), model.Page)
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "Product: unable to get all meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
}

func ProductSortTypePaging(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	vars := mux.Vars(r)

	page := 1
	page64, err := strconv.ParseInt(vars["page"], 10, 64)
	if err != nil {
		v.Set("warning", "ProductSortTypePaging: page is not integer >> "+err.Error())
		v.Set("next", false)
		return
	}
	page = int(page64)

	proCount, _ := db.CountProduct("", vars["meattype"], 0, -1)
	model := PrepareProductPageModel(header,
		"/product/sort/type="+vars["meattype"]+"&priceSort="+vars["price_sort"]+"/",
		proCount,
		page,
	)

	v.Set("next", false)
	meats, err := db.SortTypePaging(vars["meattype"], page, vars["price_sort"])
	// meats, err := db.SortType(vars["meattype"], "price")
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "ProductSortTypePaging: unable to get sorted meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
}

func ProductSearchPaging(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	vars := mux.Vars(r)

	page := 1
	page64, err := strconv.ParseInt(vars["page"], 10, 64)
	if err != nil {
		v.Set("warning", "ProductSortTypePaging: page is not integer >> "+err.Error())
		v.Set("next", false)
		return
	}
	page = int(page64)

	startPrice, err := strconv.ParseFloat(vars["startPrice"], 64)
	if err != nil {
		v.Set("warning", "ProductSearchPaging: startPrice is not a number.")
		v.Set("next", true)
		return
	}
	endPrice, err := strconv.ParseFloat(vars["endPrice"], 64)
	if err != nil {
		v.Set("warning", "ProductSearchPaging: startPrice is not a number.")
		v.Set("next", true)
		return
	}

	proCount, _ := db.CountProduct(vars["name"], "", startPrice, endPrice)
	model := PrepareProductPageModel(
		header,
		"/product/search/name="+vars["name"]+"&startPrice="+vars["startPrice"]+"&endPrice="+vars["endPrice"]+"&priceSort="+vars["price_sort"]+"/",
		proCount,
		page,
	)

	v.Set("next", false)
	meats, err := db.SearchPaging(vars["name"], startPrice, endPrice, vars["price_sort"], page)
	if err != nil {
		// meats = []dbmodel.Meat{}
		v.Set("warning", "ProductSearchPaging: unable to get searched meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	template.T.ExecuteTemplate(w, "product.html", model)
}

func EditMeat(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.MeatEdit{
		Menu:  header,
		State: pagemodel.EditMeatTxt,
	}

	vars := mux.Vars(r)
	meat, err := db.GetMeat(vars["meatID"])
	if err != nil {
		v.Set("warning", "EditMeat: unable to get all meats >> "+err.Error())
		template.T.ExecuteTemplate(w, "add-product.html", model)
		return
	}
	model.MeatModel = GetMeatModel(meat)

	v.Set("next", false)
	template.T.ExecuteTemplate(w, "add-product.html", model)
}

