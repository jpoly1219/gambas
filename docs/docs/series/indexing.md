# Indexing

You can index your `Series` with `gambas`'s built-in indexing tools.

## At

```go
func (s *Series) At(ind ...interface{}) (interface{}, error)
```

`At` returns an element at a given index.

For multiindex, you need to pass in the whole index tuple.

### Example 1: Single index

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.At(1)
if err != nil {
    fmt.Println(err)
}
fmt.Println(res)
```
```
banana
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

res, err := mySeries.At("b", "yellow")
if err != nil {
    fmt.Println(err)
}
fmt.Println(res)
```
```
key    color     |    Fruit     
a      red       |    apple     
b      yellow    |    banana    
c      red       |    cherry    
banana
```

## IAt

```go
func (s *Series) IAt(ind int) (interface{}, error)
```

`IAt` returns an element at a given integer index.

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.IAt(2)
if err != nil {
    fmt.Println(err)
}
fmt.Println(res)
```
```
cherry
```

## Loc

```go
func (s *Series) Loc(idx ...[]interface{}) (Series, error)
```

`Loc` accepts index tuples and returns a `Series` object containing data at the given rows.

Each `idx` item should contain the index of the data you would like to query. For multiindex `Series`, you can either pass in the whole index tuple, or the first index.

### Example 1: Indexing a single item

```go
myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}
myName := "Fruit"
mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.Loc([]interface{}{2})
if err != nil {
    fmt.Println(err)
}
res.Print()
```
```
     |    Fruit     
2    |    cherry
```

### Example 2: Indexing multiple items

```go
myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}
myName := "Fruit"
mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.Loc([]interface{}{2}, []interface{}{4})
if err != nil {
    fmt.Println(err)
}
res.Print()
```
```
     |    Fruit     
2    |    cherry    
4    |    kiwi
```

### Example 3: Multiindex (all index)

```go
myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}
myName := "Fruit"
myIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"c", "red"}, {"d", "brown"}, {"e", "green"}}

myIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})
if err != nil {
    fmt.Println(err)
}

mySeries, err := gambas.NewSeries(myData, myName, &myIndexData)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.Loc([]interface{}{"a", "red"}, []interface{}{"d", "brown"})
if err != nil {
    fmt.Println(err)
}
res.Print()
```
```
key    color    |    Fruit     
a      red      |    apple     
d      brown    |    durian
```

### Example 4: Multiindex (first index)

```go
myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}
myName := "Fruit"
myIndex := [][]interface{}{{"a", "red"}, {"b", "yellow"}, {"a", "red"}, {"b", "brown"}, {"a", "green"}}

myIndexData, err := gambas.NewIndexData(myIndex, []string{"key", "color"})
if err != nil {
    fmt.Println(err)
}

mySeries, err := gambas.NewSeries(myData, myName, &myIndexData)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.Loc([]interface{}{"a"})
if err != nil {
    fmt.Println(err)
}
res.Print()
```
```
key    color    |    Fruit     
a      red      |    apple     
a      red      |    cherry    
a      green    |    kiwi
```

## LocItems

```go
func (s *Series) LocItems(idx ...[]interface{}) ([]interface{}, error)
```

`LocItems` acts the exact same as Loc, but returns data as `[]interface{}` instead of `Series`. For usage, refer to `Loc`.

```go
myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.LocItems([]interface{}{1}, []interface{}{3}, []interface{}{4})
if err != nil {
    fmt.Println(err)
}
fmt.Println(res)
```
```
[banana durian kiwi]
```

## ILoc

```go
func (s *Series) ILoc(min, max int) ([]interface{}, error)
```

`ILoc` returns an array of elements at a given integer index range.

```go
myData := []interface{}{"apple", "banana", "cherry", "durian", "kiwi"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

res, err := mySeries.ILoc(1, 3)
if err != nil {
    fmt.Println(err)
}
fmt.Println(res)
```
```
[banana cherry]
```