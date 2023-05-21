package controller

import (
	"errors"
	"healthcare-service/domain/entity"
)

type CreateAppointmentRequest struct {
	PatientName     string   `json:"patientName"`
	PatientPhone    string   `json:"patientPhone"`
	PatientAddress  []string `json:"patientAddress"`
	DoctorName      string   `json:"doctorName"`
	DoctorPhone     string   `json:"doctorPhone"`
	DoctorAddress   []string `json:"doctorAddress"`
	AppointmentSlot []string `json:"appointmentSlot"`
	Fees            int      `json:"fees"`
}

func (req CreateAppointmentRequest) Validate() error {
	if len(req.PatientName) == 0 || len(req.DoctorName) == 0 {
		return errors.New("invalid doctor or patient name")
	}
	if len(req.DoctorPhone) != 0 || len(req.PatientPhone) == 0 {
		return errors.New("invalid doctor or patient phone number")
	}
	if len(req.DoctorAddress) != 2 || len(req.PatientAddress) != 2 {
		return errors.New("invalid doctor's address or patient's address")
	}
	if len(req.AppointmentSlot) != 2 {
		return errors.New("invalid slot timing")
	}
	return nil
}

func (req CreateAppointmentRequest) toPatientDto() entity.Patient {
	address := entity.Address{
		Latitude:  req.PatientAddress[0],
		Longitude: req.PatientAddress[1],
	}
	return entity.Patient{
		Name:    req.PatientName,
		Phone:   req.PatientPhone,
		Address: address,
	}
}

func (req CreateAppointmentRequest) toDoctorDto() entity.Doctor {
	address := entity.Address{
		Latitude:  req.DoctorAddress[0],
		Longitude: req.DoctorAddress[1],
	}

	slots := []entity.Slot{}
	slot := entity.Slot{
		StartTime: req.AppointmentSlot[0],
		EndTime:   req.AppointmentSlot[1],
	}
	slots = append(slots, slot)

	return entity.Doctor{
		Name:    req.DoctorName,
		Phone:   req.DoctorPhone,
		Address: address,
		Fees:    req.Fees,
		Slots:   slots,
	}
}
