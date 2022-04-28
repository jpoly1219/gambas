package gambas

import (
	"fmt"
	"math"
)

type DataFrame struct {
	series  []Series
	index   IndexData
	columns []string
}

// Head prints the first n items in the dataframe.
func (df DataFrame) Head(howMany int) {
	fmt.Println(df.index.names, df.columns)
	for i := 0; i < howMany; i++ {
		fmt.Print(df.index.index[i], " ")
		for _, ser := range df.series {
			fmt.Print(ser.data[i], " ")
		}
		fmt.Println()
	}
}

// LocRows returns a set of rows as a new DataFrame object, given a list of labels.
func (df DataFrame) LocRows(rows []Index) (*DataFrame, error) {
	filteredData := make([][]interface{}, 0)
	filteredColname := make([]string, 0)
	filteredIndex := IndexData{}
	for _, series := range df.series {
		located, err := series.Loc(rows)
		if err != nil {
			return nil, err
		}
		filteredData = append(filteredData, located.data)
		filteredColname = append(filteredColname, located.name)
		filteredIndex = located.index
	}

	dataframe, err := NewDataFrame(filteredData, filteredColname, filteredIndex.names)
	if err != nil {
		return nil, err
	}

	// When NewDataFrame is called, the resulting dataframe may have empty index values.
	// This is because NewDataFrame searches for index values in filtered2D,
	// but if the columns in the dataframe does not match filteredIndex.names,
	// there would be no matching columns, thus returning empty indexes.
	copy(dataframe.index.index, rows)

	for _, ser := range dataframe.series {
		copy(ser.index.index, rows)
	}

	return dataframe, nil

	// locations := make([]int, 0)

	// for _, row := range rows {
	// 	for _, index := range df.index {
	// 		for i, value := range index {
	// 			if row == value {
	// 				locations = append(locations, i)
	// 			}
	// 		}
	// 	}
	// }

	// filteredCols := make([][]interface{}, 0)
	// for _, series := range df.series {
	// 	filteredCol := make([]interface{}, 0)
	// 	for _, location := range locations {
	// 		filteredCol = append(filteredCol, series.data[location])
	// 	}
	// 	filteredCols = append(filteredCols, filteredCol)
	// }

	// dataframe, err := NewDataFrame(filteredCols, df.columns, df.indexCols)
	// if err != nil {
	// 	return nil, err
	// }

	// return dataframe, nil
}

// LocRows returns a set of columns as a new DataFrame object, given a list of labels.
func (df DataFrame) LocCols(cols []string) (*DataFrame, error) {
	filtered2D := make([][]interface{}, 0)
	for _, column := range cols {
		for _, series := range df.series {
			if series.name == column {
				filtered2D = append(filtered2D, series.data)
			}
		}
	}

	dataframe, err := NewDataFrame(filtered2D, cols, df.index.names)
	if err != nil {
		return nil, err
	}

	// When NewDataFrame is called, the resulting dataframe may have empty index values.
	// This is because NewDataFrame searches for index values in filtered2D,
	// but if the index column name is different from the column the user is trying to LocCols,
	// there would be no matching columns.

	copy(dataframe.index.index, df.index.index)

	for _, ser := range dataframe.series {
		copy(ser.index.index, df.index.index)
	}

	return dataframe, nil
}

// Loc indexes the DataFrame object given a slice of row and column labels.
func (df DataFrame) Loc(rows []Index, cols []string) (*DataFrame, error) {
	df1, err := df.LocCols(cols)
	if err != nil {
		return nil, err
	}

	df2, err := df1.LocRows(rows)
	if err != nil {
		return nil, err
	}

	return df2, nil
}

// Basic arithmetic operations for columns.

// ColAdd() adds the given value to each element in the specified column.
func (df *DataFrame) ColAdd(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v += value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot add, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColSub() subtracts the given value from each element in the specified column.
func (df *DataFrame) ColSub(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v -= value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot subtract, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColMul() multiplies each element in the specified column by the given value.
func (df *DataFrame) ColMul(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v *= value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot multiply, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColDiv() divides each element in the specified column by the given value.
func (df *DataFrame) ColDiv(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					v /= value
					series.data[i] = v
				default:
					return nil, fmt.Errorf("cannot divide, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColMod() applies modulus calculations on each element in the specified column, returning the remainder.
func (df *DataFrame) ColMod(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					series.data[i] = math.Mod(v, value)
				default:
					return nil, fmt.Errorf("cannot use modulus, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// Basic boolean operators for columns.

// ColGt() checks if each element in the specified column is greater than the given value.
func (df *DataFrame) ColGt(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					isGt := (v > value)
					series.data[i] = isGt
				default:
					return nil, fmt.Errorf("cannot compare, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColLt() checks if each element in the specified column is less than the given value.
func (df *DataFrame) ColLt(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					isLt := (v < value)
					series.data[i] = isLt
				default:
					return nil, fmt.Errorf("cannot compare, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// ColEq() checks if each element in the specified column is equal to the given value.
func (df *DataFrame) ColEq(colname string, value float64) (*DataFrame, error) {
	for _, series := range df.series {
		if series.name == colname {
			for i, data := range series.data {
				switch v := data.(type) {
				case float64:
					isEq := (v == value)
					series.data[i] = isEq
				default:
					return nil, fmt.Errorf("cannot compare, column data type is not float64")
				}
			}
			return df, nil
		}
	}
	return nil, fmt.Errorf("colname does not match any of the existing column names")
}

// NewCol() creates a new column with the given data and column name.
// To create a blank column, pass in a slice with zero values.
// Use this in conjunction with operators, like this: df.NewCol().ColAdd()
func (df *DataFrame) NewCol(colname string, data []interface{}) (*DataFrame, error) {
	newSeries, err := NewSeries(data, colname, &df.index)
	if err != nil {
		return nil, err
	}

	df.series = append(df.series, *newSeries)
	df.columns = append(df.columns, colname)

	return df, nil
}

// RenameCol renames a column in a DataFrame.
func (df *DataFrame) RenameCol(oldColname, newColname string) error {
	for i, col := range df.columns {
		if col == oldColname {
			df.columns[i] = newColname
			break
		}

		if i+1 == len(df.columns) {
			return fmt.Errorf("column %v does not exist", oldColname)
		}
	}

	for i, name := range df.index.names {
		if name == oldColname {
			df.index.names[i] = newColname
		}
	}

	return nil
}

func (df *DataFrame) SortByIndex(ascending bool) error {
	fmt.Println("before:", df)
	if len(df.series) > 0 {
		for i := range df.series {
			df.series[i].SortByIndex(ascending)
		}
	}
	fmt.Println("after:", df)
	df.index = df.series[0].index
	return nil
}
