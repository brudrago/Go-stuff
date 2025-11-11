package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

type RabbitMQConnection struct {
	cfg  RabbitMQConfig
	conn *amqp.Connection
}

func (r *RabbitMQConnection) Publish(msg []byte) error {
	c, err := r.conn.Channel()
	if err != nil {
		return err
	}
	mp := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "text/plain",
		Body:         msg,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	return c.PublishWithContext(
		ctx,
		"",
		r.cfg.TopicName,
		false,
		false,
		mp,
	)
}

func (r *RabbitMQConnection) Consume(cdto chan<- QueueDTO) error {
	//cria o canal
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}

	// cria a fila
	q, err := ch.QueueDeclare(
		r.cfg.TopicName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	// consome as mensagens
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for delivery := range msgs {
		dto := QueueDTO{}
		dto.Unmarshal(delivery.Body)
		cdto <- dto
	}

	return nil
}
