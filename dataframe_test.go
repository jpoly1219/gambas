package gambas

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDataFramePrint(t *testing.T) {
	type printTest struct {
		arg1 DataFrame
	}
	printTests := []printTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
		},
	}

	for _, test := range printTests {
		test.arg1.Print()
	}
}

func TestDataFramePrintRange(t *testing.T) {
	type printRangeTest struct {
		arg1 DataFrame
	}
	printRangeTests := []printRangeTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
		},
	}

	for _, test := range printRangeTests {
		test.arg1.PrintRange(1, 2)
	}
}

func TestDataFrameHead(t *testing.T) {
	type headTest struct {
		arg1 DataFrame
	}
	headTests := []headTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
		},
	}

	for _, test := range headTests {
		test.arg1.Head(2)
	}
}

func TestDataFrameTail(t *testing.T) {
	type tailTest struct {
		arg1 DataFrame
	}
	tailTests := []tailTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
		},
	}

	for _, test := range tailTests {
		test.arg1.Tail(3)
	}
}
func TestDataFrameLocRows(t *testing.T) {
	type dataframeLocRowsTest struct {
		arg1     DataFrame
		arg2     [][]interface{}
		expected *DataFrame
	}
	dataframeLocRowsTests := []dataframeLocRowsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[][]interface{}{{"Avery"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Jae Crowder"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{99.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Number",
					},
					{
						[]interface{}{"SF"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Position",
					},
					{
						[]interface{}{25.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"6-6"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{235.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Weight",
					},
					{
						[]interface{}{"Marquette"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{6796117.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Jae Crowder"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder"}, {"Avery Bradley"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Jae Crowder", "Avery Bradley"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{99.0, 0.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Number",
					},
					{
						[]interface{}{"SF", "PG"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"6-6", "6-2"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{235.0, 180.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Weight",
					},
					{
						[]interface{}{"Marquette", "Texas"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{6796117.0, 7730337.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder", 25.0}, {"Avery Bradley", 25.0}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Jae Crowder", "Avery Bradley"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Team",
					},
					{
						[]interface{}{99.0, 0.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Number",
					},
					{
						[]interface{}{"SF", "PG"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Age",
					},
					{
						[]interface{}{"6-6", "6-2"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Height",
					},
					{
						[]interface{}{235.0, 180.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Weight",
					},
					{
						[]interface{}{"Marquette", "Texas"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"College",
					},
					{
						[]interface{}{6796117.0, 7730337.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
					[]string{"Name", "Age"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder", 22.0}, {"Avery Bradley", 25.0}},
			nil,
		},
	}

	for _, test := range dataframeLocRowsTests {
		output, err := test.arg1.LocRows(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLocRowsItems(t *testing.T) {
	type dataframeLocRowsItemsTest struct {
		arg1     DataFrame
		arg2     [][]interface{}
		expected [][]interface{}
	}
	dataframeLocRowsItemsTests := []dataframeLocRowsItemsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[][]interface{}{{"Avery"}},
			[][]interface{}{{"Avery", 19, "Male"}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder"}},
			[][]interface{}{{"Jae Crowder", "Boston Celtics", 99.0, "SF", 25.0, "6-6", 235.0, "Marquette", 6796117.0}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder"}, {"Avery Bradley"}},
			[][]interface{}{
				{"Jae Crowder", "Boston Celtics", 99.0, "SF", 25.0, "6-6", 235.0, "Marquette", 6796117.0},
				{"Avery Bradley", "Boston Celtics", 0.0, "PG", 25.0, "6-2", 180.0, "Texas", 7730337.0},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder", 25.0}, {"Avery Bradley", 25.0}},
			[][]interface{}{
				{"Jae Crowder", "Boston Celtics", 99.0, "SF", 25.0, "6-6", 235.0, "Marquette", 6796117.0},
				{"Avery Bradley", "Boston Celtics", 0.0, "PG", 25.0, "6-2", 180.0, "Texas", 7730337.0},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[][]interface{}{{"Jae Crowder", 22.0}, {"Avery Bradley", 25.0}},
			nil,
		},
	}

	for _, test := range dataframeLocRowsItemsTests {
		output, err := test.arg1.LocRowsItems(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLocCols(t *testing.T) {
	type dataframeLocColsTest struct {
		arg1     DataFrame
		arg2     []string
		expected *DataFrame
	}
	dataframeLocColsTests := []dataframeLocColsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{19, 27, 22},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccols1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Position"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Position",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery Bradley"}},
						{1, []interface{}{"Jae Crowder"}},
						{2, []interface{}{"John Holland"}},
						{3, []interface{}{"R.J. Hunter"}},
					},
					[]string{"Name"},
				},
				[]string{"Position"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccols1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Age", "College", "Name"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Name",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery Bradley"}},
						{1, []interface{}{"Jae Crowder"}},
						{2, []interface{}{"John Holland"}},
						{3, []interface{}{"R.J. Hunter"}},
					},
					[]string{"Name"},
				},
				[]string{"Age", "College", "Name"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccols1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Avery Bradley"},
			nil,
		},
	}
	for _, test := range dataframeLocColsTests {
		output, err := test.arg1.LocCols(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLocColsItems(t *testing.T) {
	type dataframeLocColsItemsTest struct {
		arg1     DataFrame
		arg2     []string
		expected [][]interface{}
	}
	dataframeLocColsItemsTests := []dataframeLocColsItemsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			[][]interface{}{{19, 27, 22}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccolsitems1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Position"},
			[][]interface{}{{"PG", "SF", "SG", "SG"}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccolsitems1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Age", "College", "Name"},
			[][]interface{}{
				{25.0, 25.0, 27.0, 22.0},
				{"Texas", "Marquette", "Boston University", "Georgia State"},
				{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccolsitems1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Avery Bradley"},
			nil,
		},
	}
	for _, test := range dataframeLocColsItemsTests {
		output, err := test.arg1.LocColsItems(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLoc(t *testing.T) {
	type dataframeLocTest struct {
		arg1     DataFrame
		arg2     []string
		arg3     [][]interface{}
		expected *DataFrame
	}
	dataframeLocTests := []dataframeLocTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			[][]interface{}{{"Bradley"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{27},
						IndexData{
							[]Index{{1, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Bradley"}}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloc1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			[]string{"Age", "Name"},
			[][]interface{}{{"John Holland"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{27.0},
						IndexData{
							[]Index{{2, []interface{}{"John Holland"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"John Holland"},
						IndexData{
							[]Index{{2, []interface{}{"John Holland"}}},
							[]string{"Name"},
						},
						"Name",
					},
				},
				IndexData{
					[]Index{{2, []interface{}{"John Holland"}}},
					[]string{"Name"},
				},
				[]string{"Age", "Name"},
			},
		},
	}
	for _, test := range dataframeLocTests {
		output, err := test.arg1.Loc(test.arg2, test.arg3...)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColAdd(t *testing.T) {
	type colAddTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colAddTests := []colAddTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{24.0, 32.0, 27.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colAddTests {
		output, err := test.arg1.ColAdd(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColSub(t *testing.T) {
	type colSubTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colSubTests := []colSubTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{14.0, 22.0, 17.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colSubTests {
		output, err := test.arg1.ColSub(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColMul(t *testing.T) {
	type colMulTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colMulTests := []colMulTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			2.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{38.0, 54.0, 44.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colMulTests {
		output, err := test.arg1.ColMul(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColDiv(t *testing.T) {
	type colDivTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colDivTests := []colDivTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{3.8, 5.4, 4.4},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colDivTests {
		output, err := test.arg1.ColDiv(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColMod(t *testing.T) {
	type colModTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colModTests := []colModTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{4.0, 2.0, 2.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colModTests {
		output, err := test.arg1.ColMod(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColGt(t *testing.T) {
	type colGtTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colGtTests := []colGtTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			25.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{false, true, false},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colGtTests {
		output, err := test.arg1.ColGt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColLt(t *testing.T) {
	type colLtTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colLtTests := []colLtTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			22.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{true, false, false},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colLtTests {
		output, err := test.arg1.ColLt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestColEq(t *testing.T) {
	type colEqTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     float64
		expected *DataFrame
	}

	colEqTests := []colEqTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			nil,
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			19.0,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{true, false, false},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colEqTests {
		output, err := test.arg1.ColEq(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestNewCol(t *testing.T) {
	type newColTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     []interface{}
		expected *DataFrame
	}
	newColTests := []newColTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Nationality",
			[]interface{}{"USA", "UK", "Canada"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
					{
						[]interface{}{"USA", "UK", "Canada"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Nationality",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "Nationality"},
			},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age+5",
			make([]interface{}, 3),
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
					{
						[]interface{}{nil, nil, nil},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age+5",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "Age+5"},
			},
		},
	}
	for _, test := range newColTests {
		output, err := test.arg1.NewCol(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestNewDerivedCol(t *testing.T) {
	type newDerivedColTest struct {
		arg1     *DataFrame
		arg2     string
		arg3     string
		expected *DataFrame
	}
	newDerivedColTests := []newDerivedColTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"NewAge",
			"Age",
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"NewAge",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "NewAge"},
			},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) *DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"NewCol",
			"Doesn't Exist",
			nil,
		},
	}
	for _, test := range newDerivedColTests {
		output, err := test.arg1.NewDerivedCol(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameRenameCol(t *testing.T) {
	type renameColTest struct {
		arg1     DataFrame
		arg2     map[string]string
		expected DataFrame
	}
	renameColTests := []renameColTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			map[string]string{"Name": "Names"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Names", "Age", "Sex"}, []string{"Names"}),
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			map[string]string{"Age": "HowOld"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "HowOld", "Sex"}, []string{"Name"}),
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name", "Sex"}),
			map[string]string{"Name": "Names"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Names", "Age", "Sex"}, []string{"Names", "Sex"}),
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name", "Sex"}),
			map[string]string{"Names": "Name"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name", "Sex"}),
		},
	}

	for _, test := range renameColTests {
		err := test.arg1.RenameCol(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			if fmt.Sprint(err)[:22] == "column does not exist:" || fmt.Sprint(err)[:21] == "index does not exist:" {
				continue
			}
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func TestDataFrameSortByIndex(t *testing.T) {
	type sortByIndexTest struct {
		arg1     DataFrame
		arg2     bool
		expected DataFrame
	}
	sortByIndexTests := []sortByIndexTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Bradley", "Candice", "Avery"}, {27.0, 22.0, 19.0}, {"Male", "Female", "Male"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			true,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
		// {
		// 	func() DataFrame {
		// 		newDf, err := ReadCsv("./testfiles/nba.csv", []string{"Team"})
		// 		if err != nil {
		// 			t.Error(err)
		// 		}
		// 		return *newDf
		// 	}(),
		// 	true,
		// 	func() DataFrame {
		// 		newDf, err := ReadCsv("./testfiles/nba.csv", []string{"Team"})
		// 		if err != nil {
		// 			t.Error(err)
		// 		}
		// 		return *newDf
		// 	}(),
		// },
	}
	for _, test := range sortByIndexTests {
		test.arg1.Print()
		err := test.arg1.SortByIndex(test.arg2)
		test.arg1.Print()
		fmt.Println("----------")
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func TestDataFrameSortByValues(t *testing.T) {
	type sortByValuesTest struct {
		arg1     DataFrame
		arg2     string
		arg3     bool
		expected DataFrame
	}
	sortByValuesTests := []sortByValuesTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Bradley", "Candice", "Avery"}, {27.0, 22.0, 19.0}, {"Male", "Female", "Male"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			true,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Candice", "Bradley"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 22.0, 27.0},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Female", "Male"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}

	for _, test := range sortByValuesTests {
		test.arg1.Print()
		err := test.arg1.SortByValues(test.arg2, test.arg3)
		test.arg1.Print()
		fmt.Println("----------")
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func TestDropNaN(t *testing.T) {
	type dropNaNTest struct {
		arg1     DataFrame
		arg2     int
		expected DataFrame
	}
	dropNaNTests := []dropNaNTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "R.J. Hunter"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 28.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Number",
					},
					{
						[]interface{}{"PG", "SF", "SG"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 185.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Georgia State"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, 1148640.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan2.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			1,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, 5000000.0, 1148640.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
	}

	for _, test := range dropNaNTests {
		err := test.arg1.DropNaN(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func TestDataFramePivot(t *testing.T) {
	type pivotTest struct {
		arg1     DataFrame
		arg2     string
		arg3     string
		expected *DataFrame
	}
	pivotTests := []pivotTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfpivot1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			"Sex",
			"Height",
			&DataFrame{
				[]Series{
					{
						[]interface{}{172.0, 180.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Male",
					},
					{
						[]interface{}{math.NaN(), math.NaN(), 165.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Female",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Male", "Female"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfpivot2.csv", []string{"Time"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			"Fruit",
			"Color",
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Red", "NaN"},
						IndexData{
							[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
							[]string{"Time"},
						},
						"Apple",
					},
					{
						[]interface{}{"Yellow", "Yellow"},
						IndexData{
							[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
							[]string{"Time"},
						},
						"Banana",
					},
					{
						[]interface{}{"NaN", "Red"},
						IndexData{
							[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
							[]string{"Time"},
						},
						"Cherry",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
					[]string{"Time"},
				},
				[]string{"Apple", "Banana", "Cherry"},
			},
		},
	}

	for _, test := range pivotTests {
		output, err := test.arg1.Pivot(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFramePivotTable(t *testing.T) {
	type pivotTableTest struct {
		arg1     DataFrame
		arg2     string
		arg3     string
		arg4     string
		arg5     StatsFunc
		expected *DataFrame
	}
	pivotTableTests := []pivotTableTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfpivottable1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			"Team",
			"Height",
			"Salary",
			Mean,
			&DataFrame{
				[]Series{
					{
						[]interface{}{math.NaN(), 1500000.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"5-11",
					},
					{
						[]interface{}{6912869.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"5-9",
					},
					{
						[]interface{}{5000000.0, 958633.333},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-10",
					},
					{
						[]interface{}{math.NaN(), 1140240.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-11",
					},
					{
						[]interface{}{4777348.5, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-2",
					},
					{
						[]interface{}{math.NaN(), 2697445.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-3",
					},
					{
						[]interface{}{3431040.0, 817107.5},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-4",
					},
					{
						[]interface{}{1148640.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-5",
					},
					{
						[]interface{}{4272978.5, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-6",
					},
					{
						[]interface{}{3425510.0, 1467660.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-7",
					},
					{
						[]interface{}{1170960.0, 7330732.5},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-8",
					},
					{
						[]interface{}{7284630.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-9",
					},
					{
						[]interface{}{2391067.5, 19689000.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"7-0",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
					[]string{"Team"},
				},
				[]string{"5-11", "5-9", "6-10", "6-11", "6-2", "6-3", "6-4", "6-5", "6-6", "6-7", "6-8", "6-9", "7-0"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/airquality.csv", []string{"Time"})
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(),
			"location",
			"parameter",
			"value",
			Mean,
			&DataFrame{
				[]Series{
					{
						[]interface{}{26.951, 29.374, 29.740},
						IndexData{
							[]Index{{0, []interface{}{"BETR801"}}, {1, []interface{}{"FR04014"}}, {2, []interface{}{"London Westminster"}}},
							[]string{"location"},
						},
						"no2",
					},
					{
						[]interface{}{23.169, math.NaN(), 13.444},
						IndexData{
							[]Index{{0, []interface{}{"BETR801"}}, {1, []interface{}{"FR04014"}}, {2, []interface{}{"London Westminster"}}},
							[]string{"location"},
						},
						"pm25",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"BETR801"}}, {1, []interface{}{"FR04014"}}, {2, []interface{}{"London Westminster"}}},
					[]string{"location"},
				},
				[]string{"no2", "pm25"},
			},
		},
	}

	for _, test := range pivotTableTests {
		output, err := test.arg1.PivotTable(test.arg2, test.arg3, test.arg4, test.arg5)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
