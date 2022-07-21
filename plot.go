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

// A PlotData holds the data required for plotting.
//
// If you want to plot an arbitrary function, leave Df and Columns as nil.
// Otherwise, populate Df and Columns, and leave Function as "".
type PlotData struct {
	// Df is the DataFrame object you would like to plot.
	Df *DataFrame

	// Columns are the columns in Df that you want to plot. Usually, it's a pair of columns [xcol, ycol].
	// If you want to create a bar graph or a histogram, you can add more columns.
	Columns []string

	// Function is an arbitrary function such as sin(x) or an equation of the line of best fit.
	Function string

	// Opts are options such as `using` or `with`. `set` is passed in as an argument for other plotting functions.
	Opts []GnuplotOpt
}

// Plot plots a set of data given by the PlotData object `pd`.
//
// Pass in any `set` options you need. Refer to the [gnuplot documentation] for `set` options.
//
// [gnuplot documentation]: http://gnuplot.info/docs_5.5/loc9418.html
func Plot(pd PlotData, setOpts ...GnuplotOpt) error {
	path := ""

	if pd.Function != "" && pd.Df == nil && pd.Columns == nil {
		path = pd.Function
	} else {
		rand.Seed(time.Now().UnixNano())
		newDf, err := pd.Df.LocCols(pd.Columns...)
		if err != nil {
			return err
		}

		path = filepath.Join("/", "tmp", fmt.Sprintf("%x.csv", rand.Intn(100000000)))
		_, err = WriteCsv(newDf, path, true)
		if err != nil {
			return err
		}
	}

	var setBuf bytes.Buffer
	for _, setOpt := range setOpts {
		str := setOpt.createCmdString()
		setBuf.WriteString(str)
		setBuf.WriteString("; ")
	}

	var usingBuf bytes.Buffer
	var withBuf bytes.Buffer
	for _, opt := range pd.Opts {
		str := opt.createCmdString()
		switch opt.getOption() {
		case "using":
			usingBuf.WriteString(str)
		case "with":
			withBuf.WriteString(str)
		default:
			return fmt.Errorf("this option is not supported yet")
		}
	}

	cmdString := fmt.Sprintf(`%s %s "%s" %s %s`, setBuf.String(), "plot", path, usingBuf.String(), withBuf.String())
	cmd := exec.Command("gnuplot", "-persist", "-e", cmdString)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err, cmd.Stderr))
	}

	return nil
}

// PlotN plots several PlotData objects `pd` in one graph.
// Use PlotN when you want to compare two different datasets,
// or a dataset with a line of best fit.
//
// Refer to the [gnuplot documentation] for `set` options.
//
// [gnuplot documentation]: http://gnuplot.info/docs_5.5/loc9418.html
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
		path := ``

		if pd.Function != "" && pd.Df == nil && pd.Columns == nil {
			path = pd.Function
		} else {
			newDf, err := pd.Df.LocCols(pd.Columns...)
			if err != nil {
				return err
			}

			path = filepath.Join("/", "tmp", fmt.Sprintf("%x.csv", rand.Intn(100000000)))
			_, err = WriteCsv(newDf, path, true)
			if err != nil {
				return err
			}
			path = fmt.Sprintf(`"%s"`, path)
		}

		var usingBuf bytes.Buffer
		var withBuf bytes.Buffer
		for _, opt := range pd.Opts {
			str := opt.createCmdString()
			switch opt.getOption() {
			case "using":
				usingBuf.WriteString(str)
			case "with":
				withBuf.WriteString(str)
			default:
				return fmt.Errorf("this option is not supported yet")
			}
		}

		cmdStringPiece := fmt.Sprintf(`%s %s %s,`, path, usingBuf.String(), withBuf.String())
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

// Fit fits a user-defined function ff to data given in PlotData pd,
// and prints out the results.
//
// Pass options such as `using` in pd, but `via` in viaOpts.
func Fit(ff string, pd PlotData, viaOpts ...GnuplotOpt) error {
	rand.Seed(time.Now().UnixNano())
	newDf, err := pd.Df.LocCols(pd.Columns...)
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
	}
	for _, opt := range viaOpts {
		str := opt.createCmdString()
		if opt.getOption() == "via" {
			viaBuf.WriteString(str)
		}
	}

	cmdString := fmt.Sprintf(`%s %s "%s" %s %s`, "fit", ff, path, usingBuf.String(), viaBuf.String())
	cmd := exec.Command("gnuplot", "-persist", "-e", cmdString)
	combOutput, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", combOutput)

	return nil
}
