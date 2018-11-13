package pagemodel

type Product struct {
	Menu
	Meats     []MeatModel
	PageCount int
	Page      int
	OldLink   string
}
