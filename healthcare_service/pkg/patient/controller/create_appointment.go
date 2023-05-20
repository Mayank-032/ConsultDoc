package controller

import (
	"context"
	"encoding/json"
	"healthcare-service/domain/interfaces/controller"
	"healthcare-service/domain/interfaces/repository"
	"healthcare-service/domain/interfaces/usecase"
	"log"

	"github.com/streadway/amqp"
)

type PatientController struct {
	PatientUCase usecase.IPatientUseCase
	PatientRepo  repository.IPatientRepository
}

func NewPatientController(patientUCase usecase.IPatientUseCase, patientRepo repository.IPatientRepository) controller.IPatientController {
	return PatientController{
		PatientUCase: patientUCase,
		PatientRepo:  patientRepo,
	}
}

func (pc PatientController) CreateAppointment(ctx context.Context, data interface{}, msg amqp.Delivery) {
	patientDataBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error: %v,\n failed_to_marshal_request_body\n\n", err.Error())
		msg.Ack(false)
		return
	}

	request := CreateAppointmentRequest{}
	err = json.Unmarshal(patientDataBytes, &request)
	if err == nil {
		err = request.Validate()
	}
	if err != nil {
		log.Printf("Error: %v,\n invalid details provided\n\n", err.Error())
		msg.Ack(false)
		return
	}

	patient := request.toPatientDto()
	doctor := request.toDoctorDto()
	err = pc.PatientUCase.CreateAppointmentReceipt(ctx, patient, doctor)
	if err != nil {
		log.Printf("Error: %v,\n unable to create appointment receipt\n\n", err.Error())
		msg.Ack(false)
		return
	}
	msg.Ack(true)
}
