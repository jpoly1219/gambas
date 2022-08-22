`go test -run=^$ -bench=SeriesMax -benchmem -benchtime=20s -timeout=1h`

# Original (using sort.Float64s)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMax/1_Points-8           24760891               927.1 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-8           5213476              4539 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMax/100_Points-8          2340549             10196 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMax/1000_Points-8          582154             39314 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMax/10000_Points-8          96447            249320 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMax/100000_Points-8         10027           2394329 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMax/1000000_Points-8         1084          21500424 ns/op        83357622 B/op        102 allocs/op
BenchmarkSeriesMax/10000000_Points-8          84         256520645 ns/op        984002537 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     235.727s

# New (using quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMax/1_Points-8           25782938               926.5 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-8           5245796              4539 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMax/100_Points-8          2362231             10089 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMax/1000_Points-8          583839             39389 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMax/10000_Points-8          95119            249029 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMax/100000_Points-8          9022           2388620 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMax/1000000_Points-8         1088          21329180 ns/op        83357618 B/op        102 allocs/op
BenchmarkSeriesMax/10000000_Points-8         102         245373607 ns/op        984002524 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     240.135s

# New (using for loop over data, i, data)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMax/1_Points-8           24622449               926.5 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-8           5237313              4537 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMax/100_Points-8          2359708             10115 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMax/1000_Points-8          592321             39374 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMax/10000_Points-8          94497            250045 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMax/100000_Points-8          8892           2397657 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMax/1000000_Points-8         1057          21530674 ns/op        83357618 B/op        102 allocs/op
BenchmarkSeriesMax/10000000_Points-8          84         253558876 ns/op        984002494 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     215.008s

# New (using for loop over data, i only)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMax/1_Points-8           25651752               924.6 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-8           5228410              4540 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMax/100_Points-8          2330085             10231 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMax/1000_Points-8          585699             39241 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMax/10000_Points-8          94118            249252 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMax/100000_Points-8         10023           2390242 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMax/1000000_Points-8         1087          21264351 ns/op        83357618 B/op        102 allocs/op
BenchmarkSeriesMax/10000000_Points-8          84         243072013 ns/op        984002524 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     236.006s

# Fit
