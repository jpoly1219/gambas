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
				[]Index{
					{0, []interface{}{0}},
					{1, []interface{}{1}},
					{2, []interface{}{2}},
					{3, []interface{}{3}},
					{4, []interface{}{4}},
				},
				[]string{""},
			},
		},
		{
			10,
			IndexData{
				[]Index{
					{0, []interface{}{0}},
					{1, []interface{}{1}},
					{2, []interface{}{2}},
					{3, []interface{}{3}},
					{4, []interface{}{4}},
					{5, []interface{}{5}},
					{6, []interface{}{6}},
					{7, []interface{}{7}},
					{8, []interface{}{8}},
					{9, []interface{}{9}},
				},
				[]string{""},
			},
		},
		{
			1,
			IndexData{
				[]Index{{0, []interface{}{0}}},
				[]string{""},
			},
		},
	}

	for _, test := range createRangeIndexTests {
		output := CreateRangeIndex(test.arg1)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(IndexData{}, Index{})) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestNewSeries(t *testing.T) {
	type newSeriesTest struct {
		arg1     []interface{}
		arg2     string
		arg3     *IndexData
		expected Series
	}

	newSeriesTests := []newSeriesTest{
		{
			[]interface{}{"alice", "bob", "charlie"},
			"People",
			&IndexData{
				[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
				[]string{""},
			},
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
					[]string{""},
				},
				"People",
				"string",
			},
		},
		{
			[]interface{}{"apple", "banana", "cherry"},
			"Fruit",
			&IndexData{
				[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
				[]string{""},
			},
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
					[]string{""},
				},
				"Fruit",
				"string",
			},
		},
		{
			[]interface{}{"apple", 2, "cherry"},
			"Fruit",
			&IndexData{
				[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
				[]string{""},
			},
			Series{
				[]interface{}{"apple", "2", "cherry"},
				IndexData{
					[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
					[]string{""},
				},
				"Fruit",
				"string",
			},
		},
		{
			[]interface{}{"alice", "bob", "charlie"},
			"People",
			&IndexData{
				[]Index{{0, []interface{}{0, "female"}}, {1, []interface{}{1, "male"}}, {2, []interface{}{2, "male"}}},
				[]string{"id", "sex"},
			},
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{{0, []interface{}{0, "female"}}, {1, []interface{}{1, "male"}}, {2, []interface{}{2, "male"}}},
					[]string{"id", "sex"},
				},
				"People",
				"string",
			},
		},
	}

	for _, test := range newSeriesTests {
		output, err := NewSeries(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, Series{}, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestNewDataFrame(t *testing.T) {
	type newDataFrameTest struct {
		arg1     [][]interface{}
		arg2     []string
		arg3     []string
		expected DataFrame
	}

	newDataFrameTests := []newDataFrameTest{
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]string{"group a", "group b", "group c"},
			[]string{"group a"},
			DataFrame{
				[]Series{
					{
						[]interface{}{1, 2, 3},
						IndexData{
							[]Index{{0, []interface{}{1}}, {1, []interface{}{2}}, {2, []interface{}{3}}},
							[]string{"group a"},
						},
						"group a",
						"int",
					},
					{
						[]interface{}{4, 5, 6},
						IndexData{
							[]Index{{0, []interface{}{1}}, {1, []interface{}{2}}, {2, []interface{}{3}}},
							[]string{"group a"},
						},
						"group b",
						"int",
					},
					{
						[]interface{}{7, 8, 9},
						IndexData{
							[]Index{{0, []interface{}{1}}, {1, []interface{}{2}}, {2, []interface{}{3}}},
							[]string{"group a"},
						},
						"group c",
						"int",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{1}}, {1, []interface{}{2}}, {2, []interface{}{3}}},
					[]string{"group a"},
				},
				[]string{"group a", "group b", "group c"},
			},
		},
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]string{"group a", "group b", "group c"},
			[]string{"group a", "group c"},
			DataFrame{
				[]Series{
					{
						[]interface{}{1, 2, 3},
						IndexData{
							[]Index{{0, []interface{}{1, 7}}, {1, []interface{}{2, 8}}, {2, []interface{}{3, 9}}},
							[]string{"group a", "group c"},
						},
						"group a",
						"int",
					},
					{
						[]interface{}{4, 5, 6},
						IndexData{
							[]Index{{0, []interface{}{1, 7}}, {1, []interface{}{2, 8}}, {2, []interface{}{3, 9}}},
							[]string{"group a", "group c"},
						},
						"group b",
						"int",
					},
					{
						[]interface{}{7, 8, 9},
						IndexData{
							[]Index{{0, []interface{}{1, 7}}, {1, []interface{}{2, 8}}, {2, []interface{}{3, 9}}},
							[]string{"group a", "group c"},
						},
						"group c",
						"int",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{1, 7}}, {1, []interface{}{2, 8}}, {2, []interface{}{3, 9}}},
					[]string{"group a", "group c"},
				},
				[]string{"group a", "group b", "group c"},
			},
		},
		{
			[][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]string{"group a", "group b", "group c"},
			nil,
			DataFrame{
				[]Series{
					{
						[]interface{}{1, 2, 3},
						IndexData{
							[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
							[]string{""},
						},
						"group a",
						"int",
					},
					{
						[]interface{}{4, 5, 6},
						IndexData{
							[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
							[]string{""},
						},
						"group b",
						"int",
					},
					{
						[]interface{}{7, 8, 9},
						IndexData{
							[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
							[]string{""},
						},
						"group c",
						"int",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{0}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
					[]string{""},
				},
				[]string{"group a", "group b", "group c"},
			},
		},
	}

	for _, test := range newDataFrameTests {
		output, err := NewDataFrame(test.arg1, test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
