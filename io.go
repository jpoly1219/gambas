package gambas

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ReadCsv reads a CSV file and returns a new DataFrame object.
// Path to file should be generated using `filepath.Join`.
func ReadCsv(pathToFile string) (*DataFrame, error) {
	// read line by line
	f, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvr := csv.NewReader(f)

	rowNum := 0
	columnArray := make([]interface{}, 0)
	data2DArray := make([][]interface{}, 0)
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		// first line is column name
		if rowNum == 0 {
			// add to columnArray
			for _, v := range row {
				columnArray = append(columnArray, v)
			}
			rowNum++
			continue
		}
		// second line onwards is the actual data
		for i, v := range row {
			// add to data2DArray
			if len(data2DArray) < len(row) {
				data2DArray = append(data2DArray, make([]interface{}, 0))
			}
			// each data should be checked to see what type it is
			vChecked := checkType(v)
			data2DArray[i] = append(data2DArray[i], vChecked)
		}
		rowNum++
	}
	// create new DataFrame object and return it
	df, err := NewDataFrame(data2DArray, columnArray, []interface{}{columnArray[0]})
	if err != nil {
		return nil, err
	}

	return df, nil
}
