# Indexing

You can index your `DataFrame` with `gambas`'s built-in indexing tools.

## LocRows

```go
func (df *DataFrame) LocRows(rows ...[]interface{}) (DataFrame, error)
```
`LocRows` returns a set of rows as a new `DataFrame` object, given a list of labels.

You are only allowed to pass in the indices of the `DataFrame` as rows. For multiindex, pass in either the entire index tuple or just the first index.

### Example 1: Single index

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), []string{"Country or region"})
if err != nil {
    fmt.Println(err)
}

res, err := df.LocRows([]interface{}{"Canada"}, []interface{}{"United States"}, []interface{}{"Mexico"})
if err != nil {
    fmt.Println(err)
}

res.Print()
```
```
Country or region    |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
Canada               |    9               Canada               7.278    1.365             1.505             1.039                      0.584                           0.285         0.308                        
United States        |    19              United States        6.892    1.433             1.457             0.874                      0.454                           0.28          0.128                        
Mexico               |    23              Mexico               6.595    1.07              1.323             0.861                      0.433                           0.074         0.073                        
```

### Example 2: Multiindex

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), []string{"Overall rank", "Country or region"})
if err != nil {
    fmt.Println(err)
}

// this will accomplish the same thing as well:
// res, err := df.LocRows([]interface{}{9, "Canada"}, []interface{}{19, "United States"}, []interface{}{23, "Mexico"})
res, err := df.LocRows([]interface{}{9}, []interface{}{19}, []interface{}{23})
if err != nil {
    fmt.Println(err)
}

res.Print()
```
```
Country or region    |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
Canada               |    9               Canada               7.278    1.365             1.505             1.039                      0.584                           0.285         0.308                        
United States        |    19              United States        6.892    1.433             1.457             0.874                      0.454                           0.28          0.128                        
Mexico               |    23              Mexico               6.595    1.07              1.323             0.861                      0.433                           0.074         0.073
```

## LocRowsItems

```go
func (df *DataFrame) LocRowsItems(rows ...[]interface{}) ([][]interface{}, error)
```

`LocRowsItems` acts the exact same as `LocRows`, but returns data as `[][]interface{}` instead of `DataFrame`. For usage, refer to `LocRows`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), []string{"Country or region"})
if err != nil {
    fmt.Println(err)
}

res, err := df.LocRowsItems([]interface{}{"Canada"}, []interface{}{"United States"}, []interface{}{"Mexico"})
if err != nil {
    fmt.Println(err)
}

fmt.Println(res)
```
```
[[9 Canada 7.278 1.365 1.505 1.039 0.584 0.285 0.308] [19 United States 6.892 1.433 1.457 0.874 0.454 0.28 0.128] [23 Mexico 6.595 1.07 1.323 0.861 0.433 0.074 0.073]]
```

## LocCol

```go
func (df *DataFrame) LocCol(col string) (Series, error) 
```

`LocCol` returns a column as a new `Series` object.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

res, err := df.LocCol("Country or region")
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Country or region     
0    |    Finland               
1    |    Denmark               
2    |    Norway                
3    |    Iceland               
4    |    Netherlands
```

## LocCols

```go
func (df *DataFrame) LocCols(cols ...string) (DataFrame, error)
```

`LocCols` returns a set of columns as a new `DataFrame` object, given a list of labels.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

res, err := df.LocCols("Country or region", "GDP per capita", "Healthy life expectancy")
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Country or region    GDP per capita    Healthy life expectancy    
0    |    Finland              1.34              0.986                      
1    |    Denmark              1.383             0.996                      
2    |    Norway               1.488             1.028                      
3    |    Iceland              1.38              1.026                      
4    |    Netherlands          1.396             0.999
```

## LocColsItems

```go
func (df *DataFrame) LocColsItems(cols ...string) ([][]interface{}, error)
```

`LocColsItems` acts the exact same as `LocCols`, but returns data as `[][]interface{}` instead of `DataFrame`. For usage, refer to `LocCols`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

res, err := df.LocColsItems("Country or region", "GDP per capita", "Healthy life expectancy")
if err != nil {
    fmt.Println(err)
}

fmt.Println(res)
```
```
[[Finland Denmark Norway Iceland Netherlands ...] [1.34 1.383 1.488 1.38 1.396 ...] [0.986 0.996 1.028 1.026 0.999 ...]]

(truncated in documentation due to length constraints)
```

## Loc

```go
func (df *DataFrame) Loc(cols []string, rows ...[]interface{}) (DataFrame, error)
```

`Loc` indexes the `DataFrame` object given a slice of row and column labels, and returns the result as a new `DataFrame` object.

You are only allowed to pass in indices of the `DataFrame` as rows.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

res, err := df.Loc([]string{"Country or region", "GDP per capita"}, []interface{}{0}, []interface{}{1}, []interface{}{2})
if err != nil {
    fmt.Println(err)
}

res.Print()
```
```
     |    Country or region    GDP per capita    
0    |    Finland              1.34              
1    |    Denmark              1.383             
2    |    Norway               1.488
```