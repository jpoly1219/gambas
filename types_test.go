package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintSeries(t *testing.T) {
	seriesArray := []Series{
		{
			Data:       map[Index]interface{}{{0}: "alice", {1}: "bob", {2}: "charlie"},
			IndexArray: []Index{{0}, {1}, {2}},
			Name:       "People",
		},
		{
			Data:       map[Index]interface{}{{"a"}: "apple", {"b"}: "banana", {"c"}: "cherry"},
			IndexArray: []Index{{"a"}, {"b"}, {"c"}},
			Name:       "Fruit",
		},
	}
	expectedArray := []string{
		"Data: map[{0}:alice {1}:bob {2}:charlie] \nIndexArray: [{0} {1} {2}] \nName: People\n",
		"Data: map[{a}:apple {b}:banana {c}:cherry] \nIndexArray: [{a} {b} {c}] \nName: Fruit\n",
	}

	for i, test := range seriesArray {
		output := test.PrintSeries()
		if output != expectedArray[i] {
			t.Fatalf("expected: %s, got: %s", expectedArray[i], output)
		}
	}
}

func TestLoc(t *testing.T) {
	type locTest struct {
		arg1     Series
		arg2     Index
		expected interface{}
	}
	locTests := []locTest{
		{
			Series{
				map[Index]interface{}{{0}: "alice", {1}: "bob", {2}: "charlie"},
				[]Index{{0}, {1}, {2}},
				"People",
			},
			Index{0},
			"alice",
		},
		{
			Series{
				map[Index]interface{}{{"a"}: "apple", {"b"}: "banana", {"c"}: "cherry"},
				[]Index{{"a"}, {"b"}, {"c"}},
				"Fruit",
			},
			Index{"b"},
			"banana",
		},
	}

	for _, test := range locTests {
		output := test.arg1.Loc(test.arg2)
		if output != test.expected {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestLocM(t *testing.T) {
	type locMTest struct {
		arg1     Series
		arg2     []Index
		expected []interface{}
	}
	locMTests := []locMTest{
		{
			Series{
				map[Index]interface{}{{0}: "alice", {1}: "bob", {2}: "charlie"},
				[]Index{{0}, {1}, {2}},
				"People",
			},
			[]Index{{0}, {1}},
			[]interface{}{"alice", "bob"},
		},
		{
			Series{
				map[Index]interface{}{{"a"}: "apple", {"b"}: "banana", {"c"}: "cherry"},
				[]Index{{"a"}, {"b"}, {"c"}},
				"Fruit",
			},
			[]Index{{"b"}, {"c"}},
			[]interface{}{"banana", "cherry"},
		},
	}

	for _, test := range locMTests {
		output := test.arg1.LocM(test.arg2)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
