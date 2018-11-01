package pagemodel

type ProductDetail struct {
	Menu
	ID          string
	Pic         string
	ProName     string
	Type        string
	Grade       string
	Description string
	Price       float64
	Expire      string
}
