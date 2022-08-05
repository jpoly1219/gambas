# Sorting

In some occasions, you may need to sort your data in different order. `gambas` allows you to sort your `DataFrame` in a couple of different ways.

The data used in the example `2019.csv` is UN's 2019 World Happiness Report, sourced from [Kaggle](https://www.kaggle.com/datasets/unsdsn/world-happiness).

## SortByIndex

```go
func (df *DataFrame) SortByIndex(ascending bool) error
```

`SortByIndex` sorts the items by index.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), []string{"GDP per capita"})
if err != nil {
    fmt.Println(err)
}
df.Head(5)
fmt.Println("")

df.SortByIndex(true)
df.Head(5)
```
```
GDP per capita    |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
1.34              |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1.383             |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
1.488             |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
1.38              |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
1.396             |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

GDP per capita    |    Overall rank    Country or region           Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0                 |    112             Somalia                     4.668    0                 0.698             0.268                      0.559                           0.243         0.27                         
0.026             |    155             Central African Republic    3.083    0.026             0                 0.105                      0.225                           0.235         0.035                        
0.046             |    145             Burundi                     3.775    0.046             0.447             0.38                       0.22                            0.176         0.18                         
0.073             |    141             Liberia                     3.975    0.073             0.922             0.443                      0.37                            0.233         0.033                        
0.094             |    127             Congo (Kinshasa)            4.418    0.094             1.125             0.357                      0.269                           0.212         0.053                        
```

## SortByValues

```go
func (df *DataFrame) SortByValues(by string, ascending bool) error
```

`SortByValues` sorts the items by values in a selected `Series`.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}
df.Head(5)
fmt.Println("")

df.SortByValues("Score", true)
df.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

       |    Overall rank    Country or region           Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
155    |    156             South Sudan                 2.853    0.306             0.575             0.295                      0.01                            0.202         0.091                        
154    |    155             Central African Republic    3.083    0.026             0                 0.105                      0.225                           0.235         0.035                        
153    |    154             Afghanistan                 3.203    0.35              0.517             0.361                      0                               0.158         0.025                        
152    |    153             Tanzania                    3.231    0.476             0.885             0.499                      0.417                           0.276         0.147                        
151    |    152             Rwanda                      3.334    0.359             0.711             0.614                      0.555                           0.217         0.411                        
```

## SortByColumns

```go
func (df *DataFrame) SortByColumns()
```

`SortByColumns` sorts the columns of the `DataFrame` object.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}
df.Head(5)
fmt.Println("")

df.SortByColumns()
df.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        

     |    Country or region    Freedom to make life choices    GDP per capita    Generosity    Healthy life expectancy    Overall rank    Perceptions of corruption    Score    Social support    
0    |    Finland              0.596                           1.34              0.153         0.986                      1               0.393                        7.769    1.587             
1    |    Denmark              0.592                           1.383             0.252         0.996                      2               0.41                         7.6      1.573             
2    |    Norway               0.603                           1.488             0.271         1.028                      3               0.341                        7.554    1.582             
3    |    Iceland              0.591                           1.38              0.354         1.026                      4               0.118                        7.494    1.624             
4    |    Netherlands          0.557                           1.396             0.322         0.999                      5               0.298                        7.488    1.522             
```

## SortIndexColFirst

```go
func (df *DataFrame) SortIndexColFirst()
```

`SortIndexColFirst` puts the index column at the front.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), []string{"Score"})
if err != nil {
    fmt.Println(err)
}
df.Head(5)

df.SortIndexColFirst()
df.Head(5)
```
```
Score    |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
7.769    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
7.6      |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
7.554    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
7.494    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
7.488    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        
Score    |    Score    Country or region    Overall rank    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
7.769    |    7.769    Finland              1               1.34              1.587             0.986                      0.596                           0.153         0.393                        
7.6      |    7.6      Denmark              2               1.383             1.573             0.996                      0.592                           0.252         0.41                         
7.554    |    7.554    Norway               3               1.488             1.582             1.028                      0.603                           0.271         0.341                        
7.494    |    7.494    Iceland              4               1.38              1.624             1.026                      0.591                           0.354         0.118                        
7.488    |    7.488    Netherlands          5               1.396             1.522             0.999                      0.557                           0.322         0.298                        
```