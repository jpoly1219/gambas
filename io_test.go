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
		expected *DataFrame
	}

	readCsvTests := []readCsvTest{
		{
			filepath.Join("testfiles", "test1.csv"),
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradford", "Candice"},
						IndexData{
							[]Index{{0}, {1}, {2}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 25.0, 22.0},
						IndexData{
							[]Index{{0}, {1}, {2}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0}, {1}, {2}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				Index{"Name", "Age", "Sex"},
				[]interface{}{"Name"},
				[]Index{{"Avery", "Bradford", "Candice"}},
			},
		},
		{
			filepath.Join("testfiles", "test2.csv"),
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Number",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						IndexData{
							[]Index{{0}, {1}, {2}, {3}},
							[]string{"Name"},
						},
						"Salary",
					},
				},
				Index{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
				[]interface{}{"Name"},
				[]Index{{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"}},
			},
		},
	}

	for _, test := range readCsvTests {
		output, err := ReadCsv(test.arg1)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}), cmpopts.IgnoreTypes(0.0)) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
