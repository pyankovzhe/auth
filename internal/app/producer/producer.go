package producer

type Producer interface {
	Publish([]byte) error
}
