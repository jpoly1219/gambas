package gambas

import "fmt"

// CreateRangeIndex takes the length of an Index and creates a RangeIndex.
// RangeIndex is an index that spans from 0 to the length of the index.
func CreateRangeIndex(length int) Index {
	zeroToLength := make([]interface{}, length)

	for i := 0; i < length; i++ {
		zeroToLength[i] = i
	}

	return Index{zeroToLength}
}

// NewSeries created a new Series object from given parameters.
// Generally, NewSeriesFromFile will be used more often.
func NewSeries(data []interface{}, name string) (*Series, error) {
	var s Series
	ok, err := checkTypeIntegrity(data)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}
	if !ok {
		fmt.Println("2")
		return nil, fmt.Errorf("types do not match")
	}
	s.data = data
	s.index = CreateRangeIndex(len(data))
	s.name = name
	return &s, nil
}

// NewDataFrame created a new DataFrame object from given parameters.
// Generally, NewDataFrameFromFile will be used more often.
func NewDataFrame(data [][]interface{}, columns []interface{}, indexCols []interface{}) (*DataFrame, error) {
	if len(data) != len(columns) {
		return nil, fmt.Errorf("length of data (%d) and columns (%d) does not match", len(data), len(columns))
	}

	var df DataFrame
	df.series = make([]Series, len(data))
	df.index = make([]Index, 0)
	df.columns.data = columns
	df.indexCols = indexCols

	// create df.index
	// find location of index column
	for i, col := range columns {
		for _, indexCol := range indexCols {
			if col == indexCol {
				df.index = append(df.index, Index{data[i]})
			}
		}
	}

	for i, v := range data {
		series, err := NewSeries(v, fmt.Sprint(columns[i]))
		if err != nil {
			return nil, err
		}
		df.series[i] = *series
	}

	return &df, nil
}

func checkTypeIntegrity(data []interface{}) (bool, error) {
	isFloat64 := false
	isString := false
	for _, v := range data {
		switch t := v.(type) {
		case float64:
			isFloat64 = true
		case string:
			isString = true
		default:
			_, err := i2f(v)
			if err != nil {
				return false, fmt.Errorf("invalid type: %T", t)
			} else {
				isFloat64 = true
			}
		}

		if isFloat64 && isString {
			return false, nil
		} else if !isFloat64 && !isString {
			return false, nil
		}
	}

	return true, nil
}

func i2f(data interface{}) (float64, error) {
	var x float64
	switch v := data.(type) {
	case int:
		x = float64(v)
	case int8:
		x = float64(v)
	case int16:
		x = float64(v)
	case int32:
		x = float64(v)
	case int64:
		x = float64(v)
	case uint:
		x = float64(v)
	case uint8:
		x = float64(v)
	case uint16:
		x = float64(v)
	case uint32:
		x = float64(v)
	case uint64:
		x = float64(v)
	case float32:
		x = float64(v)
	case float64:
		x = v
	default:
		return 0.0, fmt.Errorf("%v is not a number", data)
	}

	return x, nil
}
