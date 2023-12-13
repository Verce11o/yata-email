package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"yata-email/config"
)

func NewAmqpConnection(cfg config.RabbitMQ) *amqp.Connection {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("err while connection to amqp: %v", err.Error())
	}

	return conn

}
