package controller

import (
	"context"
	"encoding/json"
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

	request := SendAppointmentDataReq{}
	err = json.Unmarshal(appointmentLinkRequestBytes, &request)
	if err == nil {
		err = request.Validate()
	}
	if err != nil {
		log.Printf("Error: %v,\n invalid details provided\n\n", err.Error())
		msg.Ack(false)
		return
	}

	doctorId := request.DoctorId
	appointmentLink := request.AppointmentLink
	err = dc.DoctorUCase.SendAppointmentLink(ctx, doctorId, appointmentLink)
	if err != nil {
		log.Printf("Error: %v,\n unable to send details to doctor\n\n", err.Error())
		msg.Ack(false)
		return
	}
	log.Printf("successfully send data to doctor and processed request\n\n")
	msg.Ack(true)
}