package gambas

import "fmt"

func NewSeries(data []interface{}, indexes []interface{}, name string) Series {
	var s Series
	s.data = make(map[interface{}]interface{})
	s.index.data = make([]interface{}, len(indexes))

	for i, v := range data {
		if indexes == nil {
			s.index.data = append(s.index.data, i)
			s.data[i] = v
		} else {
			s.index.data = indexes
			s.data[indexes[i]] = v
		}
	}

	s.name = name
	fmt.Println(s)

	return s
}
