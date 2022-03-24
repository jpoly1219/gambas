package gambas

import (
	"fmt"
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