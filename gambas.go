package gambas

import "fmt"

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

// NewSeriesFromFile creates a new Series object from the file provided.
// Path to file should be generated using `filepath.Join`.
func NewSeriesFromFile(pathToFile string) {

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

// NewDataFrameFromFile creates a new DataFrame object from the file provided.
// Path to file should be generated using `filepath.Join`.
func NewDataFrameFromFile(pathToFile string) {

}
