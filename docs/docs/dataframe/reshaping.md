# Reshaping

You can reshape a `DataFrame` object using these functions.

The data used in the example `neo_v2.csv` is NASA's list of Nearest Earth Objects, sourced from [Kaggle](https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects).

## Pivot

```go
func (df *DataFrame) Pivot(column, value string) (DataFrame, error)
```

`Pivot` returns an organized `Dataframe` that has values corresponding to the index and the given column.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), []string{"id"})
if err != nil {
     fmt.Println(err)
}
df.Head(5)
fmt.Println("")

res, err := df.Pivot("hazardous", "est_diameter_min")
if err != nil {
     fmt.Println(err)
}

res.Head(5)
```
```
id         |    id         name                   est_diameter_min    est_diameter_max    relative_velocity    miss_distance            orbiting_body    sentry_object    absolute_magnitude    hazardous    
2162635    |    2162635    162635 (2000 SS164)    1.1982708007        2.6794149658        13569.2492241812     5.483974408284605e+07    Earth            false            16.73                 false        
2277475    |    2277475    277475 (2005 WK4)      0.2658              0.5943468684        73588.7266634981     6.143812652395093e+07    Earth            false            20                    true         
2512244    |    2512244    512244 (2015 YE18)     0.7220295577        1.6145071727        114258.6921290512    4.979872494045679e+07    Earth            false            17.83                 false        
3596030    |    3596030    (2012 BV13)            0.096506147         0.2157943048        24764.3031380016     2.543497272075825e+07    Earth            false            22.2                  false        
3667127    |    3667127    (2014 GE35)            0.2550086879        0.5702167609        42737.7337647264     4.627556700130072e+07    Earth            false            20.09                 true         

id         |    false           true            
2162635    |    1.1982708007    NaN             
2277475    |    NaN             0.2658          
2512244    |    0.7220295577    NaN             
3596030    |    0.096506147     NaN             
3667127    |    NaN             0.2550086879    
```

## PivotTable

```go
func (df *DataFrame) PivotTable(index, column, value string, aggFunc StatsFunc) (DataFrame, error)
```
`PivotTable` rearranges the data by a given `index` and `column`.

Each value will be aggregated via an aggregation function. Pick three columns from the `DataFrame`, each to serve as the `index`, `column`, and `value`. 

`PivotTable` ignores `NaN` values.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), nil)
if err != nil {
     fmt.Println(err)
}
df.SortByValues("id", true)
df.Head(12)
fmt.Println("")

res, err := df.PivotTable("id", "hazardous", "miss_distance", gambas.Mean)
if err != nil {
     fmt.Println(err)
}

res.Head(5)
```
```
         |    id         name                      est_diameter_min    est_diameter_max    relative_velocity    miss_distance             orbiting_body    sentry_object    absolute_magnitude    hazardous    
12709    |    2000433    433 Eros (A898 PA)        23.0438466577       51.5276075896       15884.2526231559     5.468807778293672e+07     Earth            false            10.31                 false        
37651    |    2000433    433 Eros (A898 PA)        23.0438466577       51.5276075896       21402.705247412      2.6729521135077037e+07    Earth            false            10.31                 false        
56533    |    2000433    433 Eros (A898 PA)        23.0438466577       51.5276075896       21761.7034264303     3.120591927495648e+07     Earth            false            10.31                 false        
1847     |    2000719    719 Albert (A911 TB)      2.0443487103        4.5713026859        27551.5971939875     4.258288106079324e+07     Earth            false            15.57                 false        
36418    |    2001036    1036 Ganymed (A924 UB)    37.8926498379       84.7305408852       51496.9232928228     5.3721237819369085e+07    Earth            false            9.23                  false        
13527    |    2001566    1566 Icarus (1949 MA)     1.4274305148        3.1918316641        76768.6272477926     5.1882752851231776e+07    Earth            false            16.35                 true         
17077    |    2001566    1566 Icarus (1949 MA)     1.4274305148        3.1918316641        136986.6291056903    4.442794284496872e+07     Earth            false            16.35                 true         
28226    |    2001566    1566 Icarus (1949 MA)     1.4274305148        3.1918316641        120524.2906272869    6.046637755115862e+07     Earth            false            16.35                 true         
68073    |    2001566    1566 Icarus (1949 MA)     1.4274305148        3.1918316641        108801.2963741598    8.053781761441007e+06     Earth            false            16.35                 true         
73540    |    2001566    1566 Icarus (1949 MA)     1.4274305148        3.1918316641        78130.8042822814     6.564713381994684e+07     Earth            false            16.35                 true         
4156     |    2001580    1580 Betulia (1950 KA)    3.0658787593        6.8555133165        105157.7758512475    3.5573935508316e+07       Earth            false            14.69                 false        
67943    |    2001580    1580 Betulia (1950 KA)    3.0658787593        6.8555133165        109184.9019317352    5.268620819957744e+07     Earth            false            14.69                 false        

id         |    false               true                
2000433    |    3.7541172731e+07    NaN                 
2000719    |    4.2582881061e+07    NaN                 
2001036    |    5.3721237819e+07    NaN                 
2001566    |    NaN                 4.6095597766e+07    
2001580    |    4.4130071854e+07    NaN                 
```

