package producer

import (
	meiliConfig "meilisearch-loader/internal/meilisearch/configs"
	"reflect"
	"strconv"
	"testing"
)

func TestNewProducer(t *testing.T) {
	meiliCfg := meiliConfig.NewConfig()
	meiliCfg.ParseEnvs()

	p := NewProducer(meiliCfg)
	data := []struct {
		name     string
		expected string
	}{
		{"BatchSize", "50"},
		{"Index", "test-index"},
		{"IndexPrimaryKey", "id"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			r := getField(&p, d.name)
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

func getField(p *MeilisearchProducer, field string) any {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	if f.Type().String() == "string" {
		return f.String()
	}
	return f.Int()
}
