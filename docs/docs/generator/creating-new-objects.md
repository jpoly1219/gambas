# Creating New Objects

`gambas` provides you with generator functions that can create a RangeIndex (an index that spans from 0 to a certain length), a Series, or a DataFrame object. Usually, you won't be creating objects from scratch using the generator functions. Instead, you will use I/O functions to generate objects out of pre-existing data.

## NewSeries

```go
func NewSeries(data []interface{}, name string, index *IndexData) (Series, error)
```

`NewSeries` takes in a slice of data, a name, and an IndexData object, returning a Series object and an error. 

`data` should hold items of same data type. Supported data types include `int`, `float64`, `bool`, and `string`.

Leave `index` parameter as `nil`, and `NewSeries` will automatically generate a RangeIndex.

### Example 1: Single index

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
     |    Fruit     
0    |    apple     
1    |    banana    
2    |    cherry 
```

### Example 2: Multiindex

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"
myIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"c", "red"}}

myIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})
if err != nil {
    fmt.Println(err)
}

mySeries, err := gambas.NewSeries(myData, myName, &myIndexData)
if err != nil {
    fmt.Println(err)
}
mySeries.Print()
```
```
key    color     |    Fruit     
a      red       |    apple     
b      yellow    |    banana    
c      red       |    cherry
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
    fmt.Println(err)
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

## NewIndexData

```go
func NewIndexData(index [][]interface{}, names []string) (IndexData, error)
```

`NewIndexData` creates a custom `IndexData` object.

`index` should be a 2D slice that contains index tuples. For single index, the index tuple will contain only one item. For multiindex, the index tuple will contain more than one item.

`names` are labels for each index. The length should match that of `index` items.

## Example 1: Single index

```go
myIndex := [][]interface{}{{"a"}, {"b"}, {"c"}}

myIndexData, err := gambas.NewIndexData(myIndex, []string{"alphabet"})
if err != nil {
    fmt.Println(err)
}
```

### Example 2: Multiindex

```go
myIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"c", "red"}}

myIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})
if err != nil {
    fmt.Println(err)
}
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