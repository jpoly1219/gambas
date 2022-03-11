package gambas

import "fmt"

// CreateRangeIndex takes the length of an Index and creates a RangeIndex.
// RangeIndex is an index that spans from 0 to the length of the index.
func CreateRangeIndex(length int) Index {
	zeroToLength := make([]interface{}, length)

	for i := 0; i < length; i++ {
		zeroToLength[i] = i
	}

	return Index{zeroToLength}
}

// NewSeries created a new Series object from given parameters.
// Generally, NewSeriesFromFile will be used more often.
func NewSeries(data []interface{}, name string) (*Series, error) {
	var s Series
	s.data = data
	s.index = CreateRangeIndex(len(data))
	s.name = name
	fmt.Println(s)

	return &s, nil
}

// NewDataFrame created a new DataFrame object from given parameters.
// Generally, NewDataFrameFromFile will be used more often.
func NewDataFrame(data [][]interface{}, columns []interface{}, indexCols []interface{}) (*DataFrame, error) {
	if len(data) != len(columns) {
		return nil, fmt.Errorf("length of data (%d) and columns (%d) does not match", len(data), len(columns))
	}

	var df DataFrame
	df.series = make([]Series, len(data))
	df.index = make([]Index, 0)
	df.columns.data = columns
	df.indexCols = indexCols

	// create df.index
	// find location of index column
	for i, col := range columns {
		for _, indexCol := range indexCols {
			if col == indexCol {
				df.index = append(df.index, Index{data[i]})
			}
		}
	}

	for i, v := range data {
		series, err := NewSeries(v, fmt.Sprint(columns[i]))
		if err != nil {
			return nil, err
		}
		df.series[i] = *series
	}

	return &df, nil
}
