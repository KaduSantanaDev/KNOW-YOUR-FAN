package ports

type ProducerInterface interface {
	Publish(message, topic string, key []byte) error
	Close()
}
