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
// For default index (0, ..., n-1), pass in CreateRangeIndex().
func NewSeries(data []interface{}, index Index, name string) (*Series, error) {
	if len(data) != len(index.data) {
		return nil, fmt.Errorf("length of data (%d) and index (%d) does not match", len(data), len(index.data))
	}

	var s Series
	s.data = make([]interface{}, len(index.data))
	s.index = index

	for i, v := range data {
		s.data[i] = v
	}

	s.name = name
	fmt.Println(s)

	return &s, nil
}

// NewDataFrame created a new DataFrame object from given parameters.
// Generally, NewDataFrameFromFile will be used more often.
// For default index (0, ..., n-1), pass in CreateRangeIndex().
func NewDataFrame(data [][]interface{}, index Index, columns []interface{}) (*DataFrame, error) {
	if len(data) != len(columns) {
		return nil, fmt.Errorf("length of data (%d) and columns (%d) does not match", len(data), len(columns))
	}

	var df DataFrame
	df.series = make([]Series, len(data))
	df.index = index
	df.columns.data = columns

	for i, v := range data {
		series, err := NewSeries(v, index, columns[i].(string))
		if err != nil {
			return nil, err
		}
		df.series[i] = *series
	}

	return &df, nil
}
