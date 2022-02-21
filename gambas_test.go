package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSeries(t *testing.T) {
	type newSeriesTest struct {
		arg1     []interface{}
		arg2     []interface{}
		arg3     string
		expected Series
	}

	newSeriesTests := []newSeriesTest{
		{
			[]interface{}{"alice", "bob", "charlie"},
			nil,
			"People",
			Series{
				map[interface{}]interface{}{0: "alice", 1: "bob", 2: "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
		},
		{
			[]interface{}{"apple", "banana", "cherry"},
			[]interface{}{"a", "b", "c"},
			"Fruit",
			Series{
				map[interface{}]interface{}{"a": "apple", "b": "banana", "c": "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
		},
	}

	for _, test := range newSeriesTests {
		output := NewSeries(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(output, output.index)) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
