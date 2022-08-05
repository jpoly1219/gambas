# Merging

You can merge two `DataFame` objects.

The data used in the example `2019.csv` is UN's 2019 World Happiness Report, sourced from [Kaggle](https://www.kaggle.com/datasets/unsdsn/world-happiness).

## MergeDfsHorizontally

```go
func (df *DataFrame) MergeDfsHorizontally(target DataFrame) (DataFrame, error)
```

`MergeDfsHorizontally` merges two `DataFrame` objects side by side.

The target `DataFrame` will always be appended to the right of the source `DataFrame`. Index will reset and become a `RangeIndex`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df1, err := df.LocCols("Overall rank", "Country or region", "Score", "GDP per capita", "Social support")
if err != nil {
    fmt.Println(err)
}

df2, err := df.LocCols("Healthy life expectancy", "Freedom to make life choices", "Generosity", "Perceptions of corruption")
if err != nil {
    fmt.Println(err)
}

df1.Head(5)
fmt.Println("")
df2.Head(5)
fmt.Println("")

res, err := df1.MergeDfsHorizontally(df2)
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

     |    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    0.986                      0.596                           0.153         0.393                        
1    |    0.996                      0.592                           0.252         0.41                         
2    |    1.028                      0.603                           0.271         0.341                        
3    |    1.026                      0.591                           0.354         0.118                        
4    |    0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        
```

## MergeDfsVertically

```go
func (df *DataFrame) MergeDfsVertically(target DataFrame) (DataFrame, error)
```

`MergeDfsVertically` stacks two `DataFrame` objects vertically.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df1, err := df.LocRows([]interface{}{0}, []interface{}{1}, []interface{}{2})
if err != nil {
    fmt.Println(err)
}

df2, err := df.LocRows([]interface{}{3}, []interface{}{4}, []interface{}{5})
if err != nil {
    fmt.Println(err)
}

df1.Print()
fmt.Println("")
df2.Print()
fmt.Println("")

res, err := df1.MergeDfsVertically(df2)
if err != nil {
    fmt.Println(err)
}

res.Print()
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        
5    |    6               Switzerland          7.48     1.452             1.526             1.052                      0.572                           0.263         0.343                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        
5    |    6               Switzerland          7.48     1.452             1.526             1.052                      0.572                           0.263         0.343                        
```