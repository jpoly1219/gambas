package gambas

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSeriesPrint(t *testing.T) {
	type seriesPrintTest struct {
		arg1 Series
	}
	seriesPrintTests := []seriesPrintTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a", "red"}},
						{1, []interface{}{"b", "yellow"}},
						{2, []interface{}{"c", "red"}},
					},
					[]string{"ID", "Color"},
				},
				"Fruit",
			},
		},
	}

	for _, test := range seriesPrintTests {
		test.arg1.Print()
	}
}

func TestSeriesPrintRange(t *testing.T) {
	type printRangeTest struct {
		arg1 Series
	}
	printRangeTests := []printRangeTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a", "red"}},
						{1, []interface{}{"b", "yellow"}},
						{2, []interface{}{"c", "red"}},
					},
					[]string{"ID", "Color"},
				},
				"Fruit",
			},
		},
	}

	for _, test := range printRangeTests {
		test.arg1.PrintRange(1, 3)
	}
}

func TestSeriesHead(t *testing.T) {
	type headTest struct {
		arg1 Series
	}
	headTests := []headTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{"Index"},
				},
				"Fruit",
			},
		},
	}

	for _, test := range headTests {
		test.arg1.Head(2)
	}
}

func TestSeriesTail(t *testing.T) {
	type tailTest struct {
		arg1 Series
	}
	tailTests := []tailTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{"Index"},
				},
				"Fruit",
			},
		},
	}

	for _, test := range tailTests {
		test.arg1.Tail(2)
	}
}

