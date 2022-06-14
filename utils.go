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

// tryBool checks if a string can be converted into bool.
// tryBool only accepts "TRUE", "True", "true", and "FALSE", "False", "false".
func tryBool(data string) (bool, error) {
	ignored := []string{"1, t, T, 0, f, F"}
	if containsString(ignored, data) {
		return false, fmt.Errorf("ignored string")
	}

	b, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return b, nil
}

// tryInt checks if a string can be converted into int.
func tryInt(data string) (int, error) {
	s, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return s, nil
}

// tryFloat64 checks if a string can be converted into float64.
func tryFloat64(data string) (float64, error) {
	f, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// tryDataType accepts a string and tries to convert it to the correct data type.
// It will try to convert the data into a bool, then int, then float64, and finally string.
func tryDataType(data string) interface{} {
	b, err := tryBool(data)
	if err != nil {
		i, err := tryInt(data)
		if err != nil {
			f, err := tryFloat64(data)
			if err != nil {
				return data
			}
			return f
		}
		return i
	}
	return b
}

// checkType checks to see if the data can be represented as a float64.
// Because CSV is read as an array of strings, there has to be a way to check the type.
func checkCSVDataType(data string) interface{} {
	if data == "" {
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

// checkJsonDataType checks for JSON null data, and converts it into math.NaN().
func checkJsonDataType(data interface{}) interface{} {
	if data == nil {
		return math.NaN()
	}
	return data
}

// consolidateToString consolidates all data in an []interface{} to string.
// In order to stay compatible with Series.data,
// the data type of the slice is still an empty interface.
func consolidateToString(data []interface{}) []interface{} {
	result := make([]interface{}, len(data))
	for i, d := range data {
		result[i] = fmt.Sprint(d)
	}

	return result
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

// slicesAreEqual checks whether two slices are equal.
func slicesAreEqual(slice1, slice2 []interface{}) bool {
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
// For a lenient comaprison that checks for values only, use containsIndexWithoutId.
func containsIndex(indexSlice []Index, index Index) bool {
	for _, data := range indexSlice {
		if slicesAreEqual(data.value, index.value) && (data.id == index.id) {
			return true
		}
	}
	return false
}

// containsIndex checks whether an Index object exists in a slice of Index objects.
// For a strict comaprison that checks for id as well, use containsIndex.
func containsIndexWithoutId(indexSlice []Index, index Index) bool {
	for _, data := range indexSlice {
		if slicesAreEqual(data.value, index.value) {
			return true
		}
	}
	return false
}

// containsSlice checks whether a slice of interface{} exists in a slice of []interface{}.
func containsSlice(s1 [][]interface{}, s2 []interface{}) bool {
	for _, data := range s1 {
		if slicesAreEqual(data, s2) {
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
