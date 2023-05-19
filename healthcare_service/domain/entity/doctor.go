package entity

type Doctor struct {
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Address string  `json:"address"`
	Slots   []Slots `json:"slots"`
}

type Slots struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
