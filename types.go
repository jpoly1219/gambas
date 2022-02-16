package gambas

import (
	"fmt"
)

type Element struct {
	Value interface{}
}

type Index struct {
	Value interface{}
}

type Series struct {
	Elements []Element
	Indexes  []Index
	Name     string
}

func (s Series) PrintElements() {
	if s.Name != "" {
		fmt.Println(" ", s.Name)
	}
	for i, e := range s.Elements {
		fmt.Println(s.Indexes[i].Value, e.Value)
	}
}

// func main() {
// 	data := []Element{{"alice"}, {"bob"}, {"charlie"}}

// 	series1, err := NewSeries(data, nil, "")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	series1.PrintElements()
// }
