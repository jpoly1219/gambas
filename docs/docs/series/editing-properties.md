# Editing Properties

You can edit your `Series` object using these functions.

The data used in the example `neo_v2.csv` is NASA's list of Nearest Earth Objects, sourced from [Kaggle](https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects).

## RenameCol

```go
func (s *Series) RenameCol(newName string)
```

`RenameCol` renames the series.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}
fmt.Println(col1.Name())

col1.RenameCol("newName")
fmt.Println(col1.Name())
```
```
est_diameter_min
newName
```

## RenameIndex

```go
func (s *Series) RenameIndex(newNames map[string]string) error
```

`RenameIndex` renames the index of the series.

Input should be a map, where key is the index name to change and value is a new name.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id", "hazardous"})
if err != nil {
    fmt.Println(err)
}

col1, err := df.LocCol("est_diameter_min")
if err != nil {
    fmt.Println(err)
}
fmt.Println(col1.Index().Names())

col1.RenameIndex(map[string]string{"id": "newId", "hazardous": "reallyHazardous"})
fmt.Println(col1.Index().Names())
```
```
[id hazardous]
[newId reallyHazardous]
```