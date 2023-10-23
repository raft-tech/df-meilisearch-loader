package configs

import (
	"reflect"
	"testing"
)

func TestConfig_ParseEnvs_Defaults(t *testing.T) {
	c := NewConfig()
	data := []struct {
		name     string
		expected string
	}{
		{"BrokerHost", "localhost:9092"},
		{"SchemaRegUrl", "localhost:8081"},
		{"Topic", "test-topic"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			c.ParseEnvs()
			if result := getField(c, d.name); result != d.expected {
				t.Errorf("Expected %s, got %s", d.expected, result)
			}

		})
	}
}

func TestConfig_ParseEnvs(t *testing.T) {
	c := NewConfig()
	data := []struct {
		name     string
		env      string
		envValue string
	}{
		{"BrokerHost", "KAFKA_BROKER_HOST", "customhost:9092"},
		{"SchemaRegUrl", "SCHEMA_REGISTRY_URL", "customhost:8081"},
		{"Topic", "KAFKA_TOPIC", "custom-topic"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Setenv(d.env, d.envValue)
			c.ParseEnvs()
			if result := getField(c, d.name); result != d.envValue {
				t.Errorf("Expected %s, got %s", d.envValue, result)
			}
		})
	}
}

func getField(c *Config, field string) string {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
