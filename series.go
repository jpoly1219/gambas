package gambas

import (
	"fmt"
	"math"
	"sort"
)

type Series struct {
	data  []interface{}
	index IndexData
	name  string
}

func (s Series) PrintSeries() string {
	message := fmt.Sprintln("data:", s.data, "\nindexArray:", s.index, "\nname:", s.name)
	fmt.Println(message)
	return message
}

// Head prints the first n items in the series.
func (s Series) Head(howMany int) {
	fmt.Println(s.index.names, s.name)
	for i := 0; i < howMany; i++ {
		fmt.Print(s.index.index[i], " ", s.data[i])
		fmt.Println()
	}
}

// At() returns an element at a given index.
// For multiindex, you need to pass in the whole index tuple.
func (s Series) At(in Index) (interface{}, error) {
	for i, index := range s.index.index {
		isSame := true
		for j := 0; j < len(index); j++ {
			if index[j] != in[j] {
				isSame = false
				break
			}
		}
		if isSame {
			return s.data[i], nil
		}
	}
	return nil, fmt.Errorf("the given index does not match any of the index in the series: %v", in)
}

// IAt() returns an element at a given integer index.
func (s Series) IAt(in int) (interface{}, error) {
	if in >= len(s.data) {
		return nil, fmt.Errorf("index out of bounds: %v", in)
	}
	if in < 0 {
		return nil, fmt.Errorf("index can't be less than 0")
	}
	return s.data[in], nil
}

// Loc() returns a range of data at given rows.
func (s Series) Loc(in []Index) (*Series, error) {
	// This makes sure that each index passed are the same length.
	indexLength := len(in[0])
	fmt.Println(indexLength)
	for i, eachIndex := range in {
		if indexLength != len(eachIndex) {
			return nil, fmt.Errorf("index length does not match: %v, %v", in[i-1], eachIndex)
		}
	}

	allFiltered := make([]interface{}, 0)
	allFilteredIndex := IndexData{}

	for _, inputIndex := range in {
		filtered := make([]interface{}, 0)
		filteredIndex := IndexData{}

		for j, seriesIndex := range s.index.index {
			isSame := true
			for k := 0; k < indexLength; k++ {
				if inputIndex[k] != seriesIndex[k] {
					isSame = false
					break
				}
			}
			if isSame {
				filtered = append(filtered, s.data[j])
				filteredIndex.index = append(filteredIndex.index, seriesIndex)
			}
		}

		if len(filtered) == 0 {
			return nil, fmt.Errorf("no data found for index %v", inputIndex)
		}
		allFiltered = append(allFiltered, filtered...)
		allFilteredIndex.index = append(allFilteredIndex.index, filteredIndex.index...)
	}
	allFilteredIndex.names = s.index.names

	result, err := NewSeries(allFiltered, s.name, &allFilteredIndex)
	if err != nil {
		return nil, err
	}
	return result, nil

	// result := make([]interface{}, 0)
	// for _, row := range rows {
	// 	atData, err := s.At(row)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	result = append(result, atData)
	// }

	// return result, nil
}

// ILoc() returns an array of elements at a given integer index range.
func (s Series) ILoc(min, max int) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for i := min; i < max; i++ {
		iatData, err := s.IAt(i)
		if err != nil {
			return nil, err
		}
		result = append(result, iatData)
	}

	return result, nil
}

// Summary statistics functions

// Count() counts the number of non-NA elements in a column.
func (s Series) Count() int {
	count := 0
	for _, v := range s.data {
		if v != nil || v != math.NaN() {
			count++
		}
	}

	return count
}

// Mean() returns the mean of the elements in a column.
func (s Series) Mean() (float64, error) {
	mean := 0.0

	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return math.NaN(), fmt.Errorf("no elements in this column")
	}

	for _, v := range data {
		mean += v
	}

	mean /= float64(len(data))
	roundedMean := math.Round(mean*1000) / 1000

	return roundedMean, nil
}

// Median() returns the median of the elements in a column.
func (s Series) Median() (float64, error) {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return math.NaN(), fmt.Errorf("no elements in this column")
	}
	if total%2 == 0 {
		lower := data[total/2-1]
		upper := data[total/2]

		median := (lower + upper) / 2
		roundedMedian := math.Round(median*1000) / 1000

		return roundedMedian, nil
	} else {
		median := data[(total+1)/2-1]
		roundedMedian := math.Round(median*1000) / 1000

		return roundedMedian, nil
	}
}

