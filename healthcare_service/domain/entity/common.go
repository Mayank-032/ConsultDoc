package entity

type Slot struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Address struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
