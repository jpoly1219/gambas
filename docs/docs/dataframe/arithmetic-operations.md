# Arithmetic Operations

You can apply column-wide arithmetic operations to each of your `DataFrame` object's columns.

The data used in the example `2019.csv` is UN's 2019 World Happiness Report, sourced from [Kaggle](https://www.kaggle.com/datasets/unsdsn/world-happiness).

## ColAdd

```go
func (df *DataFrame) ColAdd(colname string, value float64) (DataFrame, error)
```

`ColAdd` adds the given value to each element in the specified column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColAdd("GDP per capita", 2.5)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    3.84              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      3.883             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    3.988             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    3.88              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    3.896             1.522             0.999                      0.557                           0.322         0.298                        
```

## ColSub

```go
func (df *DataFrame) ColSub(colname string, value float64) (DataFrame, error)
```

`ColSub` subtracts the given value from each element in the specified column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColSub("GDP per capita", 0.5)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    0.8400000000000001    1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      0.883                 1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    0.988                 1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    0.8799999999999999    1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    0.8959999999999999    1.522             0.999                      0.557                           0.322         0.298                        
```

## ColMul

```go
func (df *DataFrame) ColMul(colname string, value float64) (DataFrame, error)
```

`ColMul` multiplies each element in the specified column by the given value.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColMul("GDP per capita", 1.5)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    2.0100000000000002    1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      2.0745                1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    2.232                 1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    2.07                  1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    2.094                 1.522             0.999                      0.557                           0.322         0.298                        
```

## ColDiv

```go
func (df *DataFrame) ColDiv(colname string, value float64) (DataFrame, error)
```

`ColDiv` divides each element in the specified column by the given value.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColDiv("GDP per capita", 1.5)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    0.8933333333333334    1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      0.922                 1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    0.992                 1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    0.9199999999999999    1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    0.9306666666666666    1.522             0.999                      0.557                           0.322         0.298                        
```

## ColMod

```go
func (df *DataFrame) ColMod(colname string, value float64) (DataFrame, error)
```

`ColMod` applies modulus calculations on each element in the specified column, returning the remainder.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColMod("GDP per capita", 1.0)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita        Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    0.3400000000000001    1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      0.383                 1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    0.488                 1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    0.3799999999999999    1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    0.3959999999999999    1.522             0.999                      0.557                           0.322         0.298                        
```

## ColGt

```go
func (df *DataFrame) ColGt(colname string, value float64) (DataFrame, error)
```

`ColGt` checks if each element in the specified column is greater than the given value.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColGt("GDP per capita", 1.4)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    false             1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      false             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    true              1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    false             1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    false             1.522             0.999                      0.557                           0.322         0.298                        
```

## ColLt

```go
func (df *DataFrame) ColLt(colname string, value float64) (DataFrame, error)
```

`ColLt` checks if each element in the specified column is less than the given value.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColLt("GDP per capita", 1.4)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    true              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      true              1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    false             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    true              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    true              1.522             0.999                      0.557                           0.322         0.298                        
```

## ColEq

```go
func (df *DataFrame) ColEq(colname string, value float64) (DataFrame, error)
```

`ColEq` checks if each element in the specified column is equal to the given value.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
fmt.Println("")

res, err := df.ColEq("GDP per capita", 1.38)
if err != nil {
    fmt.Println(err)
}

res.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    false             1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      false             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    false             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    true              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    false             1.522             0.999                      0.557                           0.322         0.298                        
```