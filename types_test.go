package gambas

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHashKey(t *testing.T) {
	type hashKeyTest struct {
		arg1     Index
		expected *string
	}
	hashKeyTests := []hashKeyTest{
		{
			Index{0, 1, 2, 3},
			func(s string) *string {
				return &s
			}("26e5556d35bebab660c17eaf2ffebb823ddfb4b009a6e1f687e95fdf40ee637040fe258bb737d537c5d04d2558f2d2e802b2ef492944330f0b8e4bb17222e1d5"),
		},
		{
			Index{"a", "b", "c", "d"},
			func(s string) *string {
				return &s
			}("d8022f2060ad6efd297ab73dcc5355c9b214054b0d1776a136a669d26a7d3b14f73aa0d0ebff19ee333368f0164b6419a96da49e3e481753e7e96b716bdccb6f"),
		},
		{
			Index{"name", "age"},
			func(s string) *string {
				return &s
			}("00ea90c8c8aa40e1aed13c38b1f3885c30abe1bf26f9691d4d5b28f919742cd11e75ab56015dda37c4ff75faf6351f2c7b08cbeb7a0fc84e16dca60a9e2cbde9"),
		},
		{
			Index{""},
			func(s string) *string {
				return &s
			}("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"),
		},
		{
			Index{},
			nil,
		},
	}
	for _, test := range hashKeyTests {
		output, err := test.arg1.hashKey()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected: %s, got: %s, err: %s", *test.expected, *output, err)
		}
	}
}

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
			5.5677643628300215,
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
			7.913437938089859,
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
			9.600694419328905,
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
			10.182337649086268,
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

