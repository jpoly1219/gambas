package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewSeries(t *testing.T) {
	data := []interface{}{"alice", "bob", "charlie"}
	s := NewSeries(data, []Index{{"a"}, {"b"}, {"c"}}, "People")

	expected := Series{
		Data:       map[Index]interface{}{{"a"}: "alice", {"b"}: "bob", {"c"}: "charlie"},
		IndexArray: []Index{{"a"}, {"b"}, {"c"}},
		Name:       "People",
	}

	if !cmp.Equal(s, expected) {
		t.Fatalf("expected %v, got %v", expected, s)
	}
}
