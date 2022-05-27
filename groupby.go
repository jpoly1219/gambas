package gambas

type GroupBy struct {
	index   []interface{}
	columns []string
	dataMap map[string][]interface{}
}

// Agg aggregates data in the GroupBy object using the given aggFunc.
func (gb *GroupBy) Agg(columns []string, aggFunc StatsFunc) (*DataFrame, error) {
	statsResultMap := make(map[string]StatsResult, 0)
	for k, v := range gb.dataMap {
		result := aggFunc(v)
		statsResultMap[k] = result
	}
	
}
