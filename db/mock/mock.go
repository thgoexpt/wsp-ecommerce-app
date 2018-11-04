package mock

import (
	"github.com/guitarpawat/wsp-ecommerce/db"
	"github.com/guitarpawat/wsp-ecommerce/env"
)

func init() {
	if env.GetEnv() != env.Production {
		Mock()
	}
}

func Mock() {
	db.MockUser()
	db.MockMeat()
	db.MockSalesHistory()
}
