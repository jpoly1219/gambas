package gambas

import "testing"

func TestPlot(t *testing.T) {
	type plotTest struct {
		arg1 DataFrame
	}
	plotTests := []plotTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/neo_v2.csv", []string{"id"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
	}
	for _, test := range plotTests {
		// SetXdata("time"), SetTimefmt("%Y-%m-%d %H:%M:%S+%M:%S")
		err := test.arg1.Plot("est_diameter_min", "relative_velocity")
		if err != nil {
			t.Fatalf("error %v", err)
		}
	}
}
