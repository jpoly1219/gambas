package gambas

import "fmt"

func NewSeries(data []interface{}, indexes []Index, name string) Series {
	var s Series
	s.Data = make(map[Index]interface{})
	s.IndexArray = make([]Index, len(data))

	for i, v := range data {
		if indexes == nil {
			s.IndexArray[i] = Index{i}
			s.Data[Index{i}] = v
		} else {
			s.IndexArray = indexes
			s.Data[indexes[i]] = v
		}
	}

	s.Name = name
	fmt.Println(s)

	return s
}
