package handler

import (
	"net/http"

	"github.com/guitarpawat/middleware"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
	"github.com/guitarpawat/wsp-ecommerce/model/pagemodel"
)

func Owner(w http.ResponseWriter, r *http.Request, v *middleware.ValueMap) {
	header, ok := v.Get("header").(pagemodel.Menu)
	if !ok {
		header = defaultHeader
	}

	if header.UserType != dbmodel.TypeEmployee && header.UserType != dbmodel.TypeOwner {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	model := pagemodel.Owner{
		Menu: header,
	}

	v.Set("next", false)
	t.ExecuteTemplate(w, "owner.html", model)
}
