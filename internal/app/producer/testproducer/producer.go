package testproducer

type Producer struct{}

func (p *Producer) Publish(value []byte) error {
	return nil
}
