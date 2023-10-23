package unmarshall

import (
	"reflect"
	"strings"
	"testing"
)

func TestInto(t *testing.T) {
	data := []struct {
		name     string
		json     string
		expected map[string]any
		errMsg   string
	}{
		{"valid json", `{"test": "value 1"}`, map[string]any{"test": "value 1"}, ""},
		{"invalid json", `{"test" "value 1"}`, nil, "bad request: invalid character '\"' after object key"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			var msgValueJson map[string]any
			err := Into(&msgValueJson, strings.NewReader(d.json))
			if !reflect.DeepEqual(msgValueJson, d.expected) {
				t.Errorf("Expected %s, got %s", d.expected, msgValueJson)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Excpected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
