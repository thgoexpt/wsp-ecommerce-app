package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
)

func ProductSortTypePaging(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	// model := pagemodel.Product{
	// 	Menu:  header,
	// 	Meats: []pagemodel.MeatModel{},
	// }
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
		t.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	t.ExecuteTemplate(w, "product.html", model)
}

func ProductSearchPaging(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	// model := pagemodel.Product{
	// 	Menu:  header,
	// 	Meats: []pagemodel.MeatModel{},
	// }

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
		t.ExecuteTemplate(w, "product.html", model)
		return
	}

	for i := 0; i < len(meats); i++ {
		model.Meats = append(model.Meats, GetMeatModel(meats[i]))
	}

	t.ExecuteTemplate(w, "product.html", model)
}

func PrepareProductPageModel(header pagemodel.Menu, oldLink string, productCount int, page int) pagemodel.Product {
	pageCount := productCount / 10
	if productCount%10 > 0 {
		pageCount++
	}

	// var count []int
	// for i := 1; i <= len(count); i++ {
	// 	count[i] = i
	// }

	return pagemodel.Product{
		Menu:      header,
		Meats:     []pagemodel.MeatModel{},
		PageCount: pageCount,
		Page:      page,
		OldLink:   oldLink,
	}
}
