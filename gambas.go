package gambas

import "fmt"

func NewSeries(data []interface{}, indexes []Index, name string) Series {
	var s Series
	s.data = make(map[Index]interface{})
	s.indexArray = make([]Index, len(data))

	for i, v := range data {
		if indexes == nil {
			s.indexArray[i] = Index{i}
			s.data[Index{i}] = v
		} else {
			s.indexArray = indexes
			s.data[indexes[i]] = v
		}
	}

	s.name = name
	fmt.Println(s)

	return s
}
