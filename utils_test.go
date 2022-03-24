package gambas

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
