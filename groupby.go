package gambas

type GroupBy struct {
	dataFrame       *DataFrame
	colIndMap       map[string][]interface{}
	colTuples       [][]interface{}
	colTuplesLabels []string
}

// Agg aggregates data in the GroupBy object using the given aggFunc.
func (gb *GroupBy) Agg(targetCol []string, aggFunc StatsFunc) (DataFrame, error) {
	filtered, err := gb.dataFrame.LocCols(targetCol...)
	if err != nil {
		return DataFrame{}, err
	}

	newDfData := make([][]interface{}, len(gb.colTuples[0]))
	for _, colTuple := range gb.colTuples {
		for j, col := range colTuple {
			newDfData[j] = append(newDfData[j], col)
		}
	}

	results := make([]interface{}, 0)
	for _, ser := range filtered.series {
		for i, colTuple := range gb.colTuples {
			colTupleIndex := Index{i, colTuple}
			key, err := colTupleIndex.hashKeyValueOnly()
			if err != nil {
				return DataFrame{}, err
			}

			indexForData := gb.colIndMap[*key]
			data := make([]interface{}, 0)
			for _, id := range indexForData {
				d, err := ser.IAt(id.(int))
				if err != nil {
					return DataFrame{}, err
				}
				data = append(data, d)
			}
			result := aggFunc(data)
			results = append(results, result.Result)
		}
	}

	newDfData = append(newDfData, results)
	newDfColumns := make([]string, 0)
	newDfColumns = append(newDfColumns, gb.colTuplesLabels...)
	newDfColumns = append(newDfColumns, filtered.columns...)

	newDf, err := NewDataFrame(newDfData, newDfColumns, gb.colTuplesLabels)
	if err != nil {
		return DataFrame{}, err
	}

	newDf.SortByIndex(true)
	return newDf, nil
}
