# Properties

You can check your `Series`'s properties with these funcitons.

## Accessors

```go
func (s Series) Data() []interface{}

func (s Series) Index() IndexData

func (s Series) Name() string

func (s Series) Dtype() string
```

The fields of `Series` type are private, but `gambas` provides accessors to get these fields.

## ValueCounts

```go
func (s *Series) ValueCounts() (Series, error)
```

`ValueCounts` returns a `Series` containing the number of unique values in a given `Series`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

res, err := col1.ValueCounts()
if err != nil {
    fmt.Println(err)
}

res.Head(10)
```
```
Data            |    Unique Value Count of est_diameter_min     
0.0006089126    |    1                                          
0.0006832112    |    6                                          
0.0008176265    |    5                                          
0.0009216265    |    4                                          
0.0010105434    |    9                                          
0.0010581689    |    31                                         
0.001139082     |    4                                          
0.0011496218    |    3                                          
0.0011528027    |    1                                          
0.0011549282    |    1
```

## IndexHasDuplicateValues

```go
func (s *Series) IndexHasDuplicateValues() (bool, error)
```

`IndexHasDuplicateValues` checks if the `Series` have duplicate index values.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}

col1.SortByIndex(true)
col1.Head(5)

fmt.Println("")

res, err := col1.IndexHasDuplicateValues()
if err != nil {
    fmt.Println(err)
}

fmt.Println(res)
```
```
id         |    est_diameter_min     
2000433    |    23.0438466577        
2000433    |    23.0438466577        
2000433    |    23.0438466577        
2000719    |    2.0443487103         
2001036    |    37.8926498379        

true
```