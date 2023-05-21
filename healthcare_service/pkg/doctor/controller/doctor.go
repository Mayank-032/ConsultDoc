package controller

import (
	"context"
	"encoding/json"
	"healthcare-service/domain"
	"healthcare-service/domain/interfaces/controller"
	"healthcare-service/domain/interfaces/usecase"
	"log"

	"github.com/streadway/amqp"
)

type DoctorController struct {
	DoctorUCase usecase.IDoctorUseCase
}

func NewDoctorController(doctorUCase usecase.IDoctorUseCase) controller.IDoctorController {
	return DoctorController{
		DoctorUCase: doctorUCase,
	}
}

func (dc DoctorController) ProcessSendAppointmentLinkRequest(ctx context.Context, data interface{}, msg amqp.Delivery) {
	appointmentLinkRequestBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error: %v,\n failed_to_marshal_request_body\n\n", err.Error())
		msg.Ack(false)
		return
	}

	var appointmentLink string
	err = json.Unmarshal(appointmentLinkRequestBytes, &appointmentLink)
	if err != nil || appointmentLink == domain.EmptyString {
		log.Printf("Error: %v,\n invalid details provided\n\n", err.Error())
		msg.Ack(false)
		return
	}
}