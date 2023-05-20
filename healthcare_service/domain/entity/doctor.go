package entity

type Doctor struct {
	Id      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Phone   string  `json:"phone,omitempty"`
	Address Address `json:"address,omitempty"`
	Fees    int     `json:"fees,omitempty"`
	Slots   []Slot  `json:"slots,omitempty"`
}
