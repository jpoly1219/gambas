package gambas

import "fmt"

func NewSeries(data []interface{}, indexes CustomIndex, name string) (Series, error) {
	var s Series
	s.Data = make(map[interface{}]interface{})

	if indexes.Value != nil {
		s.CustomIndex = indexes
		for i, v := range data {
			s.Data[indexes.Value[i]] = v
		}
	} else {
		s.CustomIndex.Value = make([]interface{}, len(data))
		for i, v := range data {
			s.CustomIndex.Value[i] = i
			s.Data[i] = v
		}
	}

	s.Name = name
	fmt.Println(s)

	return s, nil
}
