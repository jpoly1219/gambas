package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintSeries(t *testing.T) {
	seriesArray := []Series{
		{
			[]interface{}{"alice", "bob", "charlie"},
			Index{[]interface{}{0, 1, 2}},
			"People",
		},
		{
			[]interface{}{"apple", "banana", "cherry"},
			Index{[]interface{}{"a", "b", "c"}},
			"Fruit",
		},
	}
	expectedArray := []string{
		"data: [alice bob charlie] \nindexArray: {[0 1 2]} \nname: People\n",
		"data: [apple banana cherry] \nindexArray: {[a b c]} \nname: Fruit\n",
	}

	for i, test := range seriesArray {
		output := test.PrintSeries()
		if output != expectedArray[i] {
			t.Fatalf("expected: %s, got: %s", expectedArray[i], output)
		}
	}
}

func TestSeriesLoc(t *testing.T) {
	type atTest struct {
		arg1     Series
		arg2     interface{}
		expected interface{}
	}
	atTests := []atTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
			0,
			"alice",
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
			"b",
			"banana",
		},
	}

	for _, test := range atTests {
		output, err := test.arg1.Loc(test.arg2)
		if output != test.expected || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestSeriesLocM(t *testing.T) {
	type atMTest struct {
		arg1     Series
		arg2     []interface{}
		expected []interface{}
	}
	atMTests := []atMTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
			[]interface{}{0, 1},
			[]interface{}{"alice", "bob"},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
			[]interface{}{"b", "c"},
			[]interface{}{"banana", "cherry"},
		},
	}

	for _, test := range atMTests {
		output, err := test.arg1.LocM(test.arg2)
		if !cmp.Equal(output, test.expected) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestSeriesLocR(t *testing.T) {
	type atRTest struct {
		arg1     Series
		arg2     int
		arg3     int
		expected []interface{}
	}
	atRTests := []atRTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				Index{[]interface{}{0, 1, 2}},
				"People",
			},
			0,
			2,
			[]interface{}{"alice", "bob"},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				Index{[]interface{}{"a", "b", "c"}},
				"Fruit",
			},
			0,
			1,
			[]interface{}{"apple"},
		},
	}

	for _, test := range atRTests {
		output, err := test.arg1.LocR(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
