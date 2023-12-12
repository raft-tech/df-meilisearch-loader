package configs

import (
	"os"

	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Config struct {
	BrokerHost    string
	MaxIdleTime   string
	SchemaRegUrl  string
	Topic         string
	SaslMechanism sasl.Mechanism
}

const (
	kafkaHost    = "localhost:9092"
	schemaRegUrl = "localhost:8081"
	topic        = "test-topic"
)

func NewConfig() *Config {
	cfg := &Config{
		BrokerHost:    kafkaHost,
		SchemaRegUrl:  schemaRegUrl,
		Topic:         topic,
		SaslMechanism: nil,
	}

	if kh, exists := os.LookupEnv("KAFKA_BROKER_HOST"); exists {
		cfg.BrokerHost = kh
	}
	if srh, exists := os.LookupEnv("SCHEMA_REGISTRY_URL"); exists {
		cfg.SchemaRegUrl = srh
	}
	if t, exists := os.LookupEnv("KAFKA_TOPIC"); exists {
		cfg.Topic = t
	}

	if auth, exists := os.LookupEnv("KAFKA_SASL_MECHANISM"); exists {
		var authMechanism scram.Algorithm = nil
		if auth == scram.SHA512.Name() {
			authMechanism = scram.SHA512
		} else if auth == scram.SHA256.Name() {
			authMechanism = scram.SHA256
		}
		kafkaClientUser, _ := os.LookupEnv("KAFKA_CLIENT_USERNAME")
		kafkaClientSecret, _ := os.LookupEnv("KAFKA_CLIENT_SECRET")

		if authMechanism != nil {
			cfg.SaslMechanism, _ = scram.Mechanism(authMechanism, kafkaClientUser, kafkaClientSecret)
		} else {
			cfg.SaslMechanism = plain.Mechanism{
				Username: kafkaClientUser,
				Password: kafkaClientSecret,
			}
		}
	}

	return cfg
}
