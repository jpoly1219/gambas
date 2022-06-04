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

	newDfData := make([][]interface{}, len(gb.colTuples[0]))
	for _, colTuple := range gb.colTuples {
		for j, col := range colTuple {
			newDfData[j] = append(newDfData[j], col)
		}
	}

	results := make([]interface{}, 0)
	for _, ser := range filtered.series {
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
			results = append(results, result.Result)
		}
	}

	newDfData = append(newDfData, results)
	newDfColumns := make([]string, 0)
	newDfColumns = append(newDfColumns, gb.colTuplesLabels...)
	newDfColumns = append(newDfColumns, filtered.columns...)
	fmt.Println(newDfColumns)

	newDf, err := NewDataFrame(newDfData, newDfColumns, gb.colTuplesLabels)
	if err != nil {
		return nil, err
	}

	newDf.SortByIndex(true)
	return newDf, nil
}
