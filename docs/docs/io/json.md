# JSON

## ReadJsonByColumns

```go
func ReadJsonByColumns(pathToFile string, indexCols []string) (DataFrame, error)
```

`ReadJsonByColumns` reads a JSON file and returns a new DataFrame object. It is recommended to generate `pathToFile` using the `path/filepath` package for cross-platform compatibility.

The JSON file should be in this format:
`{"col1":[val1, val2, ...], "col2":[val1, val2, ...], ...}`

You can either set a column to be the index, or set it as `nil`. If `nil`, a new RangeIndex will be created.
Your index column should not have any missing values. Columns will be alphabetically sorted, but the index column will always come first.

```go
/*
people.json

{
    "Name": ["Avery", "Bradley", "Candice"],
    "Age": [19.0, 26.0, 23.0],
    "Sex": ["Male", "Male", "Female"]
}
*/

myDf, err := gambas.ReadJsonByColumns(filepath.Join(".", "people.json"), nil)
if err != nil {
    fmt.Println(err)
}

myDf.Print()
```
```
     |    Age    Name       Sex       
0    |    19     Avery      Male      
1    |    26     Bradley    Male      
2    |    23     Candice    Female
```

## ReadJsonStream

```go
func ReadJsonStream(pathToFile string, indexCols []string) (DataFrame, error)
```
`ReadJsonStream` reads a JSON stream and returns a new DataFrame object. It is recommended to generate `pathToFile` using the `path/filepath` package for cross-platform compatibility.

The JSON file should be in this format: `{"col1":val1, "col2":val2, ...}{"col1":val1, "col2":val2, ...}`

You can either set a column to be the index, or set it as `nil`. If `nil`, a new RangeIndex will be created.
Your index column should not have any missing values. Columns will be alphabetically sorted, but the index column will always come first.

```go
/*
people.json

[
    {
        "Name": "Avery",
        "Age": 19.0,
        "Sex": "Male"
    },
    {
        "Name": "Bradley",
        "Age": 26.0,
        "Sex": "Male"
    },
    {
        "Name": "Candice",
        "Age": 23.0,
        "Sex": "Female"
    }
]
*/

myDf, err := gambas.ReadJsonStream(filepath.Join(".", "people.json"), nil)
if err != nil {
    fmt.Println(err)
}

myDf.Print()
```
```
     |    Age    Name       Sex       
0    |    19     Avery      Male      
1    |    26     Bradley    Male      
2    |    23     Candice    Female
```

## WriteJson

```go
func WriteJson(df DataFrame, pathToFile string) (os.FileInfo, error)
```

`WriteJson` writes a DataFrame object to a file. It is recommended to generate `pathToFile` using the `path/filepath` package for cross-platform compatibility.

```go
myData := [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
myCols := []string{"group a", "group b", "group c"}
myIndexCols := []string{"group a"}

myDf, err := gambas.NewDataFrame(myData, myCols, myIndexCols)
if err != nil {
    fmt.Println(err)
}

gambas.WriteJson(myDf, filepath.Join(".", "output.json"))
```
```json
// output.json

{
	"series": [
		{
			"data": [
				1,
				2,
				3
			],
			"name": "group a",
			"dtype": "int"
		},
		{
			"data": [
				4,
				5,
				6
			],
			"name": "group b",
			"dtype": "int"
		},
		{
			"data": [
				7,
				8,
				9
			],
			"name": "group c",
			"dtype": "int"
		}
	],
	"columns": [
		"group a",
		"group b",
		"group c"
	]
}
```