package gambas

import "testing"

func TestPlot(t *testing.T) {
	type plotTest struct {
		arg1 DataFrame
	}
	plotTests := []plotTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/airquality.csv", []string{"Time"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
	}
	for _, test := range plotTests {
		err := test.arg1.Plot("date.utc", "value")
		if err != nil {
			t.Fatalf("error %v", err)
		}
	}
}
