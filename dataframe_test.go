package gambas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
func TestDataFrameLocRows(t *testing.T) {
	type dataframeLocRowsTest struct {
		arg1     DataFrame
		arg2     []Index
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
			[]Index{{"Avery"}},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery"},
						IndexData{
							[]Index{{"Avery"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19},
						IndexData{
							[]Index{{"Avery"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male"},
						IndexData{
							[]Index{{"Avery"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}

	for _, test := range dataframeLocRowsTests {
		output, err := test.arg1.LocRows(test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
	}
	for _, test := range dataframeLocColsTests {
		output, err := test.arg1.LocCols(test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestDataFrameLoc(t *testing.T) {
	type dataframeLocTest struct {
		arg1     DataFrame
		arg2     []Index
		arg3     []string
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
			[]Index{{"Bradley"}},
			[]string{"Age"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{27},
						IndexData{
							[]Index{{"Bradley"}},
							[]string{"Name"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{{"Bradley"}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
	}
	for _, test := range dataframeLocTests {
		output, err := test.arg1.Loc(test.arg2, test.arg3)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{24.0, 32.0, 27.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colAddTests {
		output, err := test.arg1.ColAdd(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{14.0, 22.0, 17.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colSubTests {
		output, err := test.arg1.ColSub(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{38.0, 54.0, 44.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colMulTests {
		output, err := test.arg1.ColMul(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{3.8, 5.4, 4.4},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colDivTests {
		output, err := test.arg1.ColDiv(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{4.0, 2.0, 2.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colModTests {
		output, err := test.arg1.ColMod(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{false, true, false},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colGtTests {
		output, err := test.arg1.ColGt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{true, false, false},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colLtTests {
		output, err := test.arg1.ColLt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{true, false, false},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colEqTests {
		output, err := test.arg1.ColEq(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || (output != nil && err != nil) {
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
					{
						[]interface{}{"USA", "UK", "Canada"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Nationality",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
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
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
					{
						[]interface{}{nil, nil, nil},
						IndexData{
							[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age+5",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradley"}, {"Candice"}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "Age+5"},
			},
		},
	}
	for _, test := range newColTests {
		output, err := test.arg1.NewCol(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
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
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
		},
	}
	for _, test := range sortByIndexTests {
		err := test.arg1.SortByIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}
