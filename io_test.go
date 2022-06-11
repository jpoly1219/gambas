package gambas

import (
	"math"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestReadCsv(t *testing.T) {
	type readCsvTest struct {
		arg1     string
		arg2     []string
		expected *DataFrame
	}

	readCsvTests := []readCsvTest{
		{
			filepath.Join("testfiles", "test1.csv"),
			nil,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradford", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradford"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 25.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradford"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradford"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradford"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
		{
			filepath.Join("testfiles", "test2.csv"),
			nil,
			&DataFrame{
				[]Series{
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
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Number",
					},
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
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Weight",
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
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Salary",
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
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			filepath.Join("testfiles", "test2.csv"),
			[]string{"Position"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Number",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG"}},
								{1, []interface{}{"SF"}},
								{2, []interface{}{"SG"}},
								{3, []interface{}{"SG"}},
							},
							[]string{"Position"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"PG"}},
						{1, []interface{}{"SF"}},
						{2, []interface{}{"SG"}},
						{3, []interface{}{"SG"}},
					},
					[]string{"Position"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			filepath.Join("testfiles", "test2.csv"),
			nil,
			&DataFrame{
				[]Series{
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
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Number",
					},
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
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Weight",
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
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Salary",
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
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			filepath.Join("testfiles", "test2.csv"),
			[]string{"Position", "College"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Number",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						IndexData{
							[]Index{
								{0, []interface{}{"PG", "Texas"}},
								{1, []interface{}{"SF", "Marquette"}},
								{2, []interface{}{"SG", "Boston University"}},
								{3, []interface{}{"SG", "Georgia State"}},
							},
							[]string{"Position", "College"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"PG", "Texas"}},
						{1, []interface{}{"SF", "Marquette"}},
						{2, []interface{}{"SG", "Boston University"}},
						{3, []interface{}{"SG", "Georgia State"}},
					},
					[]string{"Position", "College"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
	}

	for _, test := range readCsvTests {
		output, err := ReadCsv(test.arg1, test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.IgnoreTypes(0.0)) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestWriteCsv(t *testing.T) {
	type writeCsvTest struct {
		arg1     DataFrame
		arg2     string
		expected error
	}
	writeCsvTests := []writeCsvTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return *newDf
			}(
				[][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}},
				[]string{"Name", "Age", "Sex"},
				[]string{"Name"},
			),
			filepath.Join("testfiles", "writeCsvTest1.csv"),
			nil,
		},
		{
			func() DataFrame {
				df, err := ReadCsv(filepath.Join("testfiles", "nba.csv"), nil)
				if err != nil {
					t.Error(err)
				}
				return *df
			}(),
			filepath.Join("testfiles", "writeCsvTest2.csv"),
			nil,
		},
	}

	for _, test := range writeCsvTests {
		output, err := WriteCsv(test.arg1, test.arg2)
		if output != nil && err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestIoReadJsonByColumns(t *testing.T) {
	type readJsonByColumnsTest struct {
		arg1     string
		arg2     []string
		expected *DataFrame
	}
	readJsonByColumnsTests := []readJsonByColumnsTest{
		{
			"testfiles/readjsonbycolumns/1.json",
			[]string{"Name"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{
							"Avery", "Bradley", "Candice",
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{
							19.0, 26.0, 23.0,
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{
							"Male", "Male", "Female",
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery"}},
						{1, []interface{}{"Bradley"}},
						{2, []interface{}{"Candice"}},
					},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
		{
			"testfiles/readjsonbycolumns/2.json",
			[]string{"Name"},
			&DataFrame{
				[]Series{
					{
						[]interface{}{
							"Avery", "Bradley", "Candice",
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{
							19.0, 26.0, math.NaN(),
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{
							"Male", "NaN", "Female",
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery"}},
						{1, []interface{}{"Bradley"}},
						{2, []interface{}{"Candice"}},
					},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range readJsonByColumnsTests {
		output, err := ReadJsonByColumns(test.arg1, test.arg2)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func TestIoReadJsonByRows(t *testing.T) {
	type readJsonByRowsTest struct {
		arg1     string
		expected *DataFrame
	}
	readJsonByRowsTests := []readJsonByRowsTest{
		{
			"testfiles/readjsonbyrows/1.json",
			&DataFrame{
				[]Series{
					{
						[]interface{}{
							"Avery", "Bradley", "Candice",
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"index"},
						},
						"index",
					},
					{
						[]interface{}{
							19.0, 26.0, 23.0,
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"index"},
						},
						"Age",
					},
					{
						[]interface{}{
							"Male", "Male", "Female",
						},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery"}},
								{1, []interface{}{"Bradley"}},
								{2, []interface{}{"Candice"}},
							},
							[]string{"index"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery"}},
						{1, []interface{}{"Bradley"}},
						{2, []interface{}{"Candice"}},
					},
					[]string{"index"},
				},
				[]string{"index", "Age", "Sex"},
			},
		},
	}
	for _, test := range readJsonByRowsTests {
		output, err := ReadJsonByRows(test.arg1)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
