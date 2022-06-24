package gambas

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func BenchmarkIndexHashKey(b *testing.B) {
	list := make([]Index, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, Index{i, []interface{}{rand.Intn(100)}})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list[rand.Intn(10000)].hashKey()
	}
}

func TestIndexHashKey(t *testing.T) {
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

func BenchmarkIndexHashKeyValueOnly(b *testing.B) {
	list := make([]Index, 0)
	for i := 0; i < 10000; i++ {
		list = append(list, Index{i, []interface{}{rand.Intn(100)}})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list[rand.Intn(10000)].hashKeyValueOnly()
	}
}

func TestIndexHashKeyValueOnly(t *testing.T) {
	type hashKeyValueOnlyTest struct {
		arg1     Index
		expected *string
	}
	hashKeyValueOnlyTests := []hashKeyValueOnlyTest{
		{
			Index{
				0,
				[]interface{}{0, 1, 2, 3},
			},
			func(s string) *string {
				return &s
			}("26e5556d35bebab660c17eaf2ffebb823ddfb4b009a6e1f687e95fdf40ee637040fe258bb737d537c5d04d2558f2d2e802b2ef492944330f0b8e4bb17222e1d5"),
		},
		{
			Index{
				0,
				[]interface{}{"a", "b", "c", "d"},
			},
			func(s string) *string {
				return &s
			}("d8022f2060ad6efd297ab73dcc5355c9b214054b0d1776a136a669d26a7d3b14f73aa0d0ebff19ee333368f0164b6419a96da49e3e481753e7e96b716bdccb6f"),
		},
		{
			Index{
				0,
				[]interface{}{"name", "age"},
			},
			func(s string) *string {
				return &s
			}("00ea90c8c8aa40e1aed13c38b1f3885c30abe1bf26f9691d4d5b28f919742cd11e75ab56015dda37c4ff75faf6351f2c7b08cbeb7a0fc84e16dca60a9e2cbde9"),
		},
		{
			Index{
				0,
				[]interface{}{""},
			},
			func(s string) *string {
				return &s
			}("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"),
		},
		{
			Index{},
			nil,
		},
	}
	for _, test := range hashKeyValueOnlyTests {
		output, err := test.arg1.hashKeyValueOnly()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected: %s, got: %s, err: %s", *test.expected, *output, err)
		}
	}
}
