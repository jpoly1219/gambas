package gambas

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// CreateRangeIndex takes the length of an Index and creates a RangeIndex.
// RangeIndex is an index that spans from 0 to the length of the index.
func CreateRangeIndex(length int) IndexData {
	rangeIndex := make([]Index, length)

	for i := 0; i < length; i++ {
		rangeIndex[i] = Index{i, []interface{}{i}}
	}

	indexData := IndexData{
		rangeIndex,
		[]string{""},
	}

	return indexData
}

// NewSeries created a new Series object from given parameters.
// Generally, NewSeriesFromFile will be used more often.
// The index parameter can be set to nil when calling NewSeries on its own.
// This field is for passing in the DataFrame's index data in NewDataFrame.
func NewSeries(data []interface{}, name string, index *IndexData) (Series, error) {
	var s Series
	dtype, err := checkTypeIntegrity(data)
	if err != nil {
		return Series{}, err
	}
	s.dtype = dtype

	switch dtype {
	case "float64":
		s.data = consolidateToFloat64(data)
	case "string":
		s.data = consolidateToString(data)
	default:
		if s.data == nil {
			s.data = data
		}
	}

	if index == nil {
		s.index = CreateRangeIndex(len(data))
	} else {
		s.index.index = make([]Index, len(index.index))
		s.index.names = make([]string, len(index.names))
		copy(s.index.index, index.index)
		copy(s.index.names, index.names)
	}

	s.name = name
	return s, nil
}

// NewDataFrame created a new DataFrame object from given parameters.
// Generally, NewDataFrameFromFile will be used more often.
func NewDataFrame(data [][]interface{}, columns []string, indexCols []string) (DataFrame, error) {
	if len(data) != len(columns) {
		return DataFrame{}, fmt.Errorf("length of data (%d) and columns (%d) does not match", len(data), len(columns))
	}

	var df DataFrame
	df.series = make([]Series, len(data))
	df.index = IndexData{}
	df.columns = columns

	// create df.index
	// find location of index column
	if indexCols == nil {
		df.index = CreateRangeIndex(len(data[0]))
	} else {
		df.index.names = indexCols
		indexColsIndex := make([]int, 0)

		for i, col := range columns {
			for _, indexCol := range indexCols {
				if col == indexCol {
					indexColsIndex = append(indexColsIndex, i)
				}
			}
		}

		for i := 0; i < len(data[0]); i++ {
			indexTuple := Index{}

			for _, location := range indexColsIndex {
				indexTuple.id = i
				indexTuple.value = append(indexTuple.value, data[location][i])
			}

			df.index.index = append(df.index.index, indexTuple)
		}
	}

	for i, v := range data {
		series, err := NewSeries(v, fmt.Sprint(columns[i]), &df.index)
		if err != nil {
			return DataFrame{}, err
		}
		df.series[i] = series
	}

	return df, nil
}

// NewIndexData creates a new IndexData object.
func NewIndexData(index [][]interface{}, names []string) (IndexData, error) {
	indexData := IndexData{}
	for i, val := range index {
		idx := Index{i, val}
		indexData.index = append(indexData.index, idx)
	}
	indexData.names = names

	return indexData, nil
}

/* Sample generator */

// NewSampleDataFrame creates a sample DataFrame object.
// colNameDtype defines each column's name and dtype.
// length controls the length of columns.
// You can choose certain columns to be indexes, or just leave this as nil.
func NewSampleDataFrame(colNameDtype map[string]string, length int, indexCols []string) (DataFrame, error) {
	rand.Seed(time.Now().UnixNano())

	newDfData := make([][]interface{}, 0)
	newDfColumns := make([]string, 0)
	for name, dtype := range colNameDtype {
		newDfColumns = append(newDfColumns, name)
		data := make([]interface{}, 0)

		switch dtype {
		case "bool":
			for i := 0; i < length; i++ {
				if rand.Intn(2) == 0 {
					data = append(data, false)
				} else {
					data = append(data, true)
				}
			}
		case "int":
			for i := 0; i < length; i++ {
				data = append(data, rand.Intn(100))
			}
		case "float64":
			for i := 0; i < length; i++ {
				data = append(data, 100*rand.Float64())
			}
		case "string":
			var sb strings.Builder
			charset := "abcdefghijklmnopqrstuvwxyz"
			for i := 0; i < length; i++ {
				for j := 0; j < 8; j++ {
					sb.WriteByte(charset[rand.Intn(len(charset))])
				}
				data = append(data, sb.String())
				sb.Reset()
			}
		default:
			return DataFrame{}, fmt.Errorf("invalid data type")
		}

		newDfData = append(newDfData, data)
	}

	newDf, err := NewDataFrame(newDfData, newDfColumns, indexCols)
	if err != nil {
		return DataFrame{}, err
	}

	return newDf, nil
}
