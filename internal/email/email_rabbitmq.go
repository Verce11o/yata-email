package email

type EmailsConsumer interface {
	StartConsumer(queueName, consumerTag, exchangeName, bindingKey string) error
}
