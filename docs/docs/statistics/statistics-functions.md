# Statistics Functions

The data used in the example `neo_v2.csv` is NASA's list of Nearest Earth Objects, sourced from [Kaggle](https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects).

## Count

```go
func Count(dataset []interface{}) StatsResult
```

`Count` counts the number of non-NaN elements in a `dataset`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Count(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Count
90836
<nil>
```

## Mean

```go
func Mean(dataset []interface{}) StatsResult
```

`Mean` returns the mean of the elements in a `dataset`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Mean(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Mean
0.127
<nil>
```

## Median

```go
func Median(dataset []interface{}) StatsResult 
```

`Median` returns the median of the elements in a `dataset`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Median(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Median
0.048
<nil>
```

## Std

```go
func Std(dataset []interface{}) StatsResult 
```

`Std` returns the sample standard deviation of the elements in a `dataset`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Std(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Std
0.299
<nil>
```

## Min

```go
func Min(dataset []interface{}) StatsResult
```

`Min` returns the smallest element in a `dataset`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Min(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Min
0.0006089126
<nil>
```

## Max

```go
func Max(dataset []interface{}) StatsResult
```

`Max` returns the largest element is a `dataset`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Max(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Max
37.8926498379
<nil>
```

## Q1

```go
func Q1(dataset []interface{}) StatsResult
```

`Q1` returns the lower quartile (25%) of the elements in a `dataset`. This does not include the median during calculation.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Q1(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Q1
0.0192555078
<nil>
```

## Q2

```go
func Q2(dataset []interface{}) StatsResult
```

`Q2` returns the middle quartile (50%) of the elements in a `dataset`. This accomplishes the same thing as `Median`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Q2(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Q2
0.048
<nil>
```

## Q3

```go
func Q3(dataset []interface{}) StatsResult
```

`Q3` returns the upper quartile (75%) of the elements in a `dataset`. This does not include the median during calculation.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocColsItems("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := gambas.Q3(col1[0])
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Q3
0.1434019235
<nil>
```