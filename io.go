package gambas

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/xuri/excelize/v2"
)

// ReadCsv reads a CSV file and returns a new DataFrame object.
// It is recommended to generate pathToFile using `filepath.Join`.
func ReadCsv(pathToFile string, indexCols []string) (DataFrame, error) {
	// read line by line
	f, err := os.Open(pathToFile)
	if err != nil {
		return DataFrame{}, err
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
			vChecked := tryDataType(v)
			data2DArray[i] = append(data2DArray[i], vChecked)
		}
		rowNum++
	}
	// create new DataFrame object and return it
	df, err := NewDataFrame(data2DArray, columnArray, indexCols)
	if err != nil {
		return DataFrame{}, err
	}

	return df, nil
}

// WriteCsv writes a DataFrame object to CSV file.
// It is recommended to generate pathToFile using `filepath.Join`.
func WriteCsv(df DataFrame, pathToFile string, skipColumnLabel bool) (os.FileInfo, error) {
	f, err := os.Create(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	// write column names in the first row
	if !skipColumnLabel {
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
	}

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

// ReadJson reads a JSON file and returns a new DataFrame object.
// It is recommended to generate pathToFile using `filepath.Join`.
// The JSON file should be in this format:
// {"col1":[val1, val2, ...], "col2":[val1, val2, ...], ...}
// You can either set a column to be the index, or set it as nil.
// If nil, a new RangeIndex will be created.
// Your index column should not have any missing values.
// Order of columns is not guaranteed, but the index column will always come first.
func ReadJsonByColumns(pathToFile string, indexCols []string) (DataFrame, error) {
	f, err := os.Open(pathToFile)
	if err != nil {
		return DataFrame{}, err
	}
	defer f.Close()

	fbyte, err := io.ReadAll(f)
	if err != nil {
		return DataFrame{}, err
	}

	var decoded map[string]interface{}
	err = json.Unmarshal(fbyte, &decoded)
	if err != nil {
		return DataFrame{}, err
	}

	newDfData := make([][]interface{}, 0)
	newDfCols := make([]string, 0)
	newDfIndexNames := make([]string, 0)

	for col, colData := range decoded {
		newDfCols = append(newDfCols, col)
		colDataAsserted := colData.([]interface{})
		for i, cda := range colDataAsserted {
			checked := checkJsonDataType(cda)
			colDataAsserted[i] = checked
		}
		newDfData = append(newDfData, colDataAsserted)

		if indexCols != nil && containsString(indexCols, col) {
			newDfIndexNames = append(newDfIndexNames, col)
		}
	}

	if indexCols == nil {
		newDfIndexNames = nil
	}

	newDf, err := NewDataFrame(newDfData, newDfCols, newDfIndexNames)
	if err != nil {
		return DataFrame{}, err
	}
	newDf.SortByColumns()
	newDf.SortByIndex(true)
	newDf.SortIndexColFirst()
	return newDf, nil
}

// ReadJsonStream reads a JSON stream and returns a new DataFrame object.
// The JSON file should be in this format:
// {"col1":val1, "col2":val2, ...}{"col1":val1, "col2":val2, ...}
func ReadJsonStream(pathToFile string, indexCols []string) (DataFrame, error) {
	f, err := os.Open(pathToFile)
	if err != nil {
		return DataFrame{}, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	newDfCols := make([]string, 0)
	colData := make(map[string][]interface{}, 0)

	_, err = dec.Token()
	if err != nil {
		return DataFrame{}, err
	}
	for dec.More() {
		var row map[string]interface{}
		err := dec.Decode(&row)
		if err == io.EOF {
			break
		} else if err != nil {
			return DataFrame{}, err
		}

		for k, v := range row {
			colData[k] = append(colData[k], checkJsonDataType(v))
		}
	}
	_, err = dec.Token()
	if err != nil {
		return DataFrame{}, err
	}

	newDfData := make([][]interface{}, 0)

	for k, v := range colData {
		newDfCols = append(newDfCols, k)
		newDfData = append(newDfData, v)
	}

	newDf, err := NewDataFrame(newDfData, newDfCols, indexCols)
	if err != nil {
		return DataFrame{}, err
	}
	newDf.SortByColumns()
	newDf.SortByIndex(true)
	newDf.SortIndexColFirst()
	return newDf, nil
}

// WriteJson writes a DataFrame object to a file.
func WriteJson(df DataFrame, pathToFile string) (os.FileInfo, error) {
	f, err := os.Create(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	jsonByte, err := json.MarshalIndent(&df, "", "\t")
	if err != nil {
		return nil, err
	}
	w.Write(jsonByte)
	w.Flush()

	info, err := os.Stat(pathToFile)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// ReadExcel reads an excel file and converts it to a DataFrame object.
// The axis depends on the layout of the data.
// Row-based data where each group represents a row will have an axis=0.
// Column-based data where each group represents a column will have an axis=1.
func ReadExcel(pathToFile, sheetName string, axis int) (DataFrame, error) {
	f, err := excelize.OpenFile(pathToFile)
	if err != nil {
		return DataFrame{}, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	newDfData := make([][]interface{}, 0)
	newDfCols := make([]string, 0)

	if axis == 0 {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return DataFrame{}, err
		}
		longestRowLength := 0
		for _, row := range rows {
			if longestRowLength < len(row)-1 {
				longestRowLength = len(row) - 1
			}
			d := make([]interface{}, 0)

			for i, cellValue := range row {
				if i == 0 {
					newDfCols = append(newDfCols, cellValue)
				} else {
					vChecked := tryDataType(cellValue)
					d = append(d, vChecked)
				}
			}

			if len(d) < longestRowLength {
				for i := 0; i < longestRowLength-len(d); i++ {
					d = append(d, math.NaN())
				}
			}
			newDfData = append(newDfData, d)
		}
	}

	if axis == 1 {
		cols, err := f.GetCols(sheetName)
		if err != nil {
			return DataFrame{}, err
		}
		for _, col := range cols {
			d := make([]interface{}, 0)
			for i, cellValue := range col {
				if i == 0 {
					newDfCols = append(newDfCols, cellValue)
				} else {
					vChecked := tryDataType(cellValue)
					d = append(d, vChecked)
				}
			}
			newDfData = append(newDfData, d)
		}
	}

	newDf, err := NewDataFrame(newDfData, newDfCols, nil)
	if err != nil {
		return DataFrame{}, err
	}
	return newDf, nil
}

// WriteExcel writes a DataFrame object into an Excel file.
func WriteExcel(df DataFrame, pathToFile string) (os.FileInfo, error) {
	f := excelize.NewFile()
	sheetIndex := f.NewSheet("Sheet1")

	for i, col := range df.columns {
		coord := fmt.Sprintf("%s1", generateAlphabets(i+1))
		f.SetCellValue("Sheet1", coord, col)
	}

	for i, ser := range df.series {
		xCoord := generateAlphabets(i + 1)
		for j, data := range ser.data {
			coord := fmt.Sprintf("%s%d", xCoord, j+2)
			f.SetCellValue("Sheet1", coord, data)
		}
	}

	f.SetActiveSheet(sheetIndex)

	err := f.SaveAs(pathToFile)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(pathToFile)
	if err != nil {
		return nil, err
	}
	return info, nil
}
