package gambas

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"path/filepath"
)

// Plotting functionality uses gnuplot as its backend.

func (df *DataFrame) Plot(xcol, ycol string) error {
	newDf, err := df.LocCols(xcol, ycol)
	if err != nil {
		return err
	}

	path := filepath.Join("/", "tmp", fmt.Sprintf("%x.csv", rand.Intn(100000000)))
	_, err = WriteCsv(newDf, path, true)
	if err != nil {
		return err
	}

	cmdString := fmt.Sprintf(`%s "%s" %s`, `set xdata time; set timefmt "%Y-%m-%d %H:%M:%S+%M:%S"; plot`, path, "using 1:2")
	cmd := exec.Command("gnuplot", "-persist", "-e", cmdString)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err, cmd.Stderr))
	}

	return nil
}
