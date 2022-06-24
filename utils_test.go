package gambas

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func BenchmarkCheckTypeIntegrity(b *testing.B) {
	list := make([]interface{}, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, rand.Float64())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		checkTypeIntegrity(list)
	}
}

func TestCheckTypeIntegrity(t *testing.T) {
	type checkTypeIntegrityTest struct {
		arg1     []interface{}
		expected string
	}
	checkTypeIntegrityTests := []checkTypeIntegrityTest{
		{
			[]interface{}{0, 1, 2, 3, 4, 5, 6},
			"int",
		},
		{
			[]interface{}{3.4, 5.6, 2.4, 6.5, 7},
			"float64",
		},
		{
			[]interface{}{1, 2, "3", "4", 5},
			"string",
		},
		{
			[]interface{}{"a", "b", "c"},
			"string",
		},
		{
			[]interface{}{true, false, true},
			"bool",
		},
		{
			[]interface{}{"", 1, 2, 3},
			"float64",
		},
		{
			[]interface{}{1.0, "", 2.0},
			"float64",
		},
		{
			[]interface{}{1.0, "", "2.0"},
			"string",
		},
	}

	for _, test := range checkTypeIntegrityTests {
		output, err := checkTypeIntegrity(test.arg1)
		if !cmp.Equal(output, test.expected) || (output != test.expected && err == nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkI2f(b *testing.B) {
	list := make([]interface{}, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, rand.Float64())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		i2f(list)
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

func BenchmarkTryBool(b *testing.B) {
	wordList := []string{"TRUE", "True", "true", "t", "1", "FALSE", "False", "false", "f", "0"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tryBool(wordList[rand.Intn(10)])
	}
}

func TestTryBool(t *testing.T) {
	type tryBoolTest struct {
		arg1          string
		expected      bool
		expectedError error
	}
	tryBoolTests := []tryBoolTest{
		{
			"true",
			true,
			nil,
		},
		{
			"True",
			true,
			nil,
		},
		{
			"TRUE",
			true,
			nil,
		},
		{
			"false",
			false,
			nil,
		},
		{
			"False",
			false,
			nil,
		},
		{
			"FALSE",
			false,
			nil,
		},
		{
			"1",
			false,
			fmt.Errorf("ignored string"),
		},
		{
			"t",
			false,
			fmt.Errorf("ignored string"),
		},
		{
			"T",
			false,
			fmt.Errorf("ignored string"),
		},
		{
			"0",
			false,
			fmt.Errorf("ignored string"),
		},
		{
			"f",
			false,
			fmt.Errorf("ignored string"),
		},
		{
			"F",
			false,
			fmt.Errorf("ignored string"),
		},
	}
	for _, test := range tryBoolTests {
		output, err := tryBool(test.arg1)
		if !cmp.Equal(output, test.expected) || !cmp.Equal(fmt.Sprint(err), fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkTryInt(b *testing.B) {
	wordList := []string{"0", "0.0", "-1", "-1.0", "1", "1.0", "1-1", "1/2"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tryInt(wordList[rand.Intn(8)])
	}
}

func TestTryInt(t *testing.T) {
	type tryIntTest struct {
		arg1          string
		expected      int
		expectedError error
	}
	tryIntTests := []tryIntTest{
		{
			"0",
			0,
			nil,
		},
		{
			"1",
			1,
			nil,
		},
		{
			"-1",
			-1,
			nil,
		},
		{
			"0.0",
			0,
			fmt.Errorf(`strconv.Atoi: parsing "0.0": invalid syntax`),
		},
		{
			"1-1",
			0,
			fmt.Errorf(`strconv.Atoi: parsing "1-1": invalid syntax`),
		},
		{
			"1.0",
			0,
			fmt.Errorf(`strconv.Atoi: parsing "1.0": invalid syntax`),
		},
		{
			"1.1",
			0,
			fmt.Errorf(`strconv.Atoi: parsing "1.1": invalid syntax`),
		},
	}
	for _, test := range tryIntTests {
		output, err := tryInt(test.arg1)
		if !cmp.Equal(output, test.expected) || !cmp.Equal(fmt.Sprint(err), fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkTryFloat64(b *testing.B) {
	wordList := []string{"0", "0.0", "-1", "-1.0", "1", "1.0", "1-1", "1/2"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tryFloat64(wordList[rand.Intn(8)])
	}
}

func TestTryFloat64(t *testing.T) {
	type tryFloat64Test struct {
		arg1          string
		expected      float64
		expectedError error
	}
	tryFloat64Tests := []tryFloat64Test{
		{
			"0.0",
			0.0,
			nil,
		},
		{
			"1.0",
			1.0,
			nil,
		},
		{
			"-1.0",
			-1.0,
			nil,
		},
		{
			"0.5",
			0.5,
			nil,
		},
		{
			"1-1",
			0.0,
			fmt.Errorf(`strconv.ParseFloat: parsing "1-1": invalid syntax`),
		},
		{
			"1",
			1.0,
			nil,
		},
		{
			"0",
			0.0,
			nil,
		},
		{
			"-1",
			-1.0,
			nil,
		},
	}
	for _, test := range tryFloat64Tests {
		output, err := tryFloat64(test.arg1)
		if !cmp.Equal(output, test.expected) || !cmp.Equal(fmt.Sprint(err), fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkTryDataType(b *testing.B) {
	wordList := []string{
		"TRUE", "True", "true", "t", "1", "FALSE", "False", "false", "f", "0", "0", "0.0", "-1", "-1.0", "1", "1.0", "1-1", "1/2",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tryDataType(wordList[rand.Intn(18)])
	}
}

func TestTryDataType(t *testing.T) {
	type tryDataTypeTest struct {
		arg1     string
		expected interface{}
	}
	tryDataTypeTests := []tryDataTypeTest{
		{
			"true",
			true,
		},
		{
			"True",
			true,
		},
		{
			"TRUE",
			true,
		},
		{
			"false",
			false,
		},
		{
			"False",
			false,
		},
		{
			"FALSE",
			false,
		},
		{
			"1",
			1,
		},
		{
			"t",
			"t",
		},
		{
			"T",
			"T",
		},
		{
			"0",
			0,
		},
		{
			"f",
			"f",
		},
		{
			"F",
			"F",
		},
		{
			"0.0",
			0.0,
		},
		{
			"1.0",
			1.0,
		},
		{
			"-1.0",
			-1.0,
		},
		{
			"0.5",
			0.5,
		},
		{
			"1-1",
			"1-1",
		},
		{
			"1",
			1,
		},
		{
			"0",
			0,
		},
		{
			"-1",
			-1,
		},
		{
			"hello",
			"hello",
		},
		{
			"5/2",
			"5/2",
		},
	}
	for _, test := range tryDataTypeTests {
		output := tryDataType(test.arg1)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkCheckCSVDataType(b *testing.B) {
	wordList := []string{
		"TRUE", "True", "true", "t", "1", "FALSE", "False", "false", "f", "0", "0", "0.0", "-1", "-1.0", "1", "1.0", "1-1", "1/2",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		checkCSVDataType(wordList[rand.Intn(18)])
	}
}

func TestCheckCSVDataType(t *testing.T) {
	type checkCSVDataTypeTest struct {
		arg1     string
		expected interface{}
	}
	checkCSVDataTypeTests := []checkCSVDataTypeTest{
		{"50.0", 50.0},
		{"abcd", "abcd"},
		{"100", 100.0},
		{"NaN", math.NaN()},
	}

	for _, test := range checkCSVDataTypeTests {
		output := checkCSVDataType(test.arg1)
		if !cmp.Equal(output, test.expected) {
			if test.arg1 == "NaN" {
				if !cmp.Equal(fmt.Sprint(output), fmt.Sprint(test.expected)) {
					t.Fatalf("expected %v, got %v", test.expected, output)
				}
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkConsolidateToFloat64(b *testing.B) {
	list := []interface{}{
		0, 0.0, 1, 1.0, 2, 2.0,
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		consolidateToFloat64(list)
	}
}

func TestConsolidateToFloat64(t *testing.T) {
	type consolidateToFloat64Test struct {
		arg1     []interface{}
		expected []interface{}
	}
	consolidateToFloat64Tests := []consolidateToFloat64Test{
		{
			[]interface{}{1.0, 2.0, 3.0},
			[]interface{}{1.0, 2.0, 3.0},
		},
		{
			[]interface{}{"", 2.0, 3.0},
			[]interface{}{math.NaN(), 2.0, 3.0},
		},
		{
			[]interface{}{"", ""},
			[]interface{}{math.NaN(), math.NaN()},
		},
	}
	for _, test := range consolidateToFloat64Tests {
		output := consolidateToFloat64(test.arg1)
		if !cmp.Equal(output, test.expected, cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkConsolidateToString(b *testing.B) {
	list := []interface{}{
		0, 0.0, 1, 1.0, 2, 2.0,
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		consolidateToString(list)
	}
}

func TestConsolidateToString(t *testing.T) {
	type consolidateToStringTest struct {
		arg1     []interface{}
		expected []interface{}
	}
	consolidateToStringTests := []consolidateToStringTest{
		{
			[]interface{}{1.0, 2.0, 3.0},
			[]interface{}{"1", "2", "3"},
		},
		{
			[]interface{}{"", 2.0, 3.0},
			[]interface{}{"", "2", "3"},
		},
		{
			[]interface{}{"", ""},
			[]interface{}{"", ""},
		},
		{
			[]interface{}{true, "false"},
			[]interface{}{"true", "false"},
		},
		{
			[]interface{}{1, 2, 3, 4, ""},
			[]interface{}{"1", "2", "3", "4", ""},
		},
	}
	for _, test := range consolidateToStringTests {
		output := consolidateToString(test.arg1)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkInterface2F64Slice(b *testing.B) {
	list := make([]interface{}, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, rand.Float64())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		interface2F64Slice(list)
	}
}

func TestInterface2F64Slice(t *testing.T) {
	type interface2F64DataTest struct {
		arg1     []interface{}
		expected []float64
	}
	interface2F64DataTests := []interface2F64DataTest{
		{
			[]interface{}{0.0, 1.0, 2.0, 3.0, 4.0},
			[]float64{0.0, 1.0, 2.0, 3.0, 4.0},
		},
		{
			[]interface{}{0, 1, 2, 3, 4},
			nil,
		},
		{
			[]interface{}{"0", "1", "2", "3", "4"},
			nil,
		},
		{
			[]interface{}{"a", "b", "c", "d", "e"},
			nil,
		},
		{
			[]interface{}{0.5, 1, 1.5, 2, 2.5},
			nil,
		},
		{
			[]interface{}{"a", 1, "b", 2, "c", 3},
			nil,
		},
	}
	for _, test := range interface2F64DataTests {
		output, err := interface2F64Slice(test.arg1)
		if !cmp.Equal(output, test.expected) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkInterface2StringSlice(b *testing.B) {
	list := make([]interface{}, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		interface2StringSlice(list)
	}
}

func TestInterface2StringSlice(t *testing.T) {
	type interface2StringDataTest struct {
		arg1     []interface{}
		expected []string
	}
	interface2StringDataTests := []interface2StringDataTest{
		{
			[]interface{}{0.0, 1.0, 2.0, 3.0, 4.0},
			nil,
		},
		{
			[]interface{}{0, 1, 2, 3, 4},
			nil,
		},
		{
			[]interface{}{"0", "1", "2", "3", "4"},
			[]string{"0", "1", "2", "3", "4"},
		},
		{
			[]interface{}{"a", "b", "c", "d", "e"},
			[]string{"a", "b", "c", "d", "e"},
		},
		{
			[]interface{}{0.5, 1, 1.5, 2, 2.5},
			nil,
		},
		{
			[]interface{}{"a", 1, "b", 2, "c", 3},
			nil,
		},
	}
	for _, test := range interface2StringDataTests {
		output, err := interface2StringSlice(test.arg1)
		if !cmp.Equal(output, test.expected) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSlicesAreEqual(b *testing.B) {
	list1 := make([]interface{}, 0)
	for i := 0; i < 10000; i++ {
		list1 = append(list1, rand.Float64())
	}
	list2 := make([]interface{}, 0)
	for i := 0; i < 10000; i++ {
		list2 = append(list2, rand.Float64())
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slicesAreEqual(list1, list2)
	}
}

func TestSlicesAreEqual(t *testing.T) {
	type slicesAreEqualTest struct {
		arg1     []interface{}
		arg2     []interface{}
		expected bool
	}
	slicesAreEqualTests := []slicesAreEqualTest{
		{
			[]interface{}{0, 1, 2, 3, 4},
			[]interface{}{0, 1, 2, 3, 4},
			true,
		},
		{
			[]interface{}{"a", "b"},
			[]interface{}{"a", "b", "c"},
			false,
		},
		{
			[]interface{}{true, true, false},
			[]interface{}{true, false, false},
			false,
		},
	}
	for _, test := range slicesAreEqualTests {
		output := slicesAreEqual(test.arg1, test.arg2)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkContainsString(b *testing.B) {
	list := make([]string, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		containsString(list, strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	}
}

func TestContainsString(t *testing.T) {
	type containsStringTest struct {
		arg1     []string
		arg2     string
		expected bool
	}
	containsStringTests := []containsStringTest{
		{
			[]string{"a", "b", "c", "d", "e"},
			"a",
			true,
		},
		{
			[]string{"a", "b", "c", "d", "e"},
			"f",
			false,
		},
	}
	for _, test := range containsStringTests {
		output := containsString(test.arg1, test.arg2)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkContainsIndex(b *testing.B) {
	list := make([]Index, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, Index{i, []interface{}{rand.Float64()}})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		containsIndex(list, Index{rand.Intn(10000), []interface{}{rand.Float64()}})
	}
}

func TestContainsIndex(t *testing.T) {
	type containsIndexTest struct {
		arg1     []Index
		arg2     Index
		expected bool
	}
	containsIndexTests := []containsIndexTest{
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{0, []interface{}{"index0"}},
			true,
		},
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{0, []interface{}{"index1"}},
			false,
		},
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{1, []interface{}{"index0"}},
			false,
		},
	}
	for _, test := range containsIndexTests {
		output := containsIndex(test.arg1, test.arg2)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkContainsIndexWithoutId(b *testing.B) {
	list := make([]Index, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, Index{i, []interface{}{rand.Float64()}})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		containsIndex(list, Index{rand.Intn(10000), []interface{}{rand.Float64()}})
	}
}

func TestContainsIndexWithoutId(t *testing.T) {
	type containsIndexWithoutIdTest struct {
		arg1     []Index
		arg2     Index
		expected bool
	}
	containsIndexWithoutIdTests := []containsIndexWithoutIdTest{
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{0, []interface{}{"index0"}},
			true,
		},
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{0, []interface{}{"index1"}},
			true,
		},
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{1, []interface{}{"index0"}},
			true,
		},
		{
			[]Index{
				{0, []interface{}{"index0"}},
				{1, []interface{}{"index1"}},
				{2, []interface{}{"index2"}},
			},
			Index{1, []interface{}{"index3"}},
			false,
		},
	}
	for _, test := range containsIndexWithoutIdTests {
		output := containsIndexWithoutId(test.arg1, test.arg2)
		if !cmp.Equal(output, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
