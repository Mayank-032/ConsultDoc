package entity

type Patient struct {
	Name            string  `json:"name"`
	Phone           string  `json:"phone"`
	Address         Address `json:"address,omitempty"`
	AppointmentLink string  `json:"appointmentLink,omitempty"`
}
