package controller

import (
	"context"

	"github.com/streadway/amqp"
)

type IPatientController interface {
	ProcessCreateAppointmentRequest(ctx context.Context, data interface{}, msg amqp.Delivery)
}
