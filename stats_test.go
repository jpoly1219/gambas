package gambas

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCount(t *testing.T) {
	type countTest struct {
		arg1     []interface{}
		expected StatsResult
	}
	countTests := []countTest{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Count",
				4.0,
				nil,
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Count",
				3.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Count",
				0.0,
				nil,
			},
		},
	}
	for _, test := range countTests {
		output := Count(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors(), cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestMean(t *testing.T) {
	type meanTest struct {
		arg1     []interface{}
		expected StatsResult
	}
	meanTests := []meanTest{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Mean",
				math.NaN(),
				fmt.Errorf("data is not float64: Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Mean",
				24.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Mean",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
	}

	for _, test := range meanTests {
		output := Mean(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestMedian(t *testing.T) {
	type medianTest struct {
		arg1     []interface{}
		expected StatsResult
	}
	medianTests := []medianTest{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Median",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Median",
				23.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Median",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Median",
				175.85,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Median",
				178.7,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Median",
				171.5,
				nil,
			},
		},
	}

	for _, test := range medianTests {
		output := Median(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestStd(t *testing.T) {
	type stdTest struct {
		arg1     []interface{}
		expected StatsResult
	}
	stdTests := []stdTest{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Std",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Std",
				5.568,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Std",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Std",
				7.913,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Std",
				9.601,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Std",
				10.182,
				nil,
			},
		},
	}
	for _, test := range stdTests {
		output := Std(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestMin(t *testing.T) {
	type minTest struct {
		arg1     []interface{}
		expected StatsResult
	}
	minTests := []minTest{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Min",
				math.NaN(),
				fmt.Errorf("data is not a float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Min",
				19.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Min",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Min",
				164.3,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Min",
				164.3,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Min",
				164.3,
				nil,
			},
		},
	}
	for _, test := range minTests {
		output := Min(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestMax(t *testing.T) {
	type maxTest struct {
		arg1     []interface{}
		expected StatsResult
	}
	maxTests := []maxTest{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Max",
				math.NaN(),
				fmt.Errorf("data is not a float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Max",
				30.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Max",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Max",
				182.5,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Max",
				182.5,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Max",
				178.7,
				nil,
			},
		},
	}
	for _, test := range maxTests {
		output := Max(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestQ1(t *testing.T) {
	type q1Test struct {
		arg1     []interface{}
		expected StatsResult
	}
	q1Tests := []q1Test{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Q1",
				math.NaN(),
				fmt.Errorf("data is not a float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Q1",
				19.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Q1",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Q1",
				168.65,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Q1",
				164.3,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Q1",
				164.3,
				nil,
			},
		},
	}
	for _, test := range q1Tests {
		output := Q1(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestQ2(t *testing.T) {
	type q2Test struct {
		arg1     []interface{}
		expected StatsResult
	}
	q2Tests := []q2Test{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Q2",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Q2",
				23.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Q2",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Q2",
				175.85,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Q2",
				178.7,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Q2",
				171.5,
				nil,
			},
		},
	}
	for _, test := range q2Tests {
		output := Q2(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func TestQ3(t *testing.T) {
	type q3Test struct {
		arg1     []interface{}
		expected StatsResult
	}
	q3Tests := []q3Test{
		{
			[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
			StatsResult{
				"Q3",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			[]interface{}{30.0, 23.0, 19.0},
			StatsResult{
				"Q3",
				30.0,
				nil,
			},
		},
		{
			[]interface{}{},
			StatsResult{
				"Q3",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			[]interface{}{164.3, 182.5, 173.0, 178.7},
			StatsResult{
				"Q3",
				180.6,
				nil,
			},
		},
		{
			[]interface{}{164.3, 182.5, math.NaN(), 178.7},
			StatsResult{
				"Q3",
				182.5,
				nil,
			},
		},
		{
			[]interface{}{164.3, math.NaN(), 178.7},
			StatsResult{
				"Q3",
				178.7,
				nil,
			},
		},
	}
	for _, test := range q3Tests {
		output := Q3(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}
