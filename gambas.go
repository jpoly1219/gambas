package gambas

import "fmt"

// CreateRangeIndex takes the length of an Index and creates a RangeIndex.
// RangeIndex is an index that spans from 0 to the length of the index.
func CreateRangeIndex(length int) IndexData {
	rangeIndex := make([]Index, length)

	for i := 0; i < length; i++ {
		rangeIndex[i] = Index{i}
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
func NewSeries(data []interface{}, name string, index *IndexData) (*Series, error) {
	var s Series
	ok, err := checkTypeIntegrity(data)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("types do not match")
	}
	s.data = data

	if index == nil {
		s.index = CreateRangeIndex(len(data))
	} else {
		s.index.index = make([]Index, len(index.index))
		s.index.names = make([]string, len(index.names))
		copy(s.index.index, index.index)
		copy(s.index.names, index.names)
	}

	s.name = name
	return &s, nil
}

// NewDataFrame created a new DataFrame object from given parameters.
// Generally, NewDataFrameFromFile will be used more often.
func NewDataFrame(data [][]interface{}, columns []string, indexCols []string) (*DataFrame, error) {
	if len(data) != len(columns) {
		return nil, fmt.Errorf("length of data (%d) and columns (%d) does not match", len(data), len(columns))
	}

	var df DataFrame
	df.series = make([]Series, len(data))
	df.index = IndexData{}
	df.index.names = indexCols
	df.columns = columns

	// create df.index
	// find location of index column
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
			indexTuple = append(indexTuple, data[location][i])
		}

		df.index.index = append(df.index.index, indexTuple)
	}

	for i, v := range data {
		series, err := NewSeries(v, fmt.Sprint(columns[i]), &df.index)
		if err != nil {
			return nil, err
		}
		df.series[i] = *series
	}

	return &df, nil
}
