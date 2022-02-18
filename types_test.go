package gambas

import "testing"

var seriesArray = []Series{
	{
		Data:       map[Index]interface{}{{0}: "alice", {1}: "bob", {2}: "charlie"},
		IndexArray: []Index{{0}, {1}, {2}},
		Name:       "People",
	},
	{
		Data:       map[Index]interface{}{{"a"}: "apple", {"b"}: "banana", {"c"}: "cherry"},
		IndexArray: []Index{{"a"}, {"b"}, {"c"}},
		Name:       "Fruit",
	},
}

func TestPrintSeries(t *testing.T) {
	expectedArray := []string{
		"Data: map[{0}:alice {1}:bob {2}:charlie] \nIndexArray: [{0} {1} {2}] \nName: People\n",
		"Data: map[{a}:apple {b}:banana {c}:cherry] \nIndexArray: [{a} {b} {c}] \nName: Fruit\n",
	}

	for i, test := range seriesArray {
		output := test.PrintSeries()
		if output != expectedArray[i] {
			t.Fatalf("expected: %s, got: %s", expectedArray[i], output)
		}
	}
}

func TestLoc(t *testing.T) {

}
