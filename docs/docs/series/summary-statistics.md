# Summary Statistics

`gambas` provides summary statistics functions for `Series` objects. All statistics functions return a `StatsResult` object defined per below. 

```go
type StatsResult struct {
    UsedFunc string
    Result   float64
    Err      error
}
```

`UsedFunc` denotes what function has been used. `Result` is the result of the calculation. `Err` holds any errors that the function has encountered.

The data used in the example `neo_v2.csv` is NASA's list of Nearest Earth Objects, sourced from [Kaggle](https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects).

## Count

```go
func (s *Series) Count() StatsResult
```

`Count` counts the number of non-NaN elements in a column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Count()
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
func (s *Series) Mean() StatsResult
```

`Mean` returns the mean of the elements in a column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Mean()
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
func (s *Series) Median() StatsResult
```

`Median` returns the median of the elements in a column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Median()
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
func (s *Series) Std() StatsResult
```

`Std` returns the sample standard deviation of the elements in a column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Std()
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
func (s *Series) Min() StatsResult
```

`Min` returns the smallest element in a column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Min()
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
func (s *Series) Max() StatsResult
```

`Max` returns the largest element is a column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Max()
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
func (s *Series) Q1() StatsResult
```

`Q1` returns the lower quartile (25%) of the elements in a column. This does not include the median during calculation.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Q1()
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
func (s *Series) Q2() StatsResult
```

`Q2` returns the middle quartile (50%) of the elements in a column. This accomplishes the same thing as `Median`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Q2()
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
func (s *Series) Q3() StatsResult
```

`Q3` returns the upper quartile (75%) of the elements in a column. This does not include the median during calculation.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res := col1.Q3()
fmt.Println(res.UsedFunc)
fmt.Println(res.Result)
fmt.Println(res.Err)
```
```
Q3
0.1434019235
<nil>
```

## Describe

```go
func (s *Series) Describe() ([]StatsResult, error)
```

`Describe` runs through the most commonly used statistics functions and prints the output.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res, err := col1.Describe()
if err != nil {
    fmt.Println(err)
}
fmt.Println(res)
```
```
Count: 90836
Mean: 0.127
Median: 0.048
Std: 0.299
Min: 0.0006089126
Max: 37.8926498379
Q1: 0.0192555078
Q2: 0.048
Q3: 0.1434019235
[{Count 90836 <nil>} {Mean 0.127 <nil>} {Median 0.048 <nil>} {Std 0.299 <nil>} {Min 0.0006089126 <nil>} {Max 37.8926498379 <nil>} {Q1 0.0192555078 <nil>} {Q2 0.048 <nil>} {Q3 0.1434019235 <nil>}]
```