func TestDataFrameLocRows(t *testing.T) {
	type dataframeLocRowsTest struct {
		arg1     DataFrame
		arg2     []Index
		expected *DataFrame
	}
	dataframeLocRowsTests := []dataframeLocRowsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]Index{{"Avery"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery"},
						IndexData{
							[]Index{{"Avery"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19},
						IndexData{
							[]Index{{"Avery"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male"},
						IndexData{
							[]Index{{"Avery"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}

	for _, test := range dataframeLocRowsTests {
		output, err := test.arg1.LocRows(test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLocCols(t *testing.T) {
	type dataframeLocColsTest struct {
		arg1     DataFrame
		arg2     []string
		expected *DataFrame
	}
	dataframeLocColsTests := []dataframeLocColsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{19, 27, 22},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
	}
	for _, test := range dataframeLocColsTests {
		output, err := test.arg1.LocCols(test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLoc(t *testing.T) {
	type dataframeLocTest struct {
		arg1     DataFrame
		arg2     []Index
		arg3     []string
		expected *DataFrame
	}
	dataframeLocTests := []dataframeLocTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]Index{{"Bradley"}},
			[]string{"Age"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{27},
						IndexData{
							[]Index{{"Bradley"}},
							[]string{"Name"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{{"Bradley"}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
	}
	for _, test := range dataframeLocTests {
		output, err := test.arg1.Loc(test.arg2, test.arg3)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColAdd(t *testing.T) {
	type colAddTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colAddTests := []colAddTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{24.0, 32.0, 27.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colAddTests {
		output, err := test.arg1.ColAdd(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColSub(t *testing.T) {
	type colSubTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colSubTests := []colSubTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{14.0, 22.0, 17.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colSubTests {
		output, err := test.arg1.ColSub(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

// func TestColMul(t *testing.T) {
// 	type colMulTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     float64
// 		expected *DataFrame
// 	}

// 	colMulTests := []colMulTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Name",
// 			5.0,
// 			nil,
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19.0, 25.0, 22.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age",
// 			5.0,
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{95.0, 125.0, 110.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"RandomName",
// 			5.0,
// 			nil,
// 		},
// 	}
// 	for _, test := range colMulTests {
// 		output, err := test.arg1.ColMul(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || (output != nil && err != nil) {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }

// func TestColDiv(t *testing.T) {
// 	type colDivTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     float64
// 		expected *DataFrame
// 	}

// 	colDivTests := []colDivTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Name",
// 			5.0,
// 			nil,
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19.0, 25.0, 22.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age",
// 			5.0,
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{3.8, 5.0, 4.4},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"RandomName",
// 			5.0,
// 			nil,
// 		},
// 	}
// 	for _, test := range colDivTests {
// 		output, err := test.arg1.ColDiv(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || (output != nil && err != nil) {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }

// func TestColMod(t *testing.T) {
// 	type colModTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     float64
// 		expected *DataFrame
// 	}

// 	colModTests := []colModTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Name",
// 			5.0,
// 			nil,
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19.0, 25.0, 22.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age",
// 			5.0,
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{4.0, 0.0, 2.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"RandomName",
// 			5.0,
// 			nil,
// 		},
// 	}
// 	for _, test := range colModTests {
// 		output, err := test.arg1.ColMod(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || (output != nil && err != nil) {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }

// func TestColGt(t *testing.T) {
// 	type colGtTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     float64
// 		expected *DataFrame
// 	}

// 	colGtTests := []colGtTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Name",
// 			5.0,
// 			nil,
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19.0, 25.0, 22.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age",
// 			20.0,
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{false, true, true},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"RandomName",
// 			5.0,
// 			nil,
// 		},
// 	}
// 	for _, test := range colGtTests {
// 		output, err := test.arg1.ColGt(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || (output != nil && err != nil) {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }

// func TestColLt(t *testing.T) {
// 	type colLtTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     float64
// 		expected *DataFrame
// 	}

// 	colLtTests := []colLtTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Name",
// 			5.0,
// 			nil,
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19.0, 25.0, 22.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age",
// 			20.0,
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{true, false, false},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"RandomName",
// 			5.0,
// 			nil,
// 		},
// 	}
// 	for _, test := range colLtTests {
// 		output, err := test.arg1.ColLt(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || (output != nil && err != nil) {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }

// func TestColEq(t *testing.T) {
// 	type colEqTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     float64
// 		expected *DataFrame
// 	}

// 	colEqTests := []colEqTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Name",
// 			5.0,
// 			nil,
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19.0, 25.0, 22.0},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age",
// 			25.0,
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{false, true, false},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"RandomName",
// 			5.0,
// 			nil,
// 		},
// 	}
// 	for _, test := range colEqTests {
// 		output, err := test.arg1.ColEq(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || (output != nil && err != nil) {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }

// func TestNewCol(t *testing.T) {
// 	type newColTest struct {
// 		arg1     *DataFrame
// 		arg2     string
// 		arg3     []interface{}
// 		expected *DataFrame
// 	}
// 	newColTests := []newColTest{
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Nationality",
// 			[]interface{}{"USA", "UK", "Canada"},
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 					{
// 						[]interface{}{"USA", "UK", "Canada"},
// 						Index{0, 1, 2},
// 						"Nationality",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex", "Nationality"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 		{
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 			"Age+5",
// 			make([]interface{}, 3),
// 			&DataFrame{
// 				[]Series{
// 					{
// 						[]interface{}{"Avery", "Bradford", "Candice"},
// 						Index{0, 1, 2},
// 						"Name",
// 					},
// 					{
// 						[]interface{}{19, 25, 22},
// 						Index{0, 1, 2},
// 						"Age",
// 					},
// 					{
// 						[]interface{}{"Male", "Male", "Female"},
// 						Index{0, 1, 2},
// 						"Sex",
// 					},
// 					{
// 						[]interface{}{nil, nil, nil},
// 						Index{0, 1, 2},
// 						"Age+5",
// 					},
// 				},
// 				Index{"Name", "Age", "Sex", "Age+5"},
// 				[]interface{}{"Name"},
// 				[]Index{{"Avery", "Bradford", "Candice"}},
// 			},
// 		},
// 	}
// 	for _, test := range newColTests {
// 		output, err := test.arg1.NewCol(test.arg2, test.arg3)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{})) || err != nil {
// 			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
// 		}
// 	}
// }
