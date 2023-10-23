package configs

import (
	"reflect"
	"strconv"
	"testing"
)

func TestConfig_ParseEnvs_Defaults(t *testing.T) {
	c := NewConfig()
	data := []struct {
		name     string
		expected string
	}{
		{"ApiKey", "aSampleMasterKey"},
		{"BatchSize", "50"},
		{"IndexPrimaryKey", "id"},
		{"Index", "test-index"},
		{"Url", "localhost:7700"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			c.ParseEnvs()
			r := getField(c, d.name)
			var result string
			if reflect.TypeOf(r).String() == "int64" {
				result = strconv.Itoa(int(r.(int64)))
			} else {
				result = r.(string)
			}
			if result != d.expected {
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
		{"ApiKey", "MEILISEARCH_API_KEY", "testApiKey"},
		{"BatchSize", "MEILISEARCH_INSERT_BATCH_SIZE", "1000"},
		{"IndexPrimaryKey", "MEILISEARCH_INDEX_PRIMARY_KEY", "testId"},
		{"Index", "MEILISEARCH_INDEX", "custom-index"},
		{"Url", "MEILISEARCH_URL", "customhost:7700"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			t.Setenv(d.env, d.envValue)
			c.ParseEnvs()
			r := getField(c, d.name)
			var result string
			if reflect.TypeOf(r).String() == "int64" {
				result = strconv.Itoa(int(r.(int64)))
			} else {
				result = r.(string)
			}
			if result != d.envValue {
				t.Errorf("Expected %s, got %s", d.envValue, result)
			}
		})
	}
}

func getField(c *Config, field string) any {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Type().String() == "string" {
		return f.String()
	}
	return f.Int()
}
