package app

import (
	"fmt"
	"yata-email/config"
	"yata-email/internal/email/rabbitmq"
	"yata-email/internal/lib/logger"
	"yata-email/internal/lib/metrics/trace"

	"os"
	"os/signal"
	"syscall"
)

func Run() {
	log := logger.NewLogger()
	cfg := config.LoadConfig()
	tracer := trace.InitTracer("yata-email")

	fmt.Println(cfg)
	defer log.Sync()

	amqpConn := rabbitmq.NewAmqpConnection(cfg.RabbitMQ)

	emailsConsumer := rabbitmq.NewEmailConsumer(amqpConn, log, tracer.Tracer)

	go func() {
		err := emailsConsumer.StartConsumer(
			cfg.RabbitMQ.QueueName,
			cfg.RabbitMQ.ConsumerTag,
			cfg.RabbitMQ.ExchangeName,
			cfg.RabbitMQ.BindingKey,
		)

		if err != nil {
			log.Errorf("StartConsumerErr: %v", err.Error())
		}

	}()

	log.Info("Email service started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
