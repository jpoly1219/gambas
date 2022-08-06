# GroupBy

Here are the operations that you can apply on `GroupBy` objects.

The data used in the example `neo_v2.csv` is NASA's list of Nearest Earth Objects, sourced from [Kaggle](https://www.kaggle.com/datasets/sameepvani/nasa-nearest-earth-objects).

## Agg

```go
func (gb *GroupBy) Agg(targetCol []string, aggFunc StatsFunc) (DataFrame, error)
```

`Agg` aggregates data in the `GroupBy` object using the given `aggFunc`.

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