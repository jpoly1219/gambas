# Printing

`gambas` provides several options to print your `Series`.

## Print

```go
func (s *Series) Print()
```

`Print` prints all data in a `Series` object.

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

mySeries.Print()
```
```
     |    Fruit     
0    |    apple     
1    |    banana    
2    |    cherry 
```

## PrintRange

```go
func (s *Series) PrintRange(start, end int)
```

`PrintRange` prints data in a `Series` object at a given range.

Index starts at 0.

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

mySeries.PrintRange(1, 3)
```
```
     |    Fruit     
1    |    banana    
2    |    cherry
```

## Head

```go
func (s *Series) Head(howMany int)
```

`Head` prints the first `howMany` items in a `Series` object.

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

mySeries.Head(1)
```
```
     |    Fruit     
0    |    apple  
```

## Tail

```go
func (s *Series) Tail(howMany int)
```

`Tail` prints the last `howMany` items in a `Series` object.

```go
myData := []interface{}{"apple", "banana", "cherry"}
myName := "Fruit"

mySeries, err := gambas.NewSeries(myData, myName, nil)
if err != nil {
    fmt.Println(err)
}

mySeries.Tail(1)
```
```
     |    Fruit     
2    |    cherry
```