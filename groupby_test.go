package gambas

import (
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
				newDf, err := ReadCsv("./testfiles/testgbagg.csv", []string{"ID"})
				if err != nil {
					t.Error(err)
				}
				gb, _ := newDf.GroupBy("Animal")
				return *gb
			}(),
			[]string{"Max Speed"},
			Mean,
			&DataFrame{
				[]Series{
					{
						[]interface{}{"Falcon", "Parrot"},
						IndexData{
							[]Index{
								{0, []interface{}{"Falcon"}},
								{1, []interface{}{"Parrot"}},
							},
							[]string{"Animal"},
						},
						"Animal",
					},
					{
						[]interface{}{375.000, 25.000},
						IndexData{
							[]Index{
								{0, []interface{}{"Falcon"}},
								{1, []interface{}{"Parrot"}},
							},
							[]string{"Animal"},
						},
						"Max Speed",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Falcon"}},
						{1, []interface{}{"Parrot"}},
					},
					[]string{"Animal"},
				},
				[]string{"Animal", "Max Speed"},
			},
		},
		{
			func() GroupBy {
				newDf, err := ReadCsv("./testfiles/titanic.csv", []string{"PassengerId"})
				if err != nil {
					t.Error(err)
				}
				gb, _ := newDf.GroupBy("Pclass")
				return *gb
			}(),
			[]string{"Age"},
			Mean,
			&DataFrame{
				[]Series{
					{
						[]interface{}{1.0, 2.0, 3.0},
						IndexData{
							[]Index{
								{1, []interface{}{1.0}},
								{2, []interface{}{2.0}},
								{0, []interface{}{3.0}},
							},
							[]string{"Pclass"},
						},
						"Pclass",
					},
					{
						[]interface{}{38.233, 29.878, 25.141},
						IndexData{
							[]Index{
								{1, []interface{}{1.0}},
								{2, []interface{}{2.0}},
								{0, []interface{}{3.0}},
							},
							[]string{"Pclass"},
						},
						"Age",
					},
				},
				IndexData{
					[]Index{
						{1, []interface{}{1.0}},
						{2, []interface{}{2.0}},
						{0, []interface{}{3.0}},
					},
					[]string{"Pclass"},
				},
				[]string{"Pclass", "Age"},
			},
		},
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
						[]interface{}{"no2", "no2", "no2", "pm25", "pm25"},
						IndexData{
							[]Index{
								{3, []interface{}{"no2", "BETR801"}},
								{2, []interface{}{"no2", "FR04014"}},
								{4, []interface{}{"no2", "London Westminster"}},
								{0, []interface{}{"pm25", "BETR801"}},
								{1, []interface{}{"pm25", "London Westminster"}},
							},
							[]string{"parameter", "location"},
						},
						"parameter",
					},
					{
						[]interface{}{"BETR801", "FR04014", "London Westminster", "BETR801", "London Westminster"},
						IndexData{
							[]Index{
								{3, []interface{}{"no2", "BETR801"}},
								{2, []interface{}{"no2", "FR04014"}},
								{4, []interface{}{"no2", "London Westminster"}},
								{0, []interface{}{"pm25", "BETR801"}},
								{1, []interface{}{"pm25", "London Westminster"}},
							},
							[]string{"parameter", "location"},
						},
						"location",
					},
					{
						[]interface{}{26.951, 29.374, 29.740, 23.169, 13.444},
						IndexData{
							[]Index{
								{3, []interface{}{"no2", "BETR801"}},
								{2, []interface{}{"no2", "FR04014"}},
								{4, []interface{}{"no2", "London Westminster"}},
								{0, []interface{}{"pm25", "BETR801"}},
								{1, []interface{}{"pm25", "London Westminster"}},
							},
							[]string{"parameter", "location"},
						},
						"value",
					},
				},
				IndexData{
					[]Index{
						{3, []interface{}{"no2", "BETR801"}},
						{2, []interface{}{"no2", "FR04014"}},
						{4, []interface{}{"no2", "London Westminster"}},
						{0, []interface{}{"pm25", "BETR801"}},
						{1, []interface{}{"pm25", "London Westminster"}},
					},
					[]string{"parameter", "location"},
				},
				[]string{"parameter", "location", "value"},
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
