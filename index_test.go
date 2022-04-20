package gambas

import (
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
