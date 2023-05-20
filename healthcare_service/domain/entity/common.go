package entity

type Slots struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Address struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
