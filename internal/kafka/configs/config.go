package configs

import (
	"os"
)

type Config struct {
	BrokerHost    string
	MaxIdleTime   string
	SchemaRegUrl  string
	Topic         string
	SaslMechanism string
	SaslUsername  string
	SaslSecret    string
}

const (
	kafkaHost     = "localhost:9092"
	schemaRegUrl  = "localhost:8081"
	topic         = "test-topic"
	saslMechanism = ""
	saslUsername  = ""
	saslSecret    = ""
)

func NewConfig() *Config {
	cfg := &Config{
		BrokerHost:    kafkaHost,
		SchemaRegUrl:  schemaRegUrl,
		Topic:         topic,
		SaslMechanism: saslMechanism,
		SaslUsername:  saslUsername,
		SaslSecret:    saslSecret,
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

	if authMechanism, exists := os.LookupEnv("KAFKA_SASL_MECHANISM"); exists {
		cfg.SaslMechanism = authMechanism
	}

	if username, exists := os.LookupEnv("KAFKA_CLIENT_USERNAME"); exists {
		cfg.SaslUsername = username
	}

	if secret, exists := os.LookupEnv("KAFKA_CLIENT_SECRET"); exists {
		cfg.SaslSecret = secret
	}

	return cfg
}
