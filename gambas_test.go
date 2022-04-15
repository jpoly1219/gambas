package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateRangeIndex(t *testing.T) {
	type createRangeIndexTest struct {
		arg1     int
		expected IndexData
	}

	createRangeIndexTests := []createRangeIndexTest{
		{
			5,
			IndexData{
				[]Index{{0}, {1}, {2}, {3}, {4}},
				[]string{""},
			},
		},
		{
			10,
			IndexData{
				[]Index{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}},
				[]string{""},
			},
		},
		{
			1,
			IndexData{
				[]Index{{0}},
				[]string{""},
			},
		},
	}

	for _, test := range createRangeIndexTests {
		output := CreateRangeIndex(test.arg1)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(IndexData{})) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestNewSeries(t *testing.T) {
	type newSeriesTest struct {
		arg1     []interface{}
		arg2     string
		arg3     *IndexData
		expected *Series
	}

	newSeriesTests := []newSeriesTest{
		{
			[]interface{}{"alice", "bob", "charlie"},
			"People",
			&IndexData{
				[]Index{{0}, {1}, {2}},
				[]string{""},
			},
			&Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{{0}, {1}, {2}},
					[]string{""},
				},
				"People",
			},
		},
		{
			[]interface{}{"apple", "banana", "cherry"},
			"Fruit",
			&IndexData{
				[]Index{{0}, {1}, {2}},
				[]string{""},
			},
			&Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{{0}, {1}, {2}},
					[]string{""},
				},
				"Fruit",
			},
		},
		{
			[]interface{}{"apple", 2, "cherry"},
			"Fruit",
			&IndexData{
				[]Index{{0}, {1}, {2}},
				[]string{""},
			},
			nil,
		},
		{
			[]interface{}{"alice", "bob", "charlie"},
			"People",
			&IndexData{
				[]Index{{0, "female"}, {1, "male"}, {2, "male"}},
				[]string{"id", "sex"},
			},
			&Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{{0, "female"}, {1, "male"}, {2, "male"}},
					[]string{"id", "sex"},
				},
				"People",
			},
		},
	}

	for _, test := range newSeriesTests {
		output, err := NewSeries(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestNewDataFrame(t *testing.T) {
	type newDataFrameTest struct {
		arg1     [][]interface{}
		arg2     []string
		arg3     []string
		expected *DataFrame
	}

	newDataFrameTests := []newDataFrameTest{
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]string{"group a", "group b", "group c"},
			[]string{"group a"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{1, 2, 3},
						IndexData{
							[]Index{{1}, {2}, {3}},
							[]string{"group a"},
						},
						"group a",
					},
					{
						[]interface{}{4, 5, 6},
						IndexData{
							[]Index{{1}, {2}, {3}},
							[]string{"group a"},
						},
						"group b",
					},
					{
						[]interface{}{7, 8, 9},
						IndexData{
							[]Index{{1}, {2}, {3}},
							[]string{"group a"},
						},
						"group c",
					},
				},
				IndexData{
					[]Index{{1}, {2}, {3}},
					[]string{"group a"},
				},
				[]string{"group a", "group b", "group c"},
			},
		},
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]string{"group a", "group b", "group c"},
			[]string{"group a", "group c"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{1, 2, 3},
						IndexData{
							[]Index{{1, 7}, {2, 8}, {3, 9}},
							[]string{"group a", "group c"},
						},
						"group a",
					},
					{
						[]interface{}{4, 5, 6},
						IndexData{
							[]Index{{1, 7}, {2, 8}, {3, 9}},
							[]string{"group a", "group c"},
						},
						"group b",
					},
					{
						[]interface{}{7, 8, 9},
						IndexData{
							[]Index{{1, 7}, {2, 8}, {3, 9}},
							[]string{"group a", "group c"},
						},
						"group c",
					},
				},
				IndexData{
					[]Index{{1, 7}, {2, 8}, {3, 9}},
					[]string{"group a", "group c"},
				},
				[]string{"group a", "group b", "group c"},
			},
		},
	}

	for _, test := range newDataFrameTests {
		output, err := NewDataFrame(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
