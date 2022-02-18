package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSeries(t *testing.T) {
	data := []interface{}{"alice", "bob", "charlie"}
	s, err := NewSeries(data, CustomIndex{[]interface{}{"a", "b", "c"}}, "People")

	expected := Series{
		Data: map[interface{}]interface{}{"a": "alice", "b": "bob", "c": "charlie"},
		CustomIndex: CustomIndex{
			Value: []interface{}{"a", "b", "c"},
		},
		Name: "People",
	}

	if !cmp.Equal(s, expected) || err != nil {
		t.Fatalf("\nexpected %v,\ngot %v,\nerror: %v", expected, s, err)
	}
}
