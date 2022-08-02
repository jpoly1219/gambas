package gambas

import (
	"fmt"
	"math"
	"os"
	"sort"
	"text/tabwriter"
)

// A Series represents a column of data.
type Series struct {
	data  []interface{}
	index IndexData
	name  string
	dtype string
}

// Len is used to implement the sort.Sort interface.
func (s Series) Len() int {
	return len(s.data)
}

// Less is used to implement the sort.Sort interface.
func (s Series) Less(i, j int) bool {
	if fmt.Sprint(s.data[i]) == "NaN" {
		return false
	}

	if fmt.Sprint(s.data[j]) == "NaN" {
		return true
	}

	return fmt.Sprint(s.data[i]) < fmt.Sprint(s.data[j])
}

// Swap is used to implement the sort.Sort interface.
func (s Series) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Print prints all data in a Series object.
func (s *Series) Print() {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 5, 0, 4, ' ', 0)

	for i := range s.index.names {
		fmt.Fprint(w, s.index.names[i], "\t")
	}
	fmt.Fprint(w, "|", "\t")
	fmt.Fprintln(w, s.name, "\t")

	for i := 0; i < len(s.data); i++ {
		if len(s.index.index[i].value) > 1 {
			for j := range s.index.index[i].value {
				fmt.Fprint(w, s.index.index[i].value[j], "\t")
			}
		} else {
			fmt.Fprint(w, s.index.index[i].value[0], "\t")
		}

		fmt.Fprint(w, "|", "\t")
		fmt.Fprint(w, s.data[i], "\t")
		fmt.Fprintln(w)
	}
	w.Flush()
}

// PrintRange prints data in a Series object at a given range.
// Index starts at 0.
func (s *Series) PrintRange(start, end int) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 5, 0, 4, ' ', 0)

	for i := range s.index.names {
		fmt.Fprint(w, s.index.names[i], "\t")
	}
	fmt.Fprint(w, "|", "\t")
	fmt.Fprintln(w, s.name, "\t")

	for i := start; i < end; i++ {
		if len(s.index.index[i].value) > 1 {
			for j := range s.index.index[i].value {
				fmt.Fprint(w, s.index.index[i].value[j], "\t")
			}
		} else {
			fmt.Fprint(w, s.index.index[i].value[0], "\t")
		}

		fmt.Fprint(w, "|", "\t")
		fmt.Fprint(w, s.data[i], "\t")
		fmt.Fprintln(w)
	}
	w.Flush()
}

/* Indexing methods */

// Head prints the first howMany items in a Series object.
func (s *Series) Head(howMany int) {
	s.PrintRange(0, howMany)
}

// Tail prints the last howMany items in a Series object.
func (s *Series) Tail(howMany int) {
	s.PrintRange(len(s.data)-howMany, len(s.data))
}

// At returns an element at a given index.
// For multiindex, you need to pass in the whole index tuple.
func (s *Series) At(ind ...interface{}) (interface{}, error) {
	for i, index := range s.index.index {
		matchExists := false
		for j := 0; j < len(index.value); j++ {
			if index.value[j] == ind[j] {
				matchExists = true
				break
			}
		}
		if matchExists {
			return s.data[i], nil
		}
	}
	return nil, fmt.Errorf("the given index does not match any of the index in the series: %v", ind)
}

// IAt returns an element at a given integer index.
func (s *Series) IAt(ind int) (interface{}, error) {
	if ind >= len(s.data) {
		return nil, fmt.Errorf("index out of bounds: %v", ind)
	}
	if ind < 0 {
		return nil, fmt.Errorf("index can't be less than 0")
	}
	return s.data[ind], nil
}

