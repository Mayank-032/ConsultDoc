package entity

type Doctor struct {
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
	Fees    int     `json:"fees"`
	Slots   []Slots `json:"slots"`
}
