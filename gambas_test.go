package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSeries(t *testing.T) {
	data := []interface{}{"alice", "bob", "charlie"}
	s, err := NewSeries(data, nil, "People")

	expected := Series{
		Data: map[Index]interface{}{{0}: "alice", {1}: "bob", {2}: "charlie"},
		Name: "People",
	}

	if !cmp.Equal(s, expected) || err != nil {
		t.Fatalf("\nexpected %v,\ngot %v,\nerror: %v", expected, s, err)
	}
}
