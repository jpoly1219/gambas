package gambas

import "fmt"

type CustomIndex struct {
	Value []interface{}
}

type Series struct {
	Data        map[interface{}]interface{}
	CustomIndex CustomIndex
	Name        string
}

func (s Series) PrintSeries() string {
	message := fmt.Sprintln("Data:", s.Data, "\nCustomIndex:", s.CustomIndex, "\nName:", s.Name)
	fmt.Println(message)
	return message
}

type DataFrame struct {
	SeriesArray      []Series
	CustomIndexArray []CustomIndex
	Name             string
}
