package gambas

import (
	"crypto/sha512"
	"fmt"
	"math"
	"strconv"
)

// Index stores the index values of a series and dataframe.
// The 0th element must be the ID of the index.
// For example, if your data includes a column of names that you have set to be the index,
// the index may look like this: Index{0, "Alice"}, Index{1, "Bob"}, Index{2, "Charlie"}.
// Index{} with more than one value (not including the ID) is considered a multi-index.
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
		case int:
			// TODO: so all numbers in the DataFrame should be float64, we need a way to check why there is an int in the first place
			iStrData = strconv.Itoa(v)
		case float64:
			if math.IsNaN(v) {
				iStrData = "~NaN"
				break
			}
			iStrData = strconv.FormatFloat(v, 'f', -1, 64)
		}

		switch v := id.index[j][a].(type) {
		case string:
			jStrData = v
		case int:
			// TODO: so all numbers in the DataFrame should be float64, we need a way to check why there is an int in the first place
			jStrData = strconv.Itoa(v)
		case float64:
			if math.IsNaN(v) {
				jStrData = "~NaN"
				break
			}
			jStrData = strconv.FormatFloat(v, 'f', -1, 64)
		}

		if iStrData == jStrData {
			continue
		} else {
			break
		}
	}

	return iStrData < jStrData
}

func (id IndexData) Swap(i, j int) {
	id.index[i], id.index[j] = id.index[j], id.index[i]
}
