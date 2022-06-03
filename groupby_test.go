package gambas

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGroupByAgg(t *testing.T) {
	type aggTest struct {
		arg1     GroupBy
		arg2     []string
		arg3     StatsFunc
		expected *DataFrame
	}
	aggTests := []aggTest{
		{
			func() GroupBy {
				newDf, err := ReadCsv("./testfiles/airquality.csv", []string{"date.utc"})
				if err != nil {
					t.Error(err)
				}
				gb, _ := newDf.GroupBy("parameter", "location")
				return *gb
			}(),
			[]string{"value"},
			Mean,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"BETR801", "FR04014", "London Westminster"},
						IndexData{
							[]Index{
								{0, []interface{}{"BETR801"}},
								{1, []interface{}{"FR04014"}},
								{2, []interface{}{"London Westminster"}},
							},
							[]string{"location"},
						},
						"location",
					},
					{
						[]interface{}{26.951, 29.374, 29.740},
						IndexData{
							[]Index{
								{0, []interface{}{"BETR801"}},
								{1, []interface{}{"FR04014"}},
								{2, []interface{}{"London Westminster"}},
							},
							[]string{"location"},
						},
						"no2",
					},
					{
						[]interface{}{23.169, math.NaN(), 13.444},
						IndexData{
							[]Index{
								{0, []interface{}{"BETR801"}},
								{1, []interface{}{"FR04014"}},
								{2, []interface{}{"London Westminster"}},
							},
							[]string{"location"},
						},
						"pm25",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"BETR801"}},
						{1, []interface{}{"FR04014"}},
						{2, []interface{}{"London Westminster"}},
					},
					[]string{"location"},
				},
				[]string{"location", "no2", "pm25"},
			},
		},
	}
	for _, test := range aggTests {
		output, err := test.arg1.Agg(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}, GroupBy{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}
