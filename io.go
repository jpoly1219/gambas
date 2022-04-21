package gambas

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// ReadCsv reads a CSV file and returns a new DataFrame object.
// It is recommended to generate pathToFile using `filepath.Join`.
// TODO: users should be able to define custom indices.
func ReadCsv(pathToFile string, indexCols []string) (*DataFrame, error) {
	// read line by line
	f, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvr := csv.NewReader(f)

	rowNum := 0
	columnArray := make([]string, 0)
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
			columnArray = append(columnArray, row...)
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
			vChecked := checkCSVDataType(v)
			// fmt.Print(vChecked, ", ")
			data2DArray[i] = append(data2DArray[i], vChecked)
		}
		rowNum++
	}
	// create new DataFrame object and return it
	if indexCols != nil {
		df, err := NewDataFrame(data2DArray, columnArray, indexCols)
		if err != nil {
			return nil, err
		}

		return df, nil
	}
	df, err := NewDataFrame(data2DArray, columnArray, []string{columnArray[0]})
	if err != nil {
		return nil, err
	}

	return df, nil
}

// WriteCsv writes a DataFrame object to CSV file.
// It is recommended to generate pathToFile using `filepath.Join`.
func WriteCsv(df DataFrame, pathToFile string) (os.FileInfo, error) {
	f, err := os.Create(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	// write column names in the first row
	for i, col := range df.columns {
		_, err := w.WriteString(col)
		if err != nil {
			return nil, err
		}

		if i+1 != len(df.columns) {
			_, err := w.WriteString(",")
			if err != nil {
				return nil, err
			}
		}
	}
	w.WriteString("\n")

	// write the data in the following rows
	for i := range df.series[0].data {
		for j, ser := range df.series {
			_, err := w.WriteString(fmt.Sprint(ser.data[i]))
			if err != nil {
				return nil, err
			}

			if j+1 != len(df.series) {
				_, err := w.WriteString(",")
				if err != nil {
					return nil, err
				}
			}
		}

		w.WriteString("\n")
	}

	w.Flush()

	info, err := os.Stat(pathToFile)
	if err != nil {
		return nil, err
	}
	return info, nil
}
