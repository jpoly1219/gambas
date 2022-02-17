package gambas

import (
	"fmt"
)

type Index struct {
	Value interface{}
}

type Series struct {
	Data map[Index]interface{}
	Name string
}

func (s Series) PrintElements() {
	if s.Name != "" {
		fmt.Println(" ", s.Name)
	}
	for k, v := range s.Data {
		fmt.Println(k.Value, v)
	}
}

type DataFrame struct {
	SeriesArray []Series
	IndexArray  []Index
	Name        string
}
