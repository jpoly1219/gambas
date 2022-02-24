package gambas

import (
	"fmt"
	"reflect"
)

type Index struct {
	data []interface{}
}

type Series struct {
	data  map[interface{}]interface{}
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
func (s Series) At(index interface{}) interface{} {
	result := s.data[index]

	return result
}

// AtM() returns an array of elements at given indexes.
func (s Series) AtM(indexArray []interface{}) []interface{} {
	resultArray := make([]interface{}, len(indexArray))

	for i, v := range indexArray {
		result := s.data[v]
		resultArray[i] = result
	}

	return resultArray
}

// AtR() returns an array of elements at a given index range.
func (s Series) AtR(min, max int) []interface{} {
	resultArray := make([]interface{}, 0)

	for i := min; i <= max; i++ {
		key := s.index.data[i]
		result := s.data[key]
		resultArray = append(resultArray, result)
	}

	return resultArray
}

type DataFrame struct {
	series  map[interface{}]Series
	index   Index
	columns Index
}
