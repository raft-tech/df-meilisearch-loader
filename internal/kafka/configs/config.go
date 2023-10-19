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

var (
	kafkaHost    = "localhost:9092"
	schemaRegUrl = "localhost:8081"
	topic        = "test-topic"
)

func NewConfig() *Config {
	return &Config{
		BrokerHost:   kafkaHost,
		SchemaRegUrl: schemaRegUrl,
		Topic:        topic,
	}
}

func (cfg *Config) ParseEnvs() {
	if kh, exists := os.LookupEnv("KAFKA_BROKER_HOST"); exists {
		cfg.BrokerHost = kh
	}
	if srh, exists := os.LookupEnv("SCHEMA_REGISTRY_URL"); exists {
		cfg.SchemaRegUrl = srh
	}
	if t, exists := os.LookupEnv("KAFKA_TOPIC"); exists {
		cfg.Topic = t
	}
}
