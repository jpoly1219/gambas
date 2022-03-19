package gambas

import (
	"fmt"
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
				Index{[]interface{}{0, 1, 2}},
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
				[]Index{{[]interface{}{1, 2, 3}}},
			},
		},
	}

	for _, test := range newDataFrameTests {
		output, err := NewDataFrame(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(*output, output.index[0], output.series[0])) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestCheckTypeIntegrity(t *testing.T) {
	type checkTypeIntegrityTest struct {
		arg1     []interface{}
		expected bool
	}
	checkTypeIntegrityTests := []checkTypeIntegrityTest{
		{
			[]interface{}{0, 1, 2, 3, 4, 5, 6},
			true,
		},
		{
			[]interface{}{3.4, 5.6, 2.4, 6.5, 7},
			true,
		},
		{
			[]interface{}{1, 2, "3", "4", 5},
			false,
		},
		{
			[]interface{}{"a", "b", "c"},
			true,
		},
		{
			[]interface{}{true, false, true},
			false,
		},
	}

	for _, test := range checkTypeIntegrityTests {
		output, err := checkTypeIntegrity(test.arg1)
		if !cmp.Equal(output, test.expected) || (output != test.expected && err == nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestI2f(t *testing.T) {
	type i2fTest struct {
		arg1          interface{}
		expected      float64
		expectedError error
	}
	i2fTests := []i2fTest{
		{1, 1.0, nil},
		{-5, -5.0, nil},
		{"45", 0.0, fmt.Errorf("%v is not a number", "45")},
		{2.0, 2.0, nil},
		{-5.6, -5.6, nil},
		{0.0, 0.0, nil},
		{0.000, 0.0, nil},
	}

	for _, test := range i2fTests {
		output, err := i2f(test.arg1)
		if !cmp.Equal(output, test.expected) || !cmp.Equal(fmt.Sprint(err), fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
