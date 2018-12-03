package handler

import (
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
)

func PrepareProductPageModel(header pagemodel.Menu, oldLink string, productCount int, page int) pagemodel.Product {
	perPage := db.GetPerProductPage()
	pageCount := productCount / perPage
	if productCount%perPage > 0 {
		pageCount++
	}

	return pagemodel.Product{
		Menu:      header,
		Meats:     []pagemodel.MeatModel{},
		PageCount: pageCount,
		Page:      page,
		OldLink:   oldLink,
	}
}