func TestSeriesAt(t *testing.T) {
	type atTest struct {
		arg1     Series
		arg2     []interface{}
		expected interface{}
	}
	atTests := []atTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
			},
			[]interface{}{0},
			"alice",
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
			},
			[]interface{}{"b"},
			"banana",
		},
	}

	for _, test := range atTests {
		output, err := test.arg1.At(test.arg2...)
		if output != test.expected || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestSeriesIAt(t *testing.T) {
	type iatTest struct {
		arg1     Series
		arg2     int
		expected interface{}
	}
	iatTests := []iatTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
			},
			0,
			"alice",
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
			},
			1,
			"banana",
		},
	}

	for _, test := range iatTests {
		output, err := test.arg1.IAt(test.arg2)
		if output != test.expected || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
func TestSeriesLoc(t *testing.T) {
	type locTest struct {
		arg1          Series
		arg2          [][]interface{}
		expected      *Series
		expectedError error
	}
	locTests := []locTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
			},
			[][]interface{}{{0}, {1}},
			&Series{
				[]interface{}{"alice", "bob"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
					},
					[]string{""},
				},
				"People",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
			},
			[][]interface{}{{"b"}, {"c"}},
			&Series{
				[]interface{}{"banana", "cherry"},
				IndexData{
					[]Index{
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[][]interface{}{{"female"}},
			&Series{
				[]interface{}{"clara", "anna"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[][]interface{}{{"male", "volleyball"}},
			&Series{
				[]interface{}{"brian"},
				IndexData{
					[]Index{{1, []interface{}{"male", "volleyball"}}},
					[]string{"sex", "sports"},
				},
				"People",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[][]interface{}{{"male"}, {"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[][]interface{}{{"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[][]interface{}{{"female"}, {"male"}},
			&Series{
				[]interface{}{"clara", "anna", "brian", "dorian", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[][]interface{}{{"female"}, {"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
	}

	for _, test := range locTests {
		output, err := test.arg1.Loc(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestSeriesILoc(t *testing.T) {
	type ilocTest struct {
		arg1     Series
		arg2     int
		arg3     int
		expected []interface{}
	}
	ilocTests := []ilocTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
			},
			0,
			2,
			[]interface{}{"alice", "bob"},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
			},
			0,
			1,
			[]interface{}{"apple"},
		},
	}

	for _, test := range ilocTests {
		output, err := test.arg1.ILoc(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestCount(t *testing.T) {
	type countTest struct {
		arg1     Series
		expected int
	}
	countTests := []countTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			4,
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			3,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			0,
		},
	}

	for _, test := range countTests {
		output := test.arg1.Count()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestMean(t *testing.T) {
	type meanTest struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	meanTests := []meanTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not float64: Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			24.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
	}

	for _, test := range meanTests {
		output, err := test.arg1.Mean()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestMedian(t *testing.T) {
	type medianTest struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	medianTests := []medianTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			23.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			175.85,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			178.7,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			171.5,
			nil,
		},
	}

	for _, test := range medianTests {
		output, err := test.arg1.Median()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestStd(t *testing.T) {
	type stdTest struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	stdTests := []stdTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			5.568,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			7.913,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			9.601,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			10.182,
			nil,
		},
	}

	for _, test := range stdTests {
		output, err := test.arg1.Std()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestMin(t *testing.T) {
	type minTest struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	minTests := []minTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not a float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			19.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			164.3,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			164.3,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			164.3,
			nil,
		},
	}

	for _, test := range minTests {
		output, err := test.arg1.Min()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestMax(t *testing.T) {
	type maxTest struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	maxTests := []maxTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not a float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			30.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			182.5,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			182.5,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			178.7,
			nil,
		},
	}

	for _, test := range maxTests {
		output, err := test.arg1.Max()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestQ1(t *testing.T) {
	type q1Test struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	q1Tests := []q1Test{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not a float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			19.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			168.65,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			164.3,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			164.3,
			nil,
		},
	}

	for _, test := range q1Tests {
		output, err := test.arg1.Q1()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestQ2(t *testing.T) {
	type q2Test struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	q2Tests := []q2Test{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			23.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			175.85,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			178.7,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			171.5,
			nil,
		},
	}

	for _, test := range q2Tests {
		output, err := test.arg1.Q2()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestQ3(t *testing.T) {
	type q3Test struct {
		arg1          Series
		expected      float64
		expectedError error
	}
	q3Tests := []q3Test{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
			},
			math.NaN(),
			fmt.Errorf("data is not float64: %v", "Avery"),
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
			},
			30.0,
			nil,
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
			},
			math.NaN(),
			fmt.Errorf("no elements in this column"),
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			180.6,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
			},
			182.5,
			nil,
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
			},
			178.7,
			nil,
		},
	}

	for _, test := range q3Tests {
		output, err := test.arg1.Q3()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			if fmt.Sprint(output) == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
				}
			} else {
				t.Fatalf("expected %v, got %v, err %v", test.expected, output, err)
			}
		}
	}
}

func TestDescribe(t *testing.T) {
	type describeTest struct {
		arg1     Series
		expected []interface{}
	}
	describeTests := []describeTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
			nil,
		},
		{
			Series{
				[]interface{}{123.123, 456.456, 789.789},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
			[]interface{}{3, 456.456, 456.456, 333.333, 123.123, 789.789, 123.123, 456.456, 789.789},
		},
	}
	for _, test := range describeTests {
		output, err := test.arg1.Describe()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestValueCounts(t *testing.T) {
	type valueCountsTest struct {
		arg1     Series
		expected *Series
	}
	valueCountsTests := []valueCountsTest{
		{
			Series{
				[]interface{}{"Amazon", "Amazon", "Google", "Apple", "Apple", "Apple", "Facebook"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
						{4, []interface{}{4}},
						{5, []interface{}{5}},
						{6, []interface{}{6}},
					},
					[]string{"RangeIndex"},
				},
				"Workplaces",
			},
			&Series{
				[]interface{}{2, 3, 1, 1},
				IndexData{
					[]Index{
						{0, []interface{}{"Amazon"}},
						{1, []interface{}{"Apple"}},
						{2, []interface{}{"Facebook"}},
						{3, []interface{}{"Google"}},
					},
					[]string{"Data"},
				},
				"Unique Value Count of Workplaces",
			},
		},
	}
	for _, test := range valueCountsTests {
		output, err := test.arg1.ValueCounts()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{}), cmpopts.IgnoreFields(Index{}, "id")) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestSeriesRenameCol(t *testing.T) {
	type renameColTest struct {
		arg1     Series
		arg2     string
		expected Series
	}
	renameColTests := []renameColTest{
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
			),
			"Names",
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Names",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
			),
		},
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
			"Names",
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Names",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
		},
	}

	for _, test := range renameColTests {
		test.arg1.RenameCol(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func TestSeriesRenameIndex(t *testing.T) {
	type renameIndexTest struct {
		arg1     Series
		arg2     map[string]string
		expected Series
	}
	renameIndexTests := []renameIndexTest{
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"}},
			),
			map[string]string{"Name": "Names"},
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Names"},
				},
			),
		},
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
			map[string]string{"Name": "Names"},
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Names", "Age"},
				},
			),
		},
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
			map[string]string{"Names": "Name"},
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return *newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
		},
	}

	for _, test := range renameIndexTests {
		err := test.arg1.RenameIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || err != nil {
			if fmt.Sprint(err)[:21] == "index does not exist:" {
				continue
			}
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func TestSeriesSortByIndex(t *testing.T) {
	type sortByIndexTest struct {
		arg1     Series
		arg2     bool
		expected Series
	}
	sortByIndexTests := []sortByIndexTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col2",
			},
			true,
			Series{
				[]interface{}{"d", "a", "c", "b"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col2",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col3",
			},
			false,
			Series{
				[]interface{}{"b", "c", "a", "d"},
				IndexData{
					[]Index{
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{1, []interface{}{1}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col3",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", math.NaN()}},
						{1, []interface{}{"Female", 34}},
						{2, []interface{}{"NaN", 40}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"Emily", "Gina", "Michael", "William", "Chris", "Alice"},
				IndexData{
					[]Index{
						{1, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", math.NaN()}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
						{2, []interface{}{"NaN", 40}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", math.NaN()}},
						{1, []interface{}{"Female", 34}},
						{2, []interface{}{"NaN", 40}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 19}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"Emily", "Gina", "Michael", "William", "Chris", "Alice"},
				IndexData{
					[]Index{
						{1, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", math.NaN()}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 19}},
						{5, []interface{}{"Male", 50}},
						{2, []interface{}{"NaN", 40}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
	}
	for _, test := range sortByIndexTests {
		test.arg1.SortByIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func TestSeriesSortByGivenIndex(t *testing.T) {
	type sortByGivenIndexTest struct {
		arg1     Series
		arg2     IndexData
		expected Series
	}
	sortByGivenIndexTests := []sortByGivenIndexTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
			IndexData{
				[]Index{{0, []interface{}{0}}, {3, []interface{}{3}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
				[]string{""},
			},
			Series{
				[]interface{}{"a", "d", "b", "c"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{3, []interface{}{3}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
			},
			IndexData{
				[]Index{{1, []interface{}{1}}, {3, []interface{}{3}}, {2, []interface{}{2}}, {0, []interface{}{0}}},
				[]string{""},
			},
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
			},
			IndexData{
				[]Index{{3, []interface{}{3}}, {2, []interface{}{2}}, {1, []interface{}{1}}, {0, []interface{}{0}}},
				[]string{""},
			},
			Series{
				[]interface{}{"b", "c", "a", "d"},
				IndexData{
					[]Index{
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{1, []interface{}{1}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			IndexData{
				[]Index{{3, []interface{}{"Female", 16}}, {4, []interface{}{"Female", 34}}, {0, []interface{}{"Female", 40}}, {1, []interface{}{"Male", 19}}, {2, []interface{}{"Male", 25}}, {5, []interface{}{"Male", 50}}},
				[]string{"Sex", "Age"},
			},
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
	}

	for _, test := range sortByGivenIndexTests {
		test.arg1.SortByGivenIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func TestSeriesSortByValues(t *testing.T) {
	type sortByValuesTest struct {
		arg1     Series
		arg2     bool
		expected Series
	}
	sortByValuesTests := []sortByValuesTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"d", "a", "c", "b"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
			},
			false,
			Series{
				[]interface{}{"d", "c", "b", "a"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
						{1, []interface{}{1}},
					},
					[]string{""},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"Alice", "Chris", "Emily", "Gina", "Michael", "William"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{5, []interface{}{"Male", 50}},
						{4, []interface{}{"Female", 34}},
						{3, []interface{}{"Female", 16}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "NaN", "William", "NaN", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{"Alice", "Chris", "Emily", "William", "NaN", "NaN"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{5, []interface{}{"Male", 50}},
						{4, []interface{}{"Female", 34}},
						{2, []interface{}{"Male", 25}},
						{1, []interface{}{"Male", 19}},
						{3, []interface{}{"Female", 16}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
		{
			Series{
				[]interface{}{64.0, math.NaN(), 66.1, 57.24, 96.432, math.NaN()},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
			true,
			Series{
				[]interface{}{57.24, 64.0, 66.1, 96.432, math.NaN(), math.NaN()},
				IndexData{
					[]Index{
						{3, []interface{}{"Female", 16}},
						{0, []interface{}{"Female", 40}},
						{2, []interface{}{"Male", 25}},
						{4, []interface{}{"Female", 34}},
						{1, []interface{}{"Male", 19}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
			},
		},
	}
	for _, test := range sortByValuesTests {
		test.arg1.SortByValues(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}
