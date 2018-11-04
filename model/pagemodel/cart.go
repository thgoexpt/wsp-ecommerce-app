package pagemodel

type Cart struct {
	Menu
	MeatsInCart []CartMeatModel
	Total       float64
}
