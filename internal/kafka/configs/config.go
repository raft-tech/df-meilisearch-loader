package configs

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	BrokerHost    string
	MaxIdleTime   string
	SchemaRegUrl  string
	Topic         string
	SaslMechanism string
	SaslUsername  string
	SaslSecret    string
	GroupId       string
}

const (
	kafkaHost     = "localhost:9092"
	schemaRegUrl  = "localhost:8081"
	topic         = "test-topic"
	saslMechanism = ""
	saslUsername  = ""
	saslSecret    = ""
	indeName      = ""
)

// The consumer groupId needs to precise enought so we can add new instances of consumers
// to be able to scale up and not to general to consider issues with multiple ingesiton
// piplines feeding data into different indexes. Also the groupId should not change on restarts
// or re-deployments since the app should pick where it left of.

func groupId(indexName string) string {
	return strings.TrimSuffix(fmt.Sprintf("%s-%s", "df-meilisearch", indexName), "-")
}

func NewConfig(indexName string) *Config {
	cfg := &Config{
		BrokerHost:    kafkaHost,
		SchemaRegUrl:  schemaRegUrl,
		Topic:         topic,
		SaslMechanism: saslMechanism,
		SaslUsername:  saslUsername,
		SaslSecret:    saslSecret,
		GroupId:       groupId(indexName),
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

	if consumerGroupId, exists := os.LookupEnv("KAFKA_CLIENT_CONSUMER_GROUPID"); exists {
		cfg.GroupId = consumerGroupId
	}

	return cfg
}
