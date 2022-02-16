package gambas

func NewSeries(data []Element, indexes []Index, name string) (Series, error) {
	var s Series
	s.Elements = data

	if indexes == nil {
		for i := range data {
			index := Index{i}
			indexes = append(indexes, index)
		}
		s.Indexes = indexes
	} else {
		s.Indexes = indexes
	}

	s.Name = name

	return s, nil
}