// Std() returns the sample standard deviation of the elements in a column.
func (s Series) Std() (float64, error) {
	std := 0.0
	mean, err := s.Mean() // this also checks that all data can be converted to float64.
	if err != nil {
		return math.NaN(), err
	}

	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	numerator := 0.0
	for _, v := range data {
		temp := math.Pow(v-mean, 2)
		numerator += temp
	}
	std = math.Sqrt(numerator / float64(len(data)-1))
	roundedStd := math.Round(std*1000) / 1000

	return roundedStd, nil
}

// Min() returns the smallest element in a column.
func (s Series) Min() (float64, error) {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return math.NaN(), fmt.Errorf("no elements in this column")
	}

	return data[0], nil
}

// Max() returns the largest element is a column.
func (s Series) Max() (float64, error) {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return math.NaN(), fmt.Errorf("no elements in this column")
	}

	return data[total-1], nil
}

// Q1() returns the lower quartile (25%) of the elements in a column.
// This does not include the median during calculation.
func (s Series) Q1() (float64, error) {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	if len(data)%2 == 0 {
		lower := data[:len(data)/2]
		q1, err := median(lower)
		if err != nil {
			return math.NaN(), err
		}
		return q1, nil
	} else {
		lower := data[:(len(data)-1)/2]
		q1, err := median(lower)
		if err != nil {
			return math.NaN(), err
		}
		return q1, nil
	}
}

// Q2() returns the middle quartile (50%) of the elements in a column.
// This accomplishes the same thing as s.Median().
func (s Series) Q2() (float64, error) {
	q2, err := s.Median()
	if err != nil {
		return math.NaN(), err
	}
	return q2, nil
}

// Q3() returns the upper quartile (75%) of the elements in a column.
// This does not include the median during calculation.
func (s Series) Q3() (float64, error) {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return math.NaN(), err
	}
	sort.Float64s(data)

	if len(data)%2 == 0 {
		upper := data[len(data)/2:]
		q3, err := median(upper)
		if err != nil {
			return math.NaN(), err
		}
		return q3, nil
	} else {
		upper := data[(len(data)+1)/2:]
		q3, err := median(upper)
		if err != nil {
			return math.NaN(), err
		}
		return q3, nil
	}
}

func (s Series) Describe() ([]interface{}, error) {
	count := s.Count()
	fmt.Println("Count:", count)

	mean, err := s.Mean()
	if err != nil {
		return nil, err
	}
	fmt.Println("Mean:", mean)

	median, err := s.Median()
	if err != nil {
		return nil, err
	}
	fmt.Println("Median:", median)

	std, err := s.Std()
	if err != nil {
		return nil, err
	}
	fmt.Println("Std:", std)

	min, err := s.Min()
	if err != nil {
		return nil, err
	}
	fmt.Println("Min:", min)

	max, err := s.Max()
	if err != nil {
		return nil, err
	}
	fmt.Println("Max:", max)

	q1, err := s.Q1()
	if err != nil {
		return nil, err
	}
	fmt.Println("Q1:", q1)

	q2, err := s.Q2()
	if err != nil {
		return nil, err
	}
	fmt.Println("Q2:", q2)

	q3, err := s.Q3()
	if err != nil {
		return nil, err
	}
	fmt.Println("Q3:", q3)

	result := []interface{}{count, mean, median, std, min, max, q1, q2, q3}

	return result, nil
}

// SortByIndex() sorts the elements in a series by the index.
// Multiindex support is coming, but this may require an overhaul.
func (s *Series) SortByIndex(ascending bool) error {
	indDatMap := make(map[string]interface{})
	for i, data := range s.data {
		key, err := s.index.index[i].hashKey()
		if err != nil {
			return err
		}
		indDatMap[*key] = data
	}

	fmt.Println(s.index)

	if ascending {
		sort.Sort(s.index)
	} else {
		sort.Sort(sort.Reverse(s.index))
	}

	fmt.Println(s.index)

	for i, index := range s.index.index {
		key, err := index.hashKey()
		if err != nil {
			return err
		}
		s.data[i] = indDatMap[*key]
	}

	return nil
}
