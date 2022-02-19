package gambas

import (
	"fmt"
	"reflect"
)

type Index struct {
	value interface{}
}

type Series struct {
	data       map[Index]interface{}
	indexArray []Index
	name       string
}

func (s Series) PrintSeries() string {
	message := fmt.Sprintln("data:", s.data, "\nindexArray:", s.indexArray, "\nname:", s.name)
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

func (s Series) At(index Index) interface{} {
	result := s.data[index]

	return result
}

func (s Series) AtM(indexArray []Index) []interface{} {
	resultArray := make([]interface{}, len(indexArray))

	for i, v := range indexArray {
		result := s.data[v]
		resultArray[i] = result
	}

	return resultArray
}

func (s Series) AtR(min, max int) []interface{} {
	resultArray := make([]interface{}, 0)

	for i := min; i <= max; i++ {
		key := s.indexArray[i]
		result := s.data[key]
		resultArray = append(resultArray, result)
	}

	return resultArray
}

type DataFrame struct {
	SeriesArray []Series
	IndexArray  []Index
	Name        string
}
