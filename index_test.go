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
			Index{
				0,
				[]interface{}{0, 1, 2, 3},
			},
			func(s string) *string {
				return &s
			}("8b45172ee9ab0d5eeecc3410f566e4b1f257da5719d4f5f3ecfb8ba24af8a592d009dcb539e4b81657bb3150005dd3b2b2f69a1132858e5e7b8a7d58cf6d4d39"),
		},
		{
			Index{
				0,
				[]interface{}{"a", "b", "c", "d"},
			},
			func(s string) *string {
				return &s
			}("2b049388bd0f4a2f7d0415b4e2c3cac0cd9f4eb3c46178cf0fc7ec85c0fba56ec579d33e9acd8b7e37f5643652a85efe3bdc8e19716c436650e8bc5b6dee5282"),
		},
		{
			Index{
				0,
				[]interface{}{"name", "age"},
			},
			func(s string) *string {
				return &s
			}("2ac639afcf5c9da0f6a75e184ce4a4f978a26dff1ffbd26dc5ff1a4ed5524250c89bfffdc99d375539a4f961cd62b39527bdda7f76ad087ea54bb9a3cae0e8f6"),
		},
		{
			Index{
				0,
				[]interface{}{""},
			},
			func(s string) *string {
				return &s
			}("31bca02094eb78126a517b206a88c73cfa9ec6f704c7030d18212cace820f025f00bf0ea68dbf3f3a5436ca63b53bf7bf80ad8d5de7d8359d0b7fed9dbc3ab99"),
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
