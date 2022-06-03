package gambas

import "fmt"

type GroupBy struct {
	dataFrame       *DataFrame
	colIndMap       map[string][]interface{}
	colTuples       [][]interface{}
	colTuplesLabels []string
}

// Agg aggregates data in the GroupBy object using the given aggFunc.
func (gb *GroupBy) Agg(targetCol []string, aggFunc StatsFunc) (*DataFrame, error) {
	filtered, err := gb.dataFrame.LocCols(targetCol...)
	if err != nil {
		return nil, err
	}

	newDfData := make([][]interface{}, len(gb.colTuples[0])+len(targetCol))
	for i, ser := range filtered.series {
		for j, colTuple := range gb.colTuples {
			colTupleIndex := Index{j, colTuple}
			key, err := colTupleIndex.hashKeyValueOnly()
			if err != nil {
				return nil, err
			}

			indexForData := gb.colIndMap[*key]
			ifd2d := make([][]interface{}, 0)
			for _, ifd := range indexForData {
				ifd2d = append(ifd2d, []interface{}{ifd})
			}
			data, err := ser.LocItems(ifd2d...)
			if err != nil {
				return nil, err
			}
			result := aggFunc(data)

			for k := range colTuple {
				newDfData[i] = append(newDfData[i], colTuple[i])
			}
		}
	}

	fmt.Println(colTupleToResult)

	newDf, err := NewDataFrame(colTupleToResult, gb.dataFrame.columns, gb.dataFrame.index.names)
	if err != nil {
		return nil, err
	}
	return newDf, nil
}
