package gambas

import (
	"fmt"
	"reflect"
)

// checkType checks to see if the data can be represented as a float64.
// Because CSV is read as an array of strings, there has to be a way to check the type.
func checkType(data interface{}) interface{} {
	switch data {
	case data.(float64):
		return data.(float64)
	default:
		return data.(string)
	}
}

type Index struct {
	data []interface{}
}

type Series struct {
	data  []interface{}
	index Index
	name  string
}

func (s Series) PrintSeries() string {
	message := fmt.Sprintln("data:", s.data, "\nindexArray:", s.index, "\nname:", s.name)
	fmt.Println(message)
	return message
}

func (s Series) CalcMean() (float64, error) {
	sum := 0.0
	for _, v := range s.data {
		if reflect.ValueOf(v).Kind() != reflect.Float64 || reflect.ValueOf(v).Kind() != reflect.Int {
			err := fmt.Errorf("type is not int or float64: %T", v)
			fmt.Println(err)
			return 0.0, err
		}
		sum += v.(float64)
	}
	mean := sum / float64(len(s.data))

	return mean, nil
}

// At() returns the element at a given index.
func (s Series) At(index interface{}) (interface{}, error) {
	for i, v := range s.index.data {
		if v == index {
			result := s.data[i]
			return result, nil
			break
		}
	}

	return nil, fmt.Errorf("index %s is not found", index)
}

// AtM() returns an array of elements at given indexes.
func (s Series) AtM(indexArray []interface{}) ([]interface{}, error) {
	resultArray := make([]interface{}, len(indexArray))

	for _, v := range indexArray {
		result, err := s.At(v)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, result)
	}

	return resultArray, nil
}

// AtR() returns an array of elements at a given index range.
func (s Series) AtR(min, max int) ([]interface{}, error) {
	resultArray := make([]interface{}, 0)

	for i := min; i < max; i++ {
		key := s.index.data[i]
		result, err := s.At(key)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, result)
	}

	return resultArray, nil
}

type DataFrame struct {
	series  []Series
	index   Index
	columns Index
}

func Loc(row, col string) {

}
