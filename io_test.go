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
						Index{[]interface{}{0, 1, 2}},
						"Name",
					},
					{
						[]interface{}{19.0, 25.0, 22.0},
						Index{[]interface{}{0, 1, 2}},
						"Age",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						Index{[]interface{}{0, 1, 2}},
						"Sex",
					},
				},
				Index{[]interface{}{0, 1, 2}},
				Index{[]interface{}{"Name", "Age", "Sex"}},
			},
		},
		{
			filepath.Join("testfiles", "test2.csv"),
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						Index{[]interface{}{0, 1, 2, 3}},
						"Name",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						Index{[]interface{}{0, 1, 2, 3}},
						"Team",
					},
					{
						[]interface{}{0.0, 99.0, 30.0, 28.0},
						Index{[]interface{}{0, 1, 2, 3}},
						"Number",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						Index{[]interface{}{0, 1, 2, 3}},
						"Position",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						Index{[]interface{}{0, 1, 2, 3}},
						"Age",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						Index{[]interface{}{0, 1, 2, 3}},
						"Height",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						Index{[]interface{}{0, 1, 2, 3}},
						"Weight",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						Index{[]interface{}{0, 1, 2, 3}},
						"College",
					},
					{
						[]interface{}{7730337.0, 6796117.0, math.NaN(), 1148640.0},
						Index{[]interface{}{0, 1, 2, 3}},
						"Salary",
					},
				},
				Index{[]interface{}{0, 1, 2, 3}},
				Index{[]interface{}{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"}},
			},
		},
	}

	for _, test := range readCsvTests {
		output, err := ReadCsv(test.arg1)
		if !cmp.Equal(*output, *test.expected, cmp.AllowUnexported(*output, output.series[0], output.index), cmpopts.IgnoreTypes(0.0)) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
