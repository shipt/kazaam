package converter

import (
	"github.com/qntfy/kazaam/registry"
	"strconv"
	"testing"
)

func TestSton_Convert(t *testing.T) {
	registry.RegisterConverter("ston", &Ston{})
	c := registry.GetConverter("ston")

	table := []struct {
		value     string
		arguments string
		expected  string
	}{
		{`"5""`, ``, `5`,},
		{`"5.01"`, ``, `5.01`,},
		{`"-5.01"`, ``, `-5.01`,},
	}

	for _, test := range table {
		v, e := c.Convert(getTestData(), []byte(test.value), []byte(strconv.Quote(test.arguments)))

		if e != nil {
			t.Error("error running convert function")
		}

		if string(v) != test.expected {
			t.Error("unexpected result from convert")
			t.Log("Expected: {}", test.expected)
			t.Log("Actual: {}", string(v))
		}
	}
}