package controller

import (
	"context"

	"github.com/streadway/amqp"
)

type IDoctorController interface {
	ProcessSendAppointmentLinkRequest(ctx context.Context, data interface{}, msg amqp.Delivery)
}
