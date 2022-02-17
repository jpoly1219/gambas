package gambas

import "testing"

func TestPrintSeries(t *testing.T) {
	s := Series{
		Data: map[int]interface{}{0: "alice", 1: "bob", 2: "charlie"},
		Name: "People",
	}

	expected := "Data: map[0:alice 1:bob 2:charlie] \nCustomIndex: {[]} \nName: People\n"
	output := s.PrintSeries()

	if output != expected {
		t.Fatalf("\nexpected:\n%s,\ngot:\n%s\n", expected, output)
	}
}
