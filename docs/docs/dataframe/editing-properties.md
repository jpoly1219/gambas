# Editing Properties

You can edit your `DataFrame` object using these functions.

The data used in the example `2019.csv` is UN's 2019 World Happiness Report, sourced from [Kaggle](https://www.kaggle.com/datasets/unsdsn/world-happiness).

## NewCol

```go
func (df *DataFrame) NewCol(colname string, data []interface{}) (DataFrame, error)
```

`NewCol` creates a new column with the given data and column name.

To create a blank column, pass in `nil`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")
if err != nil {
    fmt.Println(err)
}

df1.Head(5)
fmt.Println("")

res, err := df1.NewCol("NewCol", nil)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    
0    |    1               Finland              7.769    1.34              1.587             
1    |    2               Denmark              7.6      1.383             1.573             
2    |    3               Norway               7.554    1.488             1.582             
3    |    4               Iceland              7.494    1.38              1.624             
4    |    5               Netherlands          7.488    1.396             1.522             

     |    Overall rank    Country or region    Score    GDP per capita    Social support    NewCol    
0    |    1               Finland              7.769    1.34              1.587             NaN       
1    |    2               Denmark              7.6      1.383             1.573             NaN       
2    |    3               Norway               7.554    1.488             1.582             NaN       
3    |    4               Iceland              7.494    1.38              1.624             NaN       
4    |    5               Netherlands          7.488    1.396             1.522             NaN       
```

## NewDerivedCol

```go
func (df *DataFrame) NewDerivedCol(colname, srcCol string) (DataFrame, error)
```

`NewDerivedCol` creates a new column derived from an existing column.

It copies over the data from `srcCol` into a new column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")
if err != nil {
    fmt.Println(err)
}

df1.Head(5)
fmt.Println("")

res, err := df1.NewDerivedCol("NewCol", "Social support")
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    
0    |    1               Finland              7.769    1.34              1.587             
1    |    2               Denmark              7.6      1.383             1.573             
2    |    3               Norway               7.554    1.488             1.582             
3    |    4               Iceland              7.494    1.38              1.624             
4    |    5               Netherlands          7.488    1.396             1.522             

     |    Overall rank    Country or region    Score    GDP per capita    Social support    NewCol    
0    |    1               Finland              7.769    1.34              1.587             1.587     
1    |    2               Denmark              7.6      1.383             1.573             1.573     
2    |    3               Norway               7.554    1.488             1.582             1.582     
3    |    4               Iceland              7.494    1.38              1.624             1.624     
4    |    5               Netherlands          7.488    1.396             1.522             1.522     
```

## RenameCol

```go
func (df *DataFrame) RenameCol(colnames map[string]string) error
```

`RenameCol` renames columns in a `DataFrame`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")
if err != nil {
    fmt.Println(err)
}

df1.Head(5)
fmt.Println("")

err = df1.RenameCol(map[string]string{"Social support": "Some kind of support"})
if err != nil {
    fmt.Println(err)
}

df1.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    
0    |    1               Finland              7.769    1.34              1.587             
1    |    2               Denmark              7.6      1.383             1.573             
2    |    3               Norway               7.554    1.488             1.582             
3    |    4               Iceland              7.494    1.38              1.624             
4    |    5               Netherlands          7.488    1.396             1.522             

index does not exist: Social support
     |    Overall rank    Country or region    Score    GDP per capita    Some kind of support    
0    |    1               Finland              7.769    1.34              1.587                   
1    |    2               Denmark              7.6      1.383             1.573                   
2    |    3               Norway               7.554    1.488             1.582                   
3    |    4               Iceland              7.494    1.38              1.624                   
4    |    5               Netherlands          7.488    1.396             1.522                   
```

## DropNaN

```go
func (df *DataFrame) DropNaN(axis int) (DataFrame, error)
```

`DropNaN` drops rows or columns with `NaN` values.

Specify axis to choose whether to remove rows with `NaN` or columns with `NaN`. `axis=0` is row, `axis=1` is column.

### Example 1: Dropping rows

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019-with-nan.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.DropNaN(0)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              NaN      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             NaN                        0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           NaN           0.118                        
4    |    5               Netherlands          7.488    1.396             NaN               0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
5    |    6               Switzerland          7.48     1.452             1.526             1.052                      0.572                           0.263         0.343                        
6    |    7               Sweden               7.343    1.387             1.487             1.009                      0.574                           0.267         0.373                        
7    |    8               New Zealand          7.307    1.303             1.557             1.026                      0.585                           0.33          0.38                         
8    |    9               Canada               7.278    1.365             1.505             1.039                      0.584                           0.285         0.308                        
```

### Example 1: Dropping columns

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019-with-nan.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.DropNaN(1)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              NaN      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             NaN                        0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           NaN           0.118                        
4    |    5               Netherlands          7.488    1.396             NaN               0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    GDP per capita    Healthy life expectancy    Freedom to make life choices    Perceptions of corruption    
0    |    1               Finland              1.34              0.986                      0.596                           0.393                        
1    |    2               Denmark              1.383             0.996                      0.592                           0.41                         
2    |    3               Norway               1.488             NaN                        0.603                           0.341                        
3    |    4               Iceland              1.38              1.026                      0.591                           0.118                        
4    |    5               Netherlands          1.396             0.999                      0.557                           0.298                        
```