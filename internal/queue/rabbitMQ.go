package queue

import (
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

type RabbitMQConnection struct {
	cfg  RabbitMQConfig
	conn *ampq.Connection
}
