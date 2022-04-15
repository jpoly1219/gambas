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
							[]Index{{"Avery"}, {"Bradford"}, {"Candice"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{19.0, 25.0, 22.0},
						IndexData{
							[]Index{{"Avery"}, {"Bradford"}, {"Candice"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{"Avery"}, {"Bradford"}, {"Candice"}},
							[]string{"Name"},
						},
						"Sex",
					},
				},
				IndexData{
					[]Index{{"Avery"}, {"Bradford"}, {"Candice"}},
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
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Number",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						IndexData{
							[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
							[]string{"Name"},
						},
						"Salary",
					},
				},
				IndexData{
					[]Index{{"Avery Bradley"}, {"Jae Crowder"}, {"John Holland"}, {"R.J. Hunter"}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
	}

	for _, test := range readCsvTests {
		output, err := ReadCsv(test.arg1, test.arg2)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}), cmpopts.IgnoreTypes(0.0)) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
