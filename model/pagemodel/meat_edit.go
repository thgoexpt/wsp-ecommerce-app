package pagemodel

const AddMeatTxt = "Add Meat"
const EditMeatTxt = "Edit Meat"

type MeatEdit struct {
	Menu
	MeatModel
	State string
}

func (m MeatEdit) IsAddMeat() bool {
	return m.State == AddMeatTxt
}
