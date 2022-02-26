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

func TestNewDataFrame(t *testing.T) {
	type newDataFrameTest struct {
		arg1     [][]interface{}
		arg2     []interface{}
		arg3     []interface{}
		expected DataFrame
	}

	newDataFrameTests := []newDataFrameTest{
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]interface{}{0, 1, 2},
			[]interface{}{"group a", "group b", "group c"},
			DataFrame{
				map[interface{}]Series{
					"group a": {
						map[interface{}]interface{}{0: 1, 1: 2, 2: 3},
						Index{[]interface{}{0, 1, 2}},
						"group a",
					},
					"group b": {
						map[interface{}]interface{}{0: 4, 1: 5, 2: 6},
						Index{[]interface{}{0, 1, 2}},
						"group b",
					},
					"group c": {
						map[interface{}]interface{}{0: 7, 1: 8, 2: 9},
						Index{[]interface{}{0, 1, 2}},
						"group c",
					},
				},
				Index{[]interface{}{0, 1, 2}},
				Index{[]interface{}{"group a", "group b", "group c"}},
			},
		},
	}

	for _, test := range newDataFrameTests {
		output := NewDataFrame(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(output, output.series[0], output.index)) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
