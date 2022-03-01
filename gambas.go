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
func NewSeries(data []interface{}, index []interface{}, name string) Series {
	var s Series
	s.data = make(map[interface{}]interface{})
	s.index.data = make([]interface{}, len(index))

	for i, v := range data {
		if index == nil {
			s.index.data = append(s.index.data, i)
			s.data[i] = v
		} else {
			s.index.data = index
			s.data[index[i]] = v
		}
	}

	s.name = name
	fmt.Println(s)

	return s
}

// NewDataFrame created a new DataFrame object from given parameters.
// Generally, NewDataFrameFromFile will be used more often.
func NewDataFrame(data [][]interface{}, index []interface{}, columns []interface{}) DataFrame {
	var df DataFrame
	df.series = make(map[interface{}]Series, len(data))
	df.index.data = make([]interface{}, len(index))
	df.columns.data = make([]interface{}, len(columns))

	for i, v := range data {
		df.series[columns[i]] = NewSeries(v, index, columns[i].(string))
	}
	df.index.data = index
	df.columns.data = columns

	return df
}