// Loc accepts index tuples and returns a Series object containing data at the given rows.
// Each idx item should contain the index of the data you would like to query.
// For multiindex Series, you can either pass in the whole index tuple, or the first index.
func (s *Series) Loc(idx ...[]interface{}) (Series, error) {
	// This makes sure that each index passed are the same length.
	indexLength := len(idx[0])
	for i, eachIndex := range idx {
		if indexLength != len(eachIndex) {
			return Series{}, fmt.Errorf("index length does not match: %v, %v", idx[i-1], eachIndex)
		}
	}

	allFiltered := make([]interface{}, 0)
	allFilteredIndex := IndexData{}

	for _, inputIndex := range idx {
		filtered := make([]interface{}, 0)
		filteredIndex := IndexData{}

		for j, seriesIndex := range s.index.index {
			isSame := true
			for k := 0; k < indexLength; k++ {
				if inputIndex[k] != seriesIndex.value[k] {
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
			return Series{}, fmt.Errorf("no data found for index %v", inputIndex)
		}
		allFiltered = append(allFiltered, filtered...)
		allFilteredIndex.index = append(allFilteredIndex.index, filteredIndex.index...)
	}
	allFilteredIndex.names = s.index.names

	result, err := NewSeries(allFiltered, s.name, &allFilteredIndex)
	if err != nil {
		return Series{}, err
	}
	return result, nil
}

// LocItems acts the exact same as Loc, but returns data as []interface{} instead of Series.
func (s *Series) LocItems(idx ...[]interface{}) ([]interface{}, error) {
	indexLength := len(idx[0])
	for i, eachIndex := range idx {
		if indexLength != len(eachIndex) {
			return nil, fmt.Errorf("index length does not match: %v, %v", idx[i-1], eachIndex)
		}
	}

	allFiltered := make([]interface{}, 0)

	for _, inputIndex := range idx {
		filtered := make([]interface{}, 0)

		for j, seriesIndex := range s.index.index {
			isSame := true
			for k := 0; k < indexLength; k++ {
				if inputIndex[k] != seriesIndex.value[k] {
					isSame = false
					break
				}
			}
			if isSame {
				filtered = append(filtered, s.data[j])
			}
		}

		if len(filtered) == 0 {
			return nil, fmt.Errorf("no data found for index %v", inputIndex)
		}
		allFiltered = append(allFiltered, filtered...)
	}

	return allFiltered, nil
}

// ILoc returns an array of elements at a given integer index range.
func (s *Series) ILoc(min, max int) ([]interface{}, error) {
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

/* Summary statistics methods. These are Series-specific, unlike the ones in stats.go. */

// Count counts the number of non-NaN elements in a column.
func (s *Series) Count() StatsResult {
	count := 0
	for _, v := range s.data {
		if v != nil || v != math.NaN() {
			count++
		}
	}

	return StatsResult{"Count", float64(count), nil}
}

// Mean returns the mean of the elements in a column.
func (s *Series) Mean() StatsResult {
	mean := 0.0

	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Mean", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Mean", math.NaN(), fmt.Errorf("no elements in this column")}
	}

	for _, v := range data {
		mean += v
	}

	mean /= float64(len(data))
	roundedMean := math.Round(mean*1000) / 1000

	return StatsResult{"Mean", roundedMean, nil}
}

// Median returns the median of the elements in a column.
func (s *Series) Median() StatsResult {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Median", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Median", math.NaN(), fmt.Errorf("no elements in this column")}
	}
	if total%2 == 0 {
		lower := data[total/2-1]
		upper := data[total/2]

		median := (lower + upper) / 2
		roundedMedian := math.Round(median*1000) / 1000

		return StatsResult{"Median", roundedMedian, nil}
	} else {
		median := data[(total+1)/2-1]
		roundedMedian := math.Round(median*1000) / 1000

		return StatsResult{"Median", roundedMedian, nil}
	}
}

// Std returns the sample standard deviation of the elements in a column.
func (s *Series) Std() StatsResult {
	std := 0.0
	meanResult := s.Mean() // this also checks that all data can be converted to float64.
	if meanResult.Err != nil {
		return StatsResult{"Std", math.NaN(), meanResult.Err}
	}

	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Std", math.NaN(), err}
	}
	sort.Float64s(data)

	numerator := 0.0
	for _, v := range data {
		temp := math.Pow(v-meanResult.Result, 2)
		numerator += temp
	}
	std = math.Sqrt(numerator / float64(len(data)-1))
	roundedStd := math.Round(std*1000) / 1000

	return StatsResult{"Std", roundedStd, nil}
}

