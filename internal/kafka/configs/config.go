package configs

import (
	"os"
)

type Config struct {
	BrokerHost   string
	MaxIdleTime  string
	SchemaRegUrl string
	Topic        string
}

const (
	kafkaHost    = "localhost:9092"
	schemaRegUrl = "localhost:8081"
	topic        = "test-topic"
)

func NewConfig() *Config {
	cfg := &Config{
		BrokerHost:   kafkaHost,
		SchemaRegUrl: schemaRegUrl,
		Topic:        topic,
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

	return cfg
}
