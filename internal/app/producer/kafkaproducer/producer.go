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
	conn, err := kafka.DialLeader(ctx, "tcp", address, topic, partition)
	if err != nil {
		return nil, err
	}

	return &Producer{conn: conn}, nil
}

func (p *Producer) Close() {
	p.conn.Close()
}

func (p *Producer) Publish(value []byte) error {
	p.conn.SetWriteDeadline(time.Now().Add(time.Second * 1))

	_, err := p.conn.WriteMessages(
		kafka.Message{
			Topic:     "accounts",
			Partition: 0,
			Value:     value,
		},
	)

	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
