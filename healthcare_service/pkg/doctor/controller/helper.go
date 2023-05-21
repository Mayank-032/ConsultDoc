package controller

import "errors"

type SendAppointmentDataReq struct {
	DoctorId        int    `json:"doctorId"`
	AppointmentLink string `json:"appointmentLink"`
}

func (req SendAppointmentDataReq) Validate() error {
	if req.DoctorId == 0 || len(req.AppointmentLink) == 0 {
		return errors.New("doctor_Id or appointment_link is missing")
	}
	return nil
}
