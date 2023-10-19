package shared

// Message is the output representation of a Kafka Message that only contains the payload.
type Message struct {
	Key   []byte
	Value []byte
}
