# Creating New Objects

`gambas` provides you with generator functions that can create a RangeIndex (an index that spans from 0 to a certain length), a Series, or a DataFrame object. Usually, you won't be creating objects from scratch using the generator functions. Instead, you will use I/O functions to generate objects out of pre-existing data.

## NewSeries

```go
func NewSeries(data []interface{}, name string, index *IndexData) (Series, error)
```

`NewSeries` takes in a slice of data, a name, and an IndexData object, returning a Series object and an error. 

`data` should hold items of same data type. Supported data types include `int`, `float64`, `bool`, and `string`.

Leave `index` parameter as `nil`, and `NewSeries` will automatically generate a RangeIndex. The `index` field is meant to be filled when it's called by other functions.

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

mySeries.Print()
```
```
     Fruit     
0    apple     
1    banana    
2    cherry
```

## NewDataFrame

```go
func NewDataFrame(data [][]interface{}, columns []string, indexCols []string) (DataFrame, error)
```

`NewDataFrame` takes in a 2D slice of data, a list of columns, and a list of index columns, returnign a DataFrame object and an error.

If `indexCols` is `nil`, `NewDataFrame` will generate a RangeIndex.

Index will be printed to the side, separated from the other columns with a `"|"` separator.

### Example 1: Specifying an index column
```go
myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
myCols := []string{"group a", "group b", "group c"}
myIndexCols := []string{"group a"}

myDf, err := gambas.NewDataFrame(myData, myCols, myIndexCols)
if err != nil {
    return err
}

myDf.Print()
```
```
group a    |    group a    group b    group c    
1          |    1          4          7          
2          |    2          5          8          
3          |    3          6          9  
```

### Example 2: Leaving `indexCols` as `nil`

```go
myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
myCols := []string{"group a", "group b", "group c"}

myDf, err := gambas.NewDataFrame(myData, myCols, nil)
if err != nil {
    fmt.Println(err)
}

myDf.Print()
```
```
     |    group a    group b    group c    
0    |    1          4          7          
1    |    2          5          8          
2    |    3          6          9       
```

### Example 3: Specifying multiple index columns

```go
myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
myCols := []string{"group a", "group b", "group c"}
myIndexCols := []string{"group a", "group b"}

myDf, err := gambas.NewDataFrame(myData, myCols, myIndexCols)
if err != nil {
    fmt.Println(err)
}

myDf.Print()
```
```
group a    group b    |    group a    group b    group c    
1          4          |    1          4          7          
2          5          |    2          5          8          
3          6          |    3          6          9        
```

## CreateRangeIndex

```go
func CreateRangeIndex(length int) IndexData
```

`CreateRangeIndex` creates a RangeIndex that spans from 0 to a specified `length`.

```go
myRangeIndex := CreateRangeIndex(5)
```
```
IndexData{
    []Index{
        {0, []interface{}{0}},
        {1, []interface{}{1}},
        {2, []interface{}{2}},
        {3, []interface{}{3}},
        {4, []interface{}{4}},
    },
    []string{""},
},
```