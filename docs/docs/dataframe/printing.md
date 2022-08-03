# Printing

`gambas` provides several options to print your `DataFrame`.

The data used in the example `2019.csv` is UN's 2019 World Happiness Report, sourced from [Kaggle](https://www.kaggle.com/datasets/unsdsn/world-happiness).

## Print

```go
func (df *DataFrame) Print()
```

`Print` prints all data in a `DataFrame` object.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Print()
```
```
       |    Overall rank    Country or region           Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0      |    1               Finland                     7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1      |    2               Denmark                     7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2      |    3               Norway                      7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3      |    4               Iceland                     7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4      |    5               Netherlands                 7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        
5      |    6               Switzerland                 7.48     1.452             1.526             1.052                      0.572                           0.263         0.343                        
6      |    7               Sweden                      7.343    1.387             1.487             1.009                      0.574                           0.267         0.373                        
7      |    8               New Zealand                 7.307    1.303             1.557             1.026                      0.585                           0.33          0.38                         
8      |    9               Canada                      7.278    1.365             1.505             1.039                      0.584                           0.285         0.308                        
9      |    10              Austria                     7.246    1.376             1.475             1.016                      0.532                           0.244         0.226                        
10     |    11              Australia                   7.228    1.372             1.548             1.036                      0.557                           0.332         0.29                         
                      
... (truncated in documentation due to length constraints)
```

## PrintRange

```go
func (df *DataFrame) PrintRange(start, end int)
```

`PrintRange` prints data in a `DataFrame` object at a given range.

Index starts at 0.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.PrintRange(1, 5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                                   
```

## Head

```go
func (df *DataFrame) Head(howMany int)
```

`Head` prints the first `howMany` items in a `DataFrame` object.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Head(5)
```
```
     |    Overall rank    Country or region    Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
0    |    1               Finland              7.769    1.34              1.587             0.986                      0.596                           0.153         0.393                        
1    |    2               Denmark              7.6      1.383             1.573             0.996                      0.592                           0.252         0.41                         
2    |    3               Norway               7.554    1.488             1.582             1.028                      0.603                           0.271         0.341                        
3    |    4               Iceland              7.494    1.38              1.624             1.026                      0.591                           0.354         0.118                        
4    |    5               Netherlands          7.488    1.396             1.522             0.999                      0.557                           0.322         0.298                        
```

## Tail

```go
func (df *DataFrame) Tail(howMany int)
```

`Tail` prints the last `howMany` items in a `DataFrame` object.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "2019.csv"), nil)
if err != nil {
    fmt.Println(err)
}

df.Tail(5)
```
```
       |    Overall rank    Country or region           Score    GDP per capita    Social support    Healthy life expectancy    Freedom to make life choices    Generosity    Perceptions of corruption    
151    |    152             Rwanda                      3.334    0.359             0.711             0.614                      0.555                           0.217         0.411                        
152    |    153             Tanzania                    3.231    0.476             0.885             0.499                      0.417                           0.276         0.147                        
153    |    154             Afghanistan                 3.203    0.35              0.517             0.361                      0                               0.158         0.025                        
154    |    155             Central African Republic    3.083    0.026             0                 0.105                      0.225                           0.235         0.035                        
155    |    156             South Sudan                 2.853    0.306             0.575             0.295                      0.01                            0.202         0.091                        
```