# Creating New Objects

`gambas` provides you with generator functions that can create a RangeIndex (an index that spans from 0 to a certain length), a Series, or a DataFrame object. Usually, you won't be creating objects from scratch using the generator functions. Instead, you will use I/O functions to generate objects out of pre-existing data.

## NewSeries

```go
func NewSeries(data []interface{}, name string, index *IndexData) (Series, error)
```

`NewSeries` takes in a slice of data, a name, and an IndexData object, returning a Series object and an error. If you leave `index` parameter as `nil`, then `NewSeries` will automatically generate a RangeIndex. The `index` field is meant to be filled when it's called by other functions, so you don't need to worry too much about providing one yourself.

## NewDataFrame

```go
func NewDataFrame(data [][]interface{}, columns []string, indexCols []string) (DataFrame, error)
```

`NewDataFrame` takes in a 2D slice of data, a list of columns, and a list of index columns, returnign a DataFrame object and an error.

## CreateRangeIndex

```go
func CreateRangeIndex(length int) IndexData
```

`CreateRangeIndex` creates a RangeIndex that spans from 0 to a specified `length`.