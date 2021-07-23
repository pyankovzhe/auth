package kafkaproducer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	conn *kafka.Conn
}

func New(ctx context.Context, address string, topic string, partition int) (*Producer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, "accounts", partition)
	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(time.Second * 10))

	return &Producer{conn: conn}, nil
}

func (p *Producer) Close() {
	p.conn.Close()
}

func (p *Producer) Publish(value []byte) error {
	_, err := p.conn.WriteMessages(
		kafka.Message{
			Topic:     "accounts",
			Partition: 0,
			Key:       []byte{23},
			Value:     value,
		},
	)

	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