## Melt

```go
func (df *DataFrame) Melt(colName, valueName string) (DataFrame, error)
```

`Melt` returns the table from wide to long format.

Use `Melt` to revert to pre-`Pivot` format.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), nil)
if err != nil {
     fmt.Println(err)
}

pivoted, err := df.PivotTable("id", "hazardous", "miss_distance", gambas.Mean)
if err != nil {
     fmt.Println(err)
}

pivoted.Head(5)
fmt.Println("")

melted, err := pivoted.Melt("hazardous", "miss_distance")
if err != nil {
     fmt.Println(err)
}

melted.Head(5)
```
```
id         |    false               true                
2000433    |    3.7541172731e+07    NaN                 
2000719    |    4.2582881061e+07    NaN                 
2001036    |    5.3721237819e+07    NaN                 
2001566    |    NaN                 4.6095597766e+07    
2001580    |    4.4130071854e+07    NaN                 

id         |    id         hazardous    miss_distance       
2000433    |    2000433    false        3.7541172731e+07    
2000433    |    2000433    true         NaN                 
2000719    |    2000719    false        4.2582881061e+07    
2000719    |    2000719    true         NaN                 
2001036    |    2001036    false        5.3721237819e+07
```

## GroupBy

```go
func (df *DataFrame) GroupBy(by ...string) (GroupBy, error)
```

`GroupBy` groups selected columns in a `DataFrame` object and returns a `GroupBy` object.

```go
df, err := gambas.ReadCsv(filepath.Join(".", "neo_v2.csv"), nil)
if err != nil {
     fmt.Println(err)
}

filtered, err := df.LocRows([]interface{}{0}, []interface{}{1}, []interface{}{2}, []interface{}{3}, []interface{}{4})
if err != nil {
     fmt.Println(err)
}

filtered.Print()
fmt.Println("")

gb, err := filtered.GroupBy("hazardous")
if err != nil {
     fmt.Println(err)
}

res, err := gb.Agg([]string{"relative_velocity"}, gambas.Mean)
if err != nil {
     fmt.Println(err)
}

res.Print()
```
```
     |    id         name                   est_diameter_min    est_diameter_max    relative_velocity    miss_distance            orbiting_body    sentry_object    absolute_magnitude    hazardous    
0    |    2162635    162635 (2000 SS164)    1.1982708007        2.6794149658        13569.2492241812     5.483974408284605e+07    Earth            false            16.73                 false        
1    |    2277475    277475 (2005 WK4)      0.2658              0.5943468684        73588.7266634981     6.143812652395093e+07    Earth            false            20                    true         
2    |    2512244    512244 (2015 YE18)     0.7220295577        1.6145071727        114258.6921290512    4.979872494045679e+07    Earth            false            17.83                 false        
3    |    3596030    (2012 BV13)            0.096506147         0.2157943048        24764.3031380016     2.543497272075825e+07    Earth            false            22.2                  false        
4    |    3667127    (2014 GE35)            0.2550086879        0.5702167609        42737.7337647264     4.627556700130072e+07    Earth            false            20.09                 true         

hazardous    |    hazardous    relative_velocity    
false        |    false        50864.081            
true         |    true         58163.23
```