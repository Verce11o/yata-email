package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
	SMTP     SMTP     `yaml:"smtp"`
}

type RabbitMQ struct {
	Username     string `yaml:"username" env-required:"true"`
	Password     string `yaml:"password" env-required:"true"`
	Host         string `yaml:"host" env-required:"true"`
	Port         string `yaml:"port" env-required:"true"`
	ExchangeName string `yaml:"exchangeName" env-required:"true"`
	QueueName    string `yaml:"queueName" env-required:"true"`
	ConsumerTag  string `yaml:"consumerTag" env-required:"true"`
	BindingKey   string `yaml:"bindingKey" env-required:"true"`
}

type SMTP struct {
	Host                  string `yaml:"host" env-required:"true"`
	Username              string `yaml:"username" env-required:"true"`
	Password              string `yaml:"password" env-required:"true"`
	PasswordEmailType     string `yaml:"password_email_type" env-required:"true"`
	EmailConfirmationType string `yaml:"email_confirmation_type" env-required:"true"`
}

func LoadConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("config.yml", &cfg); err != nil {
		log.Fatalf("error while reading config file: %s", err)
	}
	return &cfg

}
