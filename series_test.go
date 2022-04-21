package gambas

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintSeries(t *testing.T) {
	seriesArray := []Series{
		{
			[]interface{}{"alice", "bob", "charlie"},
			IndexData{
				[]Index{{0}, {1}, {2}},
				[]string{""},
			},
			"People",
		},
		{
			[]interface{}{"apple", "banana", "cherry"},
			IndexData{
				[]Index{{"a"}, {"b"}, {"c"}},
				[]string{""},
			},
			"Fruit",
		},
	}
	expectedArray := []string{
		"data: [alice bob charlie] \nindexArray: {[[0] [1] [2]] []} \nname: People\n",
		"data: [apple banana cherry] \nindexArray: {[[a] [b] [c]] []} \nname: Fruit\n",
	}

	for i, test := range seriesArray {
		output := test.PrintSeries()
		if output != expectedArray[i] {
			t.Fatalf("expected: %s, got: %s", expectedArray[i], output)
		}
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
					[]Index{{0}, {1}, {2}},
					[]string{"RangeIndex"},
				},
				"People",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{{"a"}, {"b"}, {"c"}},
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

func TestSeriesAt(t *testing.T) {
	type atTest struct {
		arg1     Series
		arg2     Index
		expected interface{}
	}
	atTests := []atTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{{0}, {1}, {2}},
					[]string{""},
				},
				"People",
			},
			Index{0},
			"alice",
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{{"a"}, {"b"}, {"c"}},
					[]string{""},
				},
				"Fruit",
			},
			Index{"b"},
			"banana",
		},
	}

	for _, test := range atTests {
		output, err := test.arg1.At(test.arg2)
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{"a"}, {"b"}, {"c"}},
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
		arg2          []Index
		expected      *Series
		expectedError error
	}
	locTests := []locTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{{0}, {1}, {2}},
					[]string{""},
				},
				"People",
			},
			[]Index{{0}, {1}},
			&Series{
				[]interface{}{"alice", "bob"},
				IndexData{
					[]Index{{0}, {1}},
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
					[]Index{{"a"}, {"b"}, {"c"}},
					[]string{""},
				},
				"Fruit",
			},
			[]Index{{"b"}, {"c"}},
			&Series{
				[]interface{}{"banana", "cherry"},
				IndexData{
					[]Index{{"b"}, {"c"}},
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
					[]Index{{"female", "basketball"}, {"male", "volleyball"}, {"male", "basketball"}, {"female", "volleyball"}, {"male", "swimming"}},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[]Index{{"female"}},
			&Series{
				[]interface{}{"clara", "anna"},
				IndexData{
					[]Index{{"female", "basketball"}, {"female", "volleyball"}},
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
					[]Index{{"female", "basketball"}, {"male", "volleyball"}, {"male", "basketball"}, {"female", "volleyball"}, {"male", "swimming"}},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[]Index{{"male", "volleyball"}},
			&Series{
				[]interface{}{"brian"},
				IndexData{
					[]Index{{"male", "volleyball"}},
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
					[]Index{{"female", "basketball"}, {"male", "volleyball"}, {"male", "basketball"}, {"female", "volleyball"}, {"male", "swimming"}},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[]Index{{"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{{"female", "basketball"}, {"male", "volleyball"}, {"male", "basketball"}, {"female", "volleyball"}, {"male", "swimming"}},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[]Index{{"female"}, {"male"}},
			&Series{
				[]interface{}{"clara", "anna", "brian", "dorian", "michael"},
				IndexData{
					[]Index{{"female", "basketball"}, {"female", "volleyball"}, {"male", "volleyball"}, {"male", "basketball"}, {"male", "swimming"}},
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
					[]Index{{"female", "basketball"}, {"male", "volleyball"}, {"male", "basketball"}, {"female", "volleyball"}, {"male", "swimming"}},
					[]string{"sex", "sports"},
				},
				"People",
			},
			[]Index{{"female"}, {"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
	}

	for _, test := range locTests {
		output, err := test.arg1.Loc(test.arg2)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{"a"}, {"b"}, {"c"}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}},
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
					[]Index{{0}, {1}, {2}, {3}},
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
					[]Index{{0}, {1}, {2}, {3}},
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

func TestSeriesSortByIndex(t *testing.T) {
	type sortByIndexTest struct {
		arg1     *Series
		arg2     bool
		expected *Series
	}
	sortByIndexTests := []sortByIndexTest{
		{
			&Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{{0}, {1}, {2}, {3}},
					[]string{""},
				},
				"col1",
			},
			true,
			&Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{{0}, {1}, {2}, {3}},
					[]string{""},
				},
				"col1",
			},
		},
		{
			&Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{{1}, {3}, {2}, {0}},
					[]string{""},
				},
				"col1",
			},
			true,
			&Series{
				[]interface{}{"d", "a", "c", "b"},
				IndexData{
					[]Index{{0}, {1}, {2}, {3}},
					[]string{""},
				},
				"col1",
			},
		},
		{
			&Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{{1}, {3}, {2}, {0}},
					[]string{""},
				},
				"col1",
			},
			false,
			&Series{
				[]interface{}{"b", "c", "a", "d"},
				IndexData{
					[]Index{{3}, {2}, {1}, {0}},
					[]string{""},
				},
				"col1",
			},
		},
	}
	for _, test := range sortByIndexTests {
		test.arg1.SortByIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}
