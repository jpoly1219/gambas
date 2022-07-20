package gambas

import (
	"testing"
)

func TestPlot(t *testing.T) {
	type plotTest struct {
		arg1 PlotData
	}
	plotTests := []plotTest{
		{
			PlotData{
				func() *DataFrame {
					newDf, err := ReadCsv("./testfiles/neo_v2.csv", []string{"id"})
					if err != nil {
						t.Error(err)
					}
					return &newDf
				}(),
				[]string{"est_diameter_min", "relative_velocity"},
				"",
				[]GnuplotOpt{Using("($0/1000):1"), With("lines lc 0")},
			},
		},
	}
	for _, test := range plotTests {
		// SetXdata("time"), SetTimefmt("%Y-%m-%d %H:%M:%S+%M:%S")
		err := Plot(test.arg1)
		if err != nil {
			t.Fatalf("error %v", err)
		}
	}
}

func TestPlotN(t *testing.T) {
	type plotNTest struct {
		arg1 []PlotData
		arg2 []GnuplotOpt
	}
	plotNTests := []plotNTest{
		{
			[]PlotData{
				{
					func() *DataFrame {
						newDf, err := ReadCsv("./testfiles/neo_v2.csv", []string{"id"})
						if err != nil {
							t.Error(err)
						}
						return &newDf
					}(),
					[]string{"est_diameter_min", "relative_velocity"},
					"",
					[]GnuplotOpt{Using("0:1 lc 0")},
				},
				{
					func() *DataFrame {
						newDf, err := ReadCsv("./testfiles/neo_v2.csv", []string{"id"})
						if err != nil {
							t.Error(err)
						}
						return &newDf
					}(),
					[]string{"est_diameter_min", "miss_distance"},
					"",
					[]GnuplotOpt{Using("0:1 lc 7")},
				},
			},
			[]GnuplotOpt{},
		},
		{
			[]PlotData{
				{
					func() *DataFrame {
						newDf, err := ReadCsv("./testfiles/neo_v2.csv", []string{"id"})
						if err != nil {
							t.Error(err)
						}
						return &newDf
					}(),
					[]string{"est_diameter_min", "relative_velocity"},
					"",
					[]GnuplotOpt{Using("0:1 lc 0")},
				},
				{
					nil,
					nil,
					"10*sin(x)",
					[]GnuplotOpt{With("lines lc 7")},
				},
			},
			[]GnuplotOpt{},
		},
		{
			[]PlotData{
				{
					nil,
					nil,
					"sin(x)",
					[]GnuplotOpt{With("lines lc 8")},
				},
				{
					nil,
					nil,
					"cos(x)",
					[]GnuplotOpt{With("lines lc 3")},
				},
				{
					nil,
					nil,
					"tan(x)",
					[]GnuplotOpt{With("lines lc 7")},
				},
			},
			[]GnuplotOpt{},
		},
	}
	for _, test := range plotNTests {
		err := PlotN(test.arg1, test.arg2...)
		if err != nil {
			t.Fatalf("error %v", err)
		}
	}
}

func TestFit(t *testing.T) {
	type fitTest struct {
		arg1 string
		arg2 PlotData
		arg3 []GnuplotOpt
	}
	fitTests := []fitTest{
		{
			"a*exp(b*x)",
			PlotData{
				func() *DataFrame {
					df, err := ReadCsv("./testfiles/airquality.csv", []string{"city"})
					if err != nil {
						t.Error(err)
					}
					newDf, _ := df.LocRows([]interface{}{"Antwerpen"})
					return &newDf
				}(),
				[]string{"date.utc", "value"},
				"",
				nil,
			},
			[]GnuplotOpt{Using("0:1"), Via("a,b")},
		},
	}
	for _, test := range fitTests {
		err := Fit(test.arg1, test.arg2, test.arg3...)
		if err != nil {
			t.Fatalf("error %v", err)
		}
	}
}
