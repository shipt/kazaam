package converter

import (
	"strconv"
	"testing"

	"github.com/semarcial/kazaam/v1/registry"
)

func TestSplit_Convert(t *testing.T) {

	registry.RegisterConverter("split", &Split{})
	c := registry.GetConverter("split")

	table := []struct {
		value     string
		arguments string
		expected  string
	}{
		{`"a|b|c"`, `|`, `["a","b","c"]`},
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
