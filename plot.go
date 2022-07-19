package gambas

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"path/filepath"
	"time"
)

// Plotting functionality uses gnuplot as its backend.

// PlotData holds the data required for plotting a dataset.
// Dataset is the DataFrame object you would like to plot.
// ColumnPair is a pair of columns [xcol, ycol].
// Opts is whatever gnuplot option you would like to set.
type PlotData struct {
	Df         DataFrame
	ColumnPair []string
	Opts       []GnuplotOpt
}

// Plot plots the DataFrame object.
// Choose two columns to use for the x axis and y axis.
// Then pass in any options you need. Refer to the gnuplot documentation for options.
// For example, `set xrange [-10:10]; set xlabel "myX"; set ylabel "myY"; plot "myDf.dat" using 0:1 lc 0 w lines`
// Plot(<xcol>, <ycol>, SetXrange("[-10:10]"), SetXlabel("myX"), SetYlabel("myY"), Using("0:1 lc 0 w lines"))
func (df *DataFrame) Plot(xcol, ycol string, opts ...GnuplotOpt) error {
	rand.Seed(time.Now().UnixNano())
	newDf, err := df.LocCols(xcol, ycol)
	if err != nil {
		return err
	}

	path := filepath.Join("/", "tmp", fmt.Sprintf("%x.csv", rand.Intn(100000000)))
	_, err = WriteCsv(newDf, path, true)
	if err != nil {
		return err
	}

	var setBuf bytes.Buffer
	var usingBuf bytes.Buffer
	for _, opt := range opts {
		str := opt.createCmdString()
		if opt.getOption() == "using" {
			usingBuf.WriteString(str)
		} else {
			setBuf.WriteString(str)
			setBuf.WriteString("; ")
		}
	}

	cmdString := fmt.Sprintf(`%s %s "%s" %s`, setBuf.String(), "plot", path, usingBuf.String())
	cmd := exec.Command("gnuplot", "-persist", "-e", cmdString)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err, cmd.Stderr))
	}

	return nil
}

// PlotN plots n number of datasets side by side.
// This is useful when you want to compare two different datasets,
// or a dataset with a line of best fit.
// Set options should be passed in separately as a parameter, not inside PlotData.
func PlotN(plotdata []PlotData, setOpts ...GnuplotOpt) error {
	rand.Seed(time.Now().UnixNano())

	var setBuf bytes.Buffer
	for _, setOpt := range setOpts {
		str := setOpt.createCmdString()
		setBuf.WriteString(str)
		setBuf.WriteString("; ")
	}

	cmdString := fmt.Sprintf(`%s %s `, setBuf.String(), "plot")

	for _, pd := range plotdata {
		newDf, err := pd.Df.LocCols(pd.ColumnPair...)
		if err != nil {
			return err
		}

		path := filepath.Join("/", "tmp", fmt.Sprintf("%x.csv", rand.Intn(100000000)))
		_, err = WriteCsv(newDf, path, true)
		if err != nil {
			return err
		}

		var usingBuf bytes.Buffer
		for _, opt := range pd.Opts {
			str := opt.createCmdString()
			if opt.getOption() == "using" {
				usingBuf.WriteString(str)
			}
		}

		cmdStringPiece := fmt.Sprintf(`"%s" %s,`, path, usingBuf.String())
		cmdString += cmdStringPiece
	}

	cmd := exec.Command("gnuplot", "-persist", "-e", cmdString)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err, cmd.Stderr))
	}

	return nil
}

// Fit calculates the line of best fit.
// ff is the fitting function.
// pd is the PlotData you would like to fit. Only the data pd.Df and pd.ColumnPair will be used.
// Pass options such as `using` or `via` in opts.
func Fit(ff string, pd PlotData, opts ...GnuplotOpt) error {
	rand.Seed(time.Now().UnixNano())
	newDf, err := pd.Df.LocCols(pd.ColumnPair...)
	if err != nil {
		return err
	}

	path := filepath.Join("/", "tmp", fmt.Sprintf("%x.csv", rand.Intn(100000000)))
	_, err = WriteCsv(newDf, path, true)
	if err != nil {
		return err
	}

	var usingBuf, viaBuf bytes.Buffer
	for _, opt := range pd.Opts {
		str := opt.createCmdString()
		if opt.getOption() == "using" {
			usingBuf.WriteString(str)
		}
		if opt.getOption() == "via" {
			viaBuf.WriteString(str)
		}
	}

	cmdString := fmt.Sprintf(`%s "%s" %s %s`, "fit f(x)", path, usingBuf.String(), viaBuf.String())
	cmd := exec.Command("gnuplot", "-persist", "-e", cmdString)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err, cmd.Stderr))
	}

	fmt.Println(cmd.Stdout)

	return nil
}
