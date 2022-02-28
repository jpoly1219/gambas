package gambas

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ReadCsv reads a CSV file and returns a new DataFrame object.
// Path to file should be generated using `filepath.Join`.
func ReadCsv(pathToFile string, dtype []string) DataFrame {
	// read line by line
	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
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
				err = nil
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
		dataArray := make([]interface{}, len(row))
		for _, v := range row {
			// add to data2DArray
			dataArray = append(dataArray, v)
		}
		data2DArray = append(data2DArray, dataArray)
		rowNum++
	}
	// create new DataFrame object and return it
	df := NewDataFrame(data2DArray, []interface{}{}, columnArray)
	return df
	// each data should be checked to see what type it is
}
