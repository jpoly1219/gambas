package gambas

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Index []interface{}

type IndexData struct {
	index []Index
	names []string
}

func (id IndexData) Len() int {
	return len(id.index)
}

func (id IndexData) Less(i, j int) bool {
	var iStrData, jStrData string

	for a := range id.index[0] {
		switch v := id.index[i][a].(type) {
		case string:
			iStrData = v
			jStrData = id.index[j][a].(string)
		case float64:
			iStrData = strconv.FormatFloat(v, 'f', -1, 64)
			jStrData = strconv.FormatFloat(id.index[j][a].(float64), 'f', -1, 64)
		}

		if iStrData == jStrData {
			continue
		} else {
			break
		}
	}
	return iStrData < jStrData
}

func (id IndexData) Sort(i, j int) {
	id.index[i], id.index[j] = id.index[j], id.index[i]
}

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
	for i, eachIndex := range in {
		if indexLength != len(eachIndex) {
			return nil, fmt.Errorf("index length does not match: %v, %v", in[i-1], eachIndex)
		}
	}

	filtered := make([]interface{}, 0)
	for _, inputIndex := range in {
		for j, seriesIndex := range s.index.index {
			if indexLength == 1 {
				if inputIndex[0] == seriesIndex[0] {
					filtered = append(filtered, s.data[j])
				}
			}

			if indexLength > 1 {
				isSame := true
				for k := 0; k < indexLength; k++ {
					if inputIndex[k] != seriesIndex[k] {
						isSame = false
						break
					}
				}
				if isSame {
					filtered = append(filtered, s.data[j])
				}
			}
		}
	}

	result, err := NewSeries(filtered, s.name)
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

	return mean, nil
}

// Median() returns the median of the elements in a column.
func (s Series) Median() (float64, error) {
	median := 0.0

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

		median = (lower + upper) / 2
	} else {
		median := data[(total+1)/2-1]
		return median, nil
	}

	return median, nil
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

	return std, nil
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

// SortByIndex() sorts the elements in a series by the index.
// Multiindex support is coming, but this may require an overhaul.
func (s *Series) SortByIndex(ascending bool) {
	indDatMap := make(map[interface{}]interface{})
	for i, data := range s.data {
		indDatMap[s.index[i]] = data
	}

	if ascending {
		sort.Sort(s.index)
	} else {
		sort.Sort(sort.Reverse(s.index))
	}

	for i, index := range s.index {
		s.data[i] = indDatMap[index]
	}
}

type DataFrame struct {
	series    []Series
	columns   Index
	indexCols []interface{}
	index     []Index
}

// LocRows returns a set of rows as a new DataFrame object, given a list of labels.
func (df DataFrame) LocRows(rows []interface{}) (*DataFrame, error) {
	locations := make([]int, 0)

	for _, row := range rows {
		for _, index := range df.index {
			for i, value := range index {
				if row == value {
					locations = append(locations, i)
				}
			}
		}
	}

	filteredCols := make([][]interface{}, 0)
	for _, series := range df.series {
		filteredCol := make([]interface{}, 0)
		for _, location := range locations {
			filteredCol = append(filteredCol, series.data[location])
		}
		filteredCols = append(filteredCols, filteredCol)
	}

	dataframe, err := NewDataFrame(filteredCols, df.columns, df.indexCols)
	if err != nil {
		return nil, err
	}

	return dataframe, nil
}

// LocRows returns a set of columns as a new DataFrame object, given a list of labels.
func (df DataFrame) LocCols(cols []interface{}) (*DataFrame, error) {
	filtered2D := make([][]interface{}, 0)
	for _, column := range cols {
		for _, series := range df.series {
			if series.name == column {
				filtered2D = append(filtered2D, series.data)
			}
		}
	}

	dataframe, err := NewDataFrame(filtered2D, cols, []interface{}{cols[0]})
	if err != nil {
		return nil, err
	}

	return dataframe, nil
}

// Loc indexes the DataFrame object given a slice of row and column labels.
func (df DataFrame) Loc(rows, cols []interface{}) (*DataFrame, error) {
	df1, err := df.LocCols(cols)
	if err != nil {
		return nil, err
	}

	df2, err := df1.LocRows(rows)
	if err != nil {
		return nil, err
	}

	return df2, nil
}

// Basic arithmetic operations for columns.

// ColAdd() adds the given value to each element in the specified column.
func (df *DataFrame) ColAdd(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v += value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot add, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColSub() subtracts the given value from each element in the specified column.
func (df *DataFrame) ColSub(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v -= value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot subtract, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColMul() multiplies each element in the specified column by the given value.
func (df *DataFrame) ColMul(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v *= value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot multiply, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColDiv() divides each element in the specified column by the given value.
func (df *DataFrame) ColDiv(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v /= value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot divide, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColMod() applies modulus calculations on each element in the specified column, returning the remainder.
func (df *DataFrame) ColMod(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					series.data[i] = math.Mod(v, value)
				default:
					return nil, fmt.Errorf("cannot use modulus, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// Basic boolean operators for columns.

// ColGt() checks if each element in the specified column is greater than the given value.
func (df *DataFrame) ColGt(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					isGt := (v > value)
					series.data[i] = isGt
				default:
					return nil, fmt.Errorf("cannot compare, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColLt() checks if each element in the specified column is less than the given value.
func (df *DataFrame) ColLt(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					isLt := (v < value)
					series.data[i] = isLt
				default:
					return nil, fmt.Errorf("cannot compare, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColEq() checks if each element in the specified column is equal to the given value.
func (df *DataFrame) ColEq(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					isEq := (v == value)
					series.data[i] = isEq
				default:
					return nil, fmt.Errorf("cannot compare, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// NewCol() creates a new column with the given data and column name.
// To create a blank column, pass in a slice with zero values.
// Use this in conjunction with operators, like this: df.NewCol().ColAdd()
func (df *DataFrame) NewCol(colname string, data []interface{}) (*DataFrame, error) {
	newSeries, err := NewSeries(data, colname)
	if err != nil {
		return nil, err
	}

	df.series = append(df.series, *newSeries)
	df.columns = append(df.columns, colname)

	return df, nil
}