// Min returns the smallest element in a column.
func (s *Series) Min() StatsResult {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Min", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Min", math.NaN(), fmt.Errorf("no elements in this column")}
	}

	return StatsResult{"Min", data[0], nil}
}

// Max returns the largest element is a column.
func (s *Series) Max() StatsResult {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Max", math.NaN(), err}
	}
	sort.Float64s(data)

	total := len(data)
	if total == 0 {
		return StatsResult{"Max", math.NaN(), fmt.Errorf("no elements in this column")}
	}

	return StatsResult{"Max", data[total-1], nil}
}

// Q1 returns the lower quartile (25%) of the elements in a column.
// This does not include the median during calculation.
func (s *Series) Q1() StatsResult {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Q1", math.NaN(), err}
	}
	sort.Float64s(data)

	if len(data)%2 == 0 {
		lower := data[:len(data)/2]
		q1, err := median(lower)
		if err != nil {
			return StatsResult{"Q1", math.NaN(), err}
		}
		return StatsResult{"Q1", q1, nil}
	} else {
		lower := data[:(len(data)-1)/2]
		q1, err := median(lower)
		if err != nil {
			return StatsResult{"Q1", math.NaN(), err}
		}
		return StatsResult{"Q1", q1, nil}
	}
}

// Q2 returns the middle quartile (50%) of the elements in a column.
// This accomplishes the same thing as Median.
func (s *Series) Q2() StatsResult {
	q2Result := s.Median()
	if q2Result.Err != nil {
		return StatsResult{"Q2", math.NaN(), q2Result.Err}
	}

	return StatsResult{"Q2", q2Result.Result, nil}
}

// Q3 returns the upper quartile (75%) of the elements in a column.
// This does not include the median during calculation.
func (s *Series) Q3() StatsResult {
	data, err := interface2F64Slice(s.data)
	if err != nil {
		return StatsResult{"Q3", math.NaN(), err}
	}
	sort.Float64s(data)

	if len(data)%2 == 0 {
		upper := data[len(data)/2:]
		q3, err := median(upper)
		if err != nil {
			return StatsResult{"Q3", math.NaN(), err}
		}
		return StatsResult{"Q3", q3, nil}
	} else {
		upper := data[(len(data)+1)/2:]
		q3, err := median(upper)
		if err != nil {
			return StatsResult{"Q3", math.NaN(), err}
		}
		return StatsResult{"Q3", q3, nil}
	}
}

// Describe runs through the most commonly used statistics functions
// and prints the output.
func (s *Series) Describe() ([]StatsResult, error) {
	count := s.Count()
	fmt.Println("Count:", count.Result)

	mean := s.Mean()
	if mean.Err != nil {
		return nil, mean.Err
	}
	fmt.Println("Mean:", mean.Result)

	median := s.Median()
	if median.Err != nil {
		return nil, median.Err
	}
	fmt.Println("Median:", median.Result)

	std := s.Std()
	if std.Err != nil {
		return nil, std.Err
	}
	fmt.Println("Std:", std.Result)

	min := s.Min()
	if min.Err != nil {
		return nil, min.Err
	}
	fmt.Println("Min:", min.Result)

	max := s.Max()
	if max.Err != nil {
		return nil, max.Err
	}
	fmt.Println("Max:", max.Result)

	q1 := s.Q1()
	if q1.Err != nil {
		return nil, q1.Err
	}
	fmt.Println("Q1:", q1.Result)

	q2 := s.Q2()
	if q2.Err != nil {
		return nil, q2.Err
	}
	fmt.Println("Q2:", q2.Result)

	q3 := s.Q3()
	if q3.Err != nil {
		return nil, q3.Err
	}
	fmt.Println("Q3:", q3.Result)

	result := []StatsResult{
		count, mean, median, std, min, max, q1, q2, q3,
	}

	return result, nil
}

