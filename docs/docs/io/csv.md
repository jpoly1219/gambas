# CSV

## ReadCsv

```go
func ReadCsv(pathToFile string, indexCols []string) (DataFrame, error)
```

`ReadCsv` reads a CSV file and returns a new DataFrame object. It is recommended to generate pathToFile using the `path/filepath` package for cross-platform compatibility.

```go
/*
people.csv

Name,Age,Height
Alice,54,163
Bob,26,167
Charlie,17,175
*/

myDf, err := gambas.ReadCsv(filepath.Join(".", "people.csv"), []string{"Name"})
if err != nil {
    fmt.Println(err)
}

myDf.Print()
```
```
Name       |    Name       Age    Height    
Alice      |    Alice      54     163       
Bob        |    Bob        26     167       
Charlie    |    Charlie    17     175   
```

## WriteCsv

```go
func WriteCsv(df DataFrame, pathToFile string, skipColumnLabel bool) (os.FileInfo, error)
```

`WriteCsv` writes a DataFrame object to CSV file. It is recommended to generate pathToFile using the `path/filepath` package for cross-platform compatibility.

You can choose whether the output CSV file should contain column labels or not. If you choose to include it, pass `false` to `skipColumnLabel`. Otherwise, pass `true`.

### Example 1: Include column labels

```go
myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
myCols := []string{"group a", "group b", "group c"}
myIndexCols := []string{"group a"}

myDf, err := gambas.NewDataFrame(myData, myCols, myIndexCols)
if err != nil {
    fmt.Println(err)
}

gambas.WriteCsv(myDf, filepath.Join(".", "output.csv"), false)
```
```
// output.csv

group a,group b,group c
1,4,7
2,5,8
3,6,9
```

### Example 2: Skip column labels

```go
myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
myCols := []string{"group a", "group b", "group c"}
myIndexCols := []string{"group a"}

myDf, err := gambas.NewDataFrame(myData, myCols, myIndexCols)
if err != nil {
    fmt.Println(err)
}

gambas.WriteCsv(myDf, filepath.Join(".", "output.csv"), true)
```
```
// output.csv

1,4,7
2,5,8
3,6,9
```