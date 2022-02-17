package gambas

func NewSeries(data []interface{}, indexes CustomIndex, name string) (Series, error) {
	var s Series
	s.Data = make(map[int]interface{})

	if indexes.Value != nil {
		s.CustomIndex = indexes
	}

	for i, v := range data {
		s.Data[i] = v
	}

	s.Name = name

	return s, nil
}
