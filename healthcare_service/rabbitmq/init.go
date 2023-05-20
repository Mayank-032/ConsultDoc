package rabbitmq

import (
	"errors"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
func Connect() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return errors.New("unable to connect to rabbitmq")
	}
	Conn = conn
	return nil
}