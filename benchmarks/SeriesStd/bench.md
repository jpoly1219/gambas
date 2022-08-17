`go test -run=^$ -bench=SeriesStd -benchmem -benchtime=20s -timeout=1h`

# No Concurrency (Mean No Concurrency)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesStd/1_Points-8           309436428               77.26 ns/op           16 B/op          2 allocs/op
BenchmarkSeriesStd/10_Points-8          44632078               533.9 ns/op           496 B/op         10 allocs/op
BenchmarkSeriesStd/100_Points-8          6783548              3472 ns/op            4080 B/op         16 allocs/op
BenchmarkSeriesStd/1000_Points-8          688208             34118 ns/op           50416 B/op         24 allocs/op
BenchmarkSeriesStd/10000_Points-8          66298            358676 ns/op          715248 B/op         38 allocs/op
BenchmarkSeriesStd/100000_Points-8          5871           3712442 ns/op         8202736 B/op         56 allocs/op
BenchmarkSeriesStd/1000000_Points-8          613          37931164 ns/op        83356145 B/op         76 allocs/op
BenchmarkSeriesStd/10000000_Points-8          50         426610185 ns/op        984001011 B/op        98 allocs/op
PASS
ok      github.com/jpoly1219/gambas     216.600s

# No Concurrency (Mean Concurrency 1)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesStd/1_Points-8           33786644               699.9 ns/op           384 B/op          7 allocs/op
BenchmarkSeriesStd/10_Points-8           9040416              2632 ns/op            1248 B/op         23 allocs/op
BenchmarkSeriesStd/100_Points-8          2957648              8123 ns/op            5120 B/op         35 allocs/op
BenchmarkSeriesStd/1000_Points-8          594800             39560 ns/op           51456 B/op         43 allocs/op
BenchmarkSeriesStd/10000_Points-8          62488            378303 ns/op          716288 B/op         57 allocs/op
BenchmarkSeriesStd/100000_Points-8          5515           3803368 ns/op         8203776 B/op         75 allocs/op
BenchmarkSeriesStd/1000000_Points-8          619          38127257 ns/op        83357189 B/op         95 allocs/op
BenchmarkSeriesStd/10000000_Points-8          50         433397237 ns/op        984002128 B/op       117 allocs/op
PASS
ok      github.com/jpoly1219/gambas     215.924s

# No Concurrency (Mean Concurrency 2)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesStd/1_Points-8           42371551               563.0 ns/op           256 B/op          5 allocs/op
BenchmarkSeriesStd/10_Points-8           8544508              2600 ns/op             928 B/op         17 allocs/op
BenchmarkSeriesStd/100_Points-8          3448404              6900 ns/op            4656 B/op         26 allocs/op
BenchmarkSeriesStd/1000_Points-8          599298             39366 ns/op           50992 B/op         34 allocs/op
BenchmarkSeriesStd/10000_Points-8          62605            376788 ns/op          715824 B/op         48 allocs/op
BenchmarkSeriesStd/100000_Points-8          5806           3746202 ns/op         8203312 B/op         66 allocs/op
BenchmarkSeriesStd/1000000_Points-8          616          37686042 ns/op        83356727 B/op         86 allocs/op
BenchmarkSeriesStd/10000000_Points-8          51         427872690 ns/op        984001657 B/op       108 allocs/op
PASS
ok      github.com/jpoly1219/gambas     213.760s

# Concurrency
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesStd/1_Points-8           23112361              1059 ns/op             560 B/op          9 allocs/op
BenchmarkSeriesStd/10_Points-8           4881670              4700 ns/op            1680 B/op         29 allocs/op
BenchmarkSeriesStd/100_Points-8          2313765             10315 ns/op            5744 B/op         44 allocs/op
BenchmarkSeriesStd/1000_Points-8          623052             39673 ns/op           52080 B/op         52 allocs/op
BenchmarkSeriesStd/10000_Points-8          93640            251847 ns/op          716912 B/op         66 allocs/op
BenchmarkSeriesStd/100000_Points-8          9910           2432205 ns/op         8204401 B/op         84 allocs/op
BenchmarkSeriesStd/1000000_Points-8         1059          21948793 ns/op        83357810 B/op        104 allocs/op
BenchmarkSeriesStd/10000000_Points-8          86         244788081 ns/op        984002723 B/op       126 allocs/op
PASS
ok      github.com/jpoly1219/gambas     240.378s

# Fit
