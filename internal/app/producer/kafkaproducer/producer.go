package kafkaproducer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	conn *kafka.Conn
}

func New(ctx context.Context, address string, topic string, partition int) (*Producer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, "accounts", 2)

	if err != nil {
		return nil, err
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second))

	return &Producer{conn: conn}, nil
}

func (p *Producer) Publish() {
	// TODO: implement Publish message to kafka
}
