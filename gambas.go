package gambas

func NewSeries(data []interface{}, indexes []Index, name string) (Series, error) {
	var s Series
	s.Data = make(map[Index]interface{})

	for i, v := range data {
		if indexes == nil {
			s.Data[Index{i}] = v
		} else {
			s.Data[indexes[i]] = v
		}
	}

	s.Name = name

	return s, nil
}
