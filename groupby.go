package gambas

type GroupBy struct {
	DataFrame *DataFrame
	colIndMap map[string][]interface{}
	colTuples [][]interface{}
}

// Agg aggregates data in the GroupBy object using the given aggFunc.
func (gb *GroupBy) Agg(targetCol []string, aggFunc StatsFunc) (*DataFrame, error) {
	valueSlice := make([]interface{}, 0)

	for i, idx := gb.index {
		for j, col := range gb.columns {

		}
	}



	for k, v := range gb.dataMap {
		result := aggFunc(v)
		statsResultMap[k] = result
	}
	
}
