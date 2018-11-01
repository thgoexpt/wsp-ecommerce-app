package pagemodel

const TimeFormat = "02/01/2006"

type MeatModel struct {
	ID          string
	Pic         string
	ProName     string
	Type        string
	Grade       string
	Description string
	Price       float64
	Expire      string
	Quantity    int
	Total       float64
}
