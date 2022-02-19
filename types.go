package gambas

import (
	"fmt"
	"reflect"
)

type Index struct {
	Value interface{}
}

type Series struct {
	Data       map[Index]interface{}
	IndexArray []Index
	Name       string
}

func (s Series) PrintSeries() string {
	message := fmt.Sprintln("Data:", s.Data, "\nIndexArray:", s.IndexArray, "\nName:", s.Name)
	fmt.Println(message)
	return message
}

func (s Series) CalcMean() (float64, error) {
	sum := 0.0
	for _, v := range s.Data {
		if reflect.ValueOf(v).Kind() != reflect.Float64 || reflect.ValueOf(v).Kind() != reflect.Int {
			err := fmt.Errorf("type is not int or float64: %T", v)
			fmt.Println(err)
			return 0.0, err
		}
		sum += v.(float64)
	}
	mean := sum / float64(len(s.Data))

	return mean, nil
}

func (s Series) At(index Index) interface{} {
	result := s.Data[index]

	return result
}

func (s Series) AtM(indexArray []Index) []interface{} {
	resultArray := make([]interface{}, len(indexArray))

	for i, v := range indexArray {
		result := s.Data[v]
		resultArray[i] = result
	}

	return resultArray
}

func (s Series) AtR(min, max int) []interface{} {
	resultArray := make([]interface{}, 0)

	for i := min; i <= max; i++ {
		key := s.IndexArray[i]
		result := s.Data[key]
		resultArray = append(resultArray, result)
	}

	return resultArray
}

type DataFrame struct {
	SeriesArray []Series
	IndexArray  []Index
	Name        string
}
