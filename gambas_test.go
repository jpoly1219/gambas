package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateRangeIndex(t *testing.T) {
	type createRangeIndexTest struct {
		arg1     int
		expected Index
	}

	createRangeIndexTests := []createRangeIndexTest{
		{
			5,
			Index{[]interface{}{0, 1, 2, 3, 4}},
		},
		{
			10,
			Index{[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			1,
			Index{[]interface{}{0}},
		},
	}

	for _, test := range createRangeIndexTests {
		output := CreateRangeIndex(test.arg1)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(output)) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestNewSeries(t *testing.T) {
	type newSeriesTest struct {
		arg1     []interface{}
		arg2     string
		expected *Series
	}

	newSeriesTests := []newSeriesTest{
		{
			[]interface{}{"alice", "bob", "charlie"},
			"People",
			&Series{
				[]interface{}{"alice", "bob", "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
		},
		{
			[]interface{}{"apple", "banana", "cherry"},
			"Fruit",
			&Series{
				[]interface{}{"apple", "banana", "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
		},
	}

	for _, test := range newSeriesTests {
		output, err := NewSeries(test.arg1, test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(*output, output.index)) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestNewDataFrame(t *testing.T) {
	type newDataFrameTest struct {
		arg1     [][]interface{}
		arg2     []interface{}
		arg3     []interface{}
		expected *DataFrame
	}

	newDataFrameTests := []newDataFrameTest{
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]interface{}{"group a", "group b", "group c"},
			[]interface{}{"group a"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{1, 2, 3},
						Index{[]interface{}{0, 1, 2}},
						"group a",
					},
					{
						[]interface{}{4, 5, 6},
						Index{[]interface{}{0, 1, 2}},
						"group b",
					},
					{
						[]interface{}{7, 8, 9},
						Index{[]interface{}{0, 1, 2}},
						"group c",
					},
				},
				Index{[]interface{}{"group a", "group b", "group c"}},
				[]interface{}{"group a"},
				[]Index{{[]interface{}{0, 1, 2}}},
			},
		},
	}

	for _, test := range newDataFrameTests {
		output, err := NewDataFrame(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(*output, output.index, output.series[0])) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
