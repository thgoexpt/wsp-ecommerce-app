package pagemodel

type Cart struct {
	Menu
	MeatsInCart []CartMeatModel
	CartTotal   float64
}
