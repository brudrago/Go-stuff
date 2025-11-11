package queue

import (
	"fmt"
	"log"
	"reflect"
)

type QueueType int

const (
	RabbitMQ QueueType = iota
)

type QueueConnection interface {
	Publish([]byte) error
	Consume(chan<- QueueDTO) error
}

type Queue struct {
	qc QueueConnection
}

func New(qt QueueType, cfg any) (q *Queue, err error) {
	rt := reflect.TypeOf(cfg)

	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("invalid config type for RabbitMQ")
		}

		conn, err := NewRabbitMQConnection(cfg.(RabbitMQConfig))
		if err != nil {
			return nil, err
		}

		q.qc = conn

	default:
		log.Fatal("unsupported queue type")
	}
	return
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume(cdto chan<- QueueDTO) error {
	return q.qc.Consume(cdto)
}
