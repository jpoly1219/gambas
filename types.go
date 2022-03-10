package gambas

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// checkType checks to see if the data can be represented as a float64.
// Because CSV is read as an array of strings, there has to be a way to check the type.
func checkType(data interface{}) interface{} {
	if data.(string) == "NaN" {
		fmt.Println("NaN detected")
		return math.NaN()
	}
	v, ok := strconv.ParseFloat(data.(string), 64)
	switch ok {
	case nil:
		return v
	default:
		return data.(string)
	}
}

type Index struct {
	data []interface{}
}

type Series struct {
	data  []interface{}
	index Index
	name  string
}

func (s Series) PrintSeries() string {
	message := fmt.Sprintln("data:", s.data, "\nindexArray:", s.index, "\nname:", s.name)
	fmt.Println(message)
	return message
}

func (s Series) CalcMean() (float64, error) {
	sum := 0.0
	for _, v := range s.data {
		if reflect.ValueOf(v).Kind() != reflect.Float64 || reflect.ValueOf(v).Kind() != reflect.Int {
			err := fmt.Errorf("type is not int or float64: %T", v)
			fmt.Println(err)
			return 0.0, err
		}
		sum += v.(float64)
	}
	mean := sum / float64(len(s.data))

	return mean, nil
}

// At() returns the element at a given index.
func (s Series) Loc(index interface{}) (interface{}, error) {
	for i, v := range s.index.data {
		if v == index {
			result := s.data[i]
			return result, nil
		}
	}

	return nil, fmt.Errorf("index %s is not found", index)
}

// AtM() returns an array of elements at given indexes.
func (s Series) LocM(indexArray []interface{}) ([]interface{}, error) {
	resultArray := make([]interface{}, len(indexArray))

	for i, v := range indexArray {
		result, err := s.Loc(v)
		if err != nil {
			return nil, err
		}
		resultArray[i] = result
	}

	return resultArray, nil
}

// AtR() returns an array of elements at a given index range.
func (s Series) LocR(min, max int) ([]interface{}, error) {
	resultArray := make([]interface{}, 0)

	for i := min; i < max; i++ {
		key := s.index.data[i]
		result, err := s.Loc(key)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, result)
	}

	return resultArray, nil
}

type DataFrame struct {
	series  []Series
	index   Index
	columns Index
}

// LocRows returns a set of rows as a new DataFrame object, given a list of labels.
func (df DataFrame) LocRows(rows []interface{}) (*DataFrame, error) {
	filtered2D := make([][]interface{}, 0)
	for _, series := range df.series {
		filtered := make([]interface{}, 0)
		for _, label := range rows {
			for j, index := range df.index.data {
				if label == index {
					filtered = append(filtered, series.data[j])
				}
			}
		}
		filtered2D = append(filtered2D, filtered)
	}

	dataframe, err := NewDataFrame(filtered2D, Index{rows}, df.columns.data)
	if err != nil {
		return nil, err
	}

	return dataframe, nil
}

// LocRows returns a set of columns as a new DataFrame object, given a list of labels.
func (df DataFrame) LocCols(cols []interface{}) (*DataFrame, error) {
	filtered2D := make([][]interface{}, 0)
	for _, column := range cols {
		for _, series := range df.series {
			if series.name == column {
				filtered2D = append(filtered2D, series.data)
			}
		}
	}

	dataframe, err := NewDataFrame(filtered2D, df.index, cols)
	if err != nil {
		return nil, err
	}

	return dataframe, nil
}

// Loc indexes the DataFrame object given a single row or column label.
func (df DataFrame) Loc(rows, cols []interface{}) (*DataFrame, error) {
	filteredCols := make([]Series, len(cols))
	for i, col := range cols {
		for _, series := range df.series {
			if series.name == col {
				filteredCols[i] = series
			}
		}
	}

	filtered2D := make([][]interface{}, len(cols))
	for i, filteredCol := range filteredCols {
		filteredRows := make([]interface{}, len(rows))

		for j, row := range rows {
			for k, index := range filteredCol.index.data {
				if index == row {
					filteredRows[j] = filteredCol.data[k]
				}
			}
		}

		filtered2D[i] = filteredRows
	}

	switch rows[0].(type) {
	case int:
		dataframe, err := NewDataFrame(filtered2D, CreateRangeIndex(len(rows)), cols)
		if err != nil {
			return nil, err
		}
		return dataframe, nil
	default:
		dataframe, err := NewDataFrame(filtered2D, Index{rows}, cols)
		if err != nil {
			return nil, err
		}
		return dataframe, nil
	}

}

// func (df DataFrame) LocM() (*DataFrame, error) {

// }

// func (df DataFrame) ILoc() (*DataFrame, error) {

// }

// I am thinking of redefining some of DataFrame at the moment.
// Notably, adding indexCol method of type []interface{} and changing the type of index to []Index.
// This is to make indexing make more sense, and allow for multi-indexing.
// There will be a major overhaul for this which I will work on during the weekend.
