package controller

import (
	"context"
	"encoding/json"
	"errors"
	"healthcare-service/domain/entity"
	"healthcare-service/rabbitmq"
	"healthcare-service/rabbitmq/publisher"
	"log"

	"github.com/streadway/amqp"
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

func publishAppointmentLinkToQueue(ctx context.Context, conn *amqp.Connection, appointmentLink string) error {
	defer conn.Close()
	if conn.IsClosed() {
		log.Printf("Closed Connect")

		err := rabbitmq.Connect()
		if err != nil {
			log.Printf("Error: %v, unable to init rabbitmq", err.Error())
			return errors.New("unable to init rabbitmq conn")
		}
		conn = rabbitmq.Conn
	}
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Printf("Error: %v,\n failed_to_create_channel", err.Error())
		return errors.New("unable to create channel")
	}
	defer amqpChannel.Close()

	publishData := publisher.PublishTaskRequestData{}
	publishData.Data = appointmentLink
	reqBytes, err := json.Marshal(publishData)
	if err != nil {
		log.Printf("Error: %v,\n invalid_json_format", err.Error())
		return errors.New("invalid json format")
	}

	publishRequest := publisher.PublishTaskRequest{}
	publishRequest.QueueName = "healthcare_service_appointment_link"
	publishRequest.ExchangeName = "healthcare_service"
	publishRequest.RoutingKey = "healthcare_service_appointment_link"
	publishRequest.ReqBytes = reqBytes
	err = publishRequest.PublishMessage(ctx, amqpChannel)
	if err != nil {
		log.Printf("Error: %v,\n failed_to_publish_message\n\n", err.Error())
		return errors.New("unable to publish message")
	}
	return nil
}