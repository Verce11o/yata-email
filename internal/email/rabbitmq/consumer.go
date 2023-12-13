package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"log"
)

type EmailConsumer struct {
	AmqpConn *amqp.Connection
	log      *zap.SugaredLogger
	trace    trace.Tracer
}

func NewEmailConsumer(amqpConn *amqp.Connection, log *zap.SugaredLogger, trace trace.Tracer) *EmailConsumer {
	return &EmailConsumer{AmqpConn: amqpConn, log: log, trace: trace}
}

func (c *EmailConsumer) createChannel(exchangeName, queueName, bindingKey string) *amqp.Channel {

	ch, err := c.AmqpConn.Channel()

	if err != nil {
		panic(err)
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchangeName,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	return ch

}

func (c *EmailConsumer) StartConsumer(queueName, consumerTag, exchangeName, bindingKey string) error {
	ch := c.createChannel(exchangeName, queueName, bindingKey)
	defer ch.Close()

	deliveries, err := ch.Consume(
		queueName,
		consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		i := i
		go c.worker(i, deliveries)
	}

	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	c.log.Infof("Notify close: %v", chanErr)

	return chanErr

}

func (c *EmailConsumer) worker(index int, messages <-chan amqp.Delivery) {
	for message := range messages {
		log.Printf("Worker #%d: %v", index, string(message.Body))
		err := message.Ack(false)
		if err != nil {
			c.log.Errorf("failed to acknowledge delivery: %v", err)
		}
	}
	c.log.Info("Channel closed")
}
