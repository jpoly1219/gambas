package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintSeries(t *testing.T) {
	seriesArray := []Series{
		{
			map[interface{}]interface{}{0: "alice", 1: "bob", 2: "charlie"},
			Index{[]interface{}{0, 1, 2}},
			"People",
		},
		{
			map[interface{}]interface{}{"a": "apple", "b": "banana", "c": "cherry"},
			Index{[]interface{}{"a", "b", "c"}},
			"Fruit",
		},
	}
	expectedArray := []string{
		"data: map[0:alice 1:bob 2:charlie] \nindexArray: {[0 1 2]} \nname: People\n",
		"data: map[a:apple b:banana c:cherry] \nindexArray: {[a b c]} \nname: Fruit\n",
	}

	for i, test := range seriesArray {
		output := test.PrintSeries()
		if output != expectedArray[i] {
			t.Fatalf("expected: %s, got: %s", expectedArray[i], output)
		}
	}
}

func TestAt(t *testing.T) {
	type atTest struct {
		arg1     Series
		arg2     interface{}
		expected interface{}
	}
	atTests := []atTest{
		{
			Series{
				map[interface{}]interface{}{0: "alice", 1: "bob", 2: "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
			0,
			"alice",
		},
		{
			Series{
				map[interface{}]interface{}{"a": "apple", "b": "banana", "c": "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
			"b",
			"banana",
		},
	}

	for _, test := range atTests {
		output := test.arg1.At(test.arg2)
		if output != test.expected {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestAtM(t *testing.T) {
	type atMTest struct {
		arg1     Series
		arg2     []interface{}
		expected []interface{}
	}
	atMTests := []atMTest{
		{
			Series{
				map[interface{}]interface{}{0: "alice", 1: "bob", 2: "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
			[]interface{}{0, 1},
			[]interface{}{"alice", "bob"},
		},
		{
			Series{
				map[interface{}]interface{}{"a": "apple", "b": "banana", "c": "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
			[]interface{}{"b", "c"},
			[]interface{}{"banana", "cherry"},
		},
	}

	for _, test := range atMTests {
		output := test.arg1.AtM(test.arg2)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestAtR(t *testing.T) {
	type atRTest struct {
		arg1     Series
		arg2     int
		arg3     int
		expected []interface{}
	}
	atRTests := []atRTest{
		{
			Series{
				map[interface{}]interface{}{0: "alice", 1: "bob", 2: "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
			0,
			2,
			[]interface{}{"alice", "bob", "charlie"},
		},
		{
			Series{
				map[interface{}]interface{}{"a": "apple", "b": "banana", "c": "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
			0,
			1,
			[]interface{}{"apple", "banana"},
		},
	}

	for _, test := range atRTests {
		output := test.arg1.AtR(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