// ValueCounts returns a Series containing the number of unique values in a given Series.
func (s *Series) ValueCounts() (Series, error) {
	valueCountMap := make(map[interface{}]int, 0)
	for _, data := range s.data {
		// if key doesn't exist, create a new key and set initial value as 1.
		// if key exists, look up the data associated with the key and add 1.
		valueCountMap[data] += 1
	}

	newSeriesValue := make([]interface{}, 0)
	newSeriesIndex := IndexData{}
	newSeriesIndex.names = []string{"Data"}

	i := 0
	for k := range valueCountMap {
		newSeriesIndex.index = append(newSeriesIndex.index, Index{i, []interface{}{k}})
		i++
	}
	sort.Sort(newSeriesIndex)

	for _, v := range newSeriesIndex.index {
		newSeriesValue = append(newSeriesValue, valueCountMap[v.value[0]])
	}

	newS, err := NewSeries(newSeriesValue, fmt.Sprintf("Unique Value Count of %v", s.name), &newSeriesIndex)
	if err != nil {
		return Series{}, err
	}
	return newS, nil
}

/* Series manipulation methods */

// RenameCol renames the series.
func (s *Series) RenameCol(newName string) {
	s.name = newName
}

// RenameIndex renames the index of the series.
// Input should be a map, where key is the index name to change and value is a new name.
func (s *Series) RenameIndex(newNames map[string]string) error {
	exists := false
	for oldName, newName := range newNames {
		exists = false
		for i, indexName := range s.index.names {
			if indexName == oldName {
				s.index.names[i] = newName
				exists = true
			}
		}
		if !exists {
			return fmt.Errorf("index does not exist: %v", oldName)
		}
	}

	return nil
}

/* Sorting methods */

// SortByIndex sorts the elements in a series by the index.
func (s *Series) SortByIndex(ascending bool) error {
	indDatMap := make(map[string]interface{})
	for i, data := range s.data {
		key, err := s.index.index[i].hashKey()
		if err != nil {
			return err
		}
		indDatMap[*key] = data
	}

	if ascending {
		sort.Sort(s.index)
	} else {
		sort.Sort(sort.Reverse(s.index))
	}

	for i, index := range s.index.index {
		key, err := index.hashKey()
		if err != nil {
			return err
		}
		s.data[i] = indDatMap[*key]
	}

	return nil
}

// SortByGivenIndex sorts the Series by a given index.
func (s *Series) SortByGivenIndex(index IndexData) error {
	rowMap := make(map[string]interface{}, 0)
	for i, data := range s.data {
		key, err := s.index.index[i].hashKey()
		if err != nil {
			return err
		}
		rowMap[*key] = data
	}

	s.index = index

	for i, index := range s.index.index {
		key, err := index.hashKey()
		if err != nil {
			return err
		}
		s.data[i] = rowMap[*key]
	}

	return nil
}

// SortByValues sorts the Series by its values.
func (s *Series) SortByValues(ascending bool) error {
	rowMap := make(map[interface{}]Index, 0)
	keyStore := make(sort.StringSlice, 0)
	for i := range s.data {
		var key string
		if fmt.Sprint(s.data[i]) == "NaN" {
			key = fmt.Sprint("~", s.data[i], s.index.index[i].id)
		} else {
			key = fmt.Sprint(s.data[i], s.index.index[i].id)
		}
		keyStore = append(keyStore, key)
		rowMap[key] = s.index.index[i]
	}

	if ascending {
		sort.Sort(s)
		sort.Sort(keyStore)
	} else {
		sort.Sort(sort.Reverse(s))
		sort.Sort(sort.Reverse(keyStore))
	}

	for i := range s.data {
		s.index.index[i] = rowMap[keyStore[i]]
	}

	return nil
}

// IndexHasDuplicateValues checks if the Series have duplicate index values.
func (s *Series) IndexHasDuplicateValues() (bool, error) {
	indexDataMap := make(map[string]interface{}, 0)

	for _, index := range s.index.index {
		key, err := index.hashKeyValueOnly()
		if err != nil {
			return false, err
		}
		val, exists := indexDataMap[*key]
		if !exists {
			indexDataMap[*key] = val
		} else {
			return true, nil
		}
	}
	return false, nil
}
