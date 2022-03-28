package gambas

import (
	"fmt"
	"math"
	"strconv"
)

func checkTypeIntegrity(data []interface{}) (bool, error) {
	isFloat64 := 0
	isString := 0
	isNil := 0
	for _, v := range data {
		switch t := v.(type) {
		case float64:
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
	if data == "NaN" {
		fmt.Println("NaN detected")
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
func interface2F64Data(data []interface{}) (F64Data, error) {
	fd := make(F64Data, 0)
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
func interface2StringData(data []interface{}) (StringData, error) {
	sd := make(StringData, 0)
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
