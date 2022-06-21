package gambas

import (
	"crypto/sha512"
	"fmt"
	"strconv"
)

// Index stores the index values of a series and dataframe.
// The 0th element must be the ID of the index.
// For example, if your data includes a column of names that you have set to be the index,
// the index may look like this: Index{0, "Alice"}, Index{1, "Bob"}, Index{2, "Charlie"}.
// Index{} with more than one value (not including the ID) is considered a multi-index.
type Index struct {
	id    int
	value []interface{}
}

func (i Index) hashKey() (*string, error) {
	if len(i.value) == 0 {
		return nil, fmt.Errorf("no index")
	}

	byteSlice := []byte{}
	byteSlice = append(byteSlice, []byte(strconv.Itoa(i.id))...)

	for _, val := range i.value {
		byteSlice = append(byteSlice, []byte(fmt.Sprint(val))...)
	}

	h := sha512.Sum512(byteSlice)

	resultHex := fmt.Sprintf("%x", h)

	return &resultHex, nil
}

func (i Index) hashKeyValueOnly() (*string, error) {
	if len(i.value) == 0 {
		return nil, fmt.Errorf("no index")
	}

	byteSlice := []byte{}

	for _, val := range i.value {
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
	var iStr, jStr string
	for a := range id.index[0].value {
		if fmt.Sprint(id.index[i].value[a]) == "NaN" {
			return false
		}

		if fmt.Sprint(id.index[j].value[a]) == "NaN" {
			return true
		}

		if id.index[i].value[a] != id.index[j].value[a] {
			iStr = fmt.Sprint(id.index[i].value[a])
			jStr = fmt.Sprint(id.index[j].value[a])
			break
		}
	}
	return iStr < jStr
}

func (id IndexData) Swap(i, j int) {
	id.index[i], id.index[j] = id.index[j], id.index[i]
}
