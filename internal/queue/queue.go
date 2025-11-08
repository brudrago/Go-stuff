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
	Consume() error
}

type Queue struct {
	qc  QueueConnection
	cfg any
}

func New(qt QueueType, cfg any) (q *Queue, err error) {
	rt := reflect.TypeOf(cfg)

	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("invalid config type for RabbitMQ")
		}
		fmt.Println("not implemented")

	default:
		log.Fatal("unsupported queue type")
	}
	return
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume() error {
	return q.qc.Consume()
}
