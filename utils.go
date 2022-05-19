package gambas

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func checkTypeIntegrity(data []interface{}) (bool, error) {
	isFloat64 := 0
	isString := 0
	isNil := 0
	for _, v := range data {
		switch t := v.(type) {
		case float64:
			if math.IsNaN(t) {
				continue
			}
			isFloat64 = 1
		case string:
			isString = 1
		case nil:
			isNil = 1
		default:
			_, err := i2f(v)
			if err != nil {
				return false, fmt.Errorf("invalid type: %T", t)
			} else {
				isFloat64 = 1
			}
		}

		if isFloat64+isString+isNil > 1 {
			return false, nil
		} else if isFloat64+isString+isNil == 0 {
			panic("type not detected")
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

// checkType checks to see if the data can be represented as a float64.
// Because CSV is read as an array of strings, there has to be a way to check the type.
func checkCSVDataType(data string) interface{} {
	if data == "NaN" || data == "" {
		return math.NaN()
	}
	v, ok := strconv.ParseFloat(data, 64)
	switch ok {
	case nil:
		return v
	default:
		return data
	}
}

// interface2F64Data() converts a slice of interface{} into F64Data.
func interface2F64Slice(data []interface{}) ([]float64, error) {
	fd := make([]float64, 0)
	for _, v := range data {
		switch converted := v.(type) {
		case float64:
			if math.IsNaN(converted) {
				continue
			}
			fd = append(fd, converted)
		default:
			return nil, fmt.Errorf("data is not a float64: %v", v)
		}
	}

	return fd, nil
}

// interface2StringData() converts a slice of interface{} into StringData.
func interface2StringSlice(data []interface{}) ([]string, error) {
	sd := make([]string, 0)
	for _, v := range data {
		switch converted := v.(type) {
		case string:
			sd = append(sd, converted)
		default:
			return nil, fmt.Errorf("data is not a string: %v", v)
		}
	}

	return sd, nil
}

// equal checks whether two slices are equal.
func equal(slice1, slice2 []interface{}) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

// containsString checks whether a string exists in a slice of strings.
func containsString(strSlice []string, str string) bool {
	for _, data := range strSlice {
		if data == str {
			return true
		}
	}
	return false
}

// containsIndex checks whether an Index object exists in a slice of Index objects.
func containsIndex(indexSlice []Index, index Index) bool {
	for _, data := range indexSlice {
		if equal(data.value, index.value) {
			return true
		}
	}
	return false
}

// Summary statistics functions (internal use only)

// median() returns the median of the elements in an array.
func median(data []float64) (float64, error) {
	median := 0.0
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
