package gambas

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"path/filepath"
)

// Plotting functionality uses gnuplot as its backend.

func (df *DataFrame) Plot(xcol, ycol string, opts ...GnuplotOpt) error {
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
