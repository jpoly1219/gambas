# Sorting

In some occasions, you may need to sort your data in different order. `gambas` allows you to sort your Series in a couple of different ways.

The data used in the example `neo_v2.csv` is NASA's list of Nearest Earth Objects, sourced from [Kaggle](https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects).

## SortByIndex

```go
func (s *Series) SortByIndex(ascending bool) error
```

`SortByIndex` sorts the elements in a `Series` by index.

Pass in `true` if you want to sort in ascending order, and `false` for descending order.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}
col1.Head(5)

fmt.Println("")

col1.SortByIndex(true)
col1.Head(5)
```
```
id         |    est_diameter_min     
2162635    |    1.1982708007         
2277475    |    0.2658               
2512244    |    0.7220295577         
3596030    |    0.096506147          
3667127    |    0.2550086879         

id         |    est_diameter_min     
2000433    |    23.0438466577        
2000433    |    23.0438466577        
2000433    |    23.0438466577        
2000719    |    2.0443487103         
2001036    |    37.8926498379
```

## SortByGivenIndex

```go
func (s *Series) SortByGivenIndex(index IndexData) error
```

`SortByGivenIndex` sorts the `Series` by a given index.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}
col1.Head(5)

fmt.Println("")

// create a custom index by shuffling the original one
ciData := make([][]interface{}, len(col1.Data()))
for i := range col1.Data() {
    ciData[i] = append(ciData[i], col1.Index().Index()[rand.Intn(len(col1.Data()))].Value()...)
}

ci, err := gambas.NewIndexData(ciData, []string{"customId"})
if err != nil {
    fmt.Println(err)
}

col1.SortByGivenIndex(ci)
col1.Head(5)
```
```
id         |    est_diameter_min     
2162635    |    1.1982708007         
2277475    |    0.2658               
2512244    |    0.7220295577         
3596030    |    0.096506147          
3667127    |    0.2550086879         

customId    |    est_diameter_min     
3321492     |    0.0930154254         
54053951    |    0.0253837029         
3164412     |    0.0802703167         
2326291     |    0.4174024334         
3445668     |    0.4135756649
```

## SortByValues

```go
func (s *Series) SortByValues(ascending bool) error
```
`SortByValues` sorts the `Series` by its values.

Pass in `true` if you want to sort in ascending order, and `false` for descending order.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}
col1.Head(5)

fmt.Println("")

col1.SortByValues(true)
col1.Head(5)
```
```
id         |    est_diameter_min     
2162635    |    1.1982708007         
2277475    |    0.2658               
2512244    |    0.7220295577         
3596030    |    0.096506147          
3667127    |    0.2550086879         

id          |    est_diameter_min     
3430497     |    0.0006089126         
54106298    |    0.0006832112         
54106298    |    0.0006832112         
54106298    |    0.0006832112         
54106298    |    0.0006832112
```