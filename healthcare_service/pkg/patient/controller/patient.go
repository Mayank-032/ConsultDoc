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
	Conn         *amqp.Connection
	PatientUCase usecase.IPatientUseCase
	PatientRepo  repository.IPatientRepository
}

func NewPatientController(conn *amqp.Connection, patientUCase usecase.IPatientUseCase, patientRepo repository.IPatientRepository) controller.IPatientController {
	return PatientController{
		Conn:         conn,
		PatientUCase: patientUCase,
		PatientRepo:  patientRepo,
	}
}

func (pc PatientController) ProcessCreateAppointmentRequest(ctx context.Context, data interface{}, msg amqp.Delivery) {
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
	appointmentLink, err := pc.PatientUCase.CreateAppointmentReceipt(ctx, patient, doctor)
	if err != nil {
		log.Printf("Error: %v,\n unable to create appointment receipt\n\n", err.Error())
		msg.Ack(false)
		return
	}

	err = publishAppointmentLinkToQueue(ctx, pc.Conn, appointmentLink, doctor.Id)
	if err != nil {
		log.Printf("Error: %v,\n unable to publish appointment link to queue\n\n", err.Error())
		msg.Ack(false)
		return
	}

	patient.AppointmentLink = appointmentLink
	err = pc.PatientRepo.SavePatientDetails(ctx, patient)
	if err != nil {
		log.Printf("Error: %v,\n unable to save patient details\n\n", err.Error())
		msg.Ack(false)
		return
	}
	log.Printf("successfully saved patient details and processed request\n\n")
	msg.Ack(true)
}
