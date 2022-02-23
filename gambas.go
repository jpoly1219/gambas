package gambas

import "fmt"

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

func NewSeriesFromFile(pathToFile ) {

}

func NewDataFrame(data [][]interface{}, index []interface{}, columns []interface{}, name string) DataFrame {
	var df DataFrame
	df.series = make(map[interface{}]Series, len(data))
	df.index.data = make([]interface{}, len(index))
	df.columns.data = make([]interface{}, len(columns))

	for i, v := range data {
		df.series[columns[i]] = NewSeries(v, index, columns[i].(string))
	}

	return df
}

func NewDataFrameFromFile() {

}
