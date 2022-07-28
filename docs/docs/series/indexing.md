# Indexing

## At

```go
func (s *Series) At(ind ...interface{}) (interface{}, error)
```

`At` returns an element at a given index.

For multiindex, you need to pass in the whole index tuple.

### Example 1: Normal

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
// tbd
```
```
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

`Loc` returns a range of data at given rows.