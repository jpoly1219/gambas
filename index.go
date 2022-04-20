package gambas

import (
	"crypto/sha512"
	"fmt"
	"strconv"
)

type Index []interface{}

func (i Index) hashKey() (*string, error) {
	if len(i) == 0 {
		return nil, fmt.Errorf("no index")
	}

	byteSlice := []byte{}

	for _, val := range i {
		byteSlice = append(byteSlice, []byte(fmt.Sprint(val))...)
	}
	h := sha512.Sum512(byteSlice)

	resultHex := fmt.Sprintf("%x", h)

	return &resultHex, nil
}

type IndexData struct {
	index []Index
	names []string
}

func (id IndexData) Len() int {
	return len(id.index)
}

func (id IndexData) Less(i, j int) bool {
	var iStrData, jStrData string

	for a := range id.index[0] {
		switch v := id.index[i][a].(type) {
		case string:
			iStrData = v
			jStrData = id.index[j][a].(string)
		case int:
			// TODO: so all numbers in the DataFrame should be float64, we need a way to check why there is an int in the first place
			iStrData = strconv.Itoa(v)
			jStrData = strconv.Itoa(id.index[j][a].(int))
		case float64:
			iStrData = strconv.FormatFloat(v, 'f', -1, 64)
			jStrData = strconv.FormatFloat(id.index[j][a].(float64), 'f', -1, 64)
		}

		if iStrData == jStrData {
			continue
		} else {
			break
		}
	}

	fmt.Println(iStrData < jStrData)
	return iStrData < jStrData
}

func (id IndexData) Swap(i, j int) {
	id.index[i], id.index[j] = id.index[j], id.index[i]
}
