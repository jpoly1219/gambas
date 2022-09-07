`go test -run=^$ -bench=SeriesMax -benchmem -benchtime=20s -timeout=1h`

# Original (using sort.Float64s)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMax/1_Points-12          10226950              2129 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-12           728778             32066 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMax/100_Points-12          654548             37748 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMax/1000_Points-12         249414             94159 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMax/10000_Points-12         63768            371655 ns/op          717360 B/op         76 allocs/op
BenchmarkSeriesMax/100000_Points-12         7306           2902158 ns/op         8204849 B/op         94 allocs/op
BenchmarkSeriesMax/1000000_Points-12         938          24541747 ns/op        83358265 B/op        114 allocs/op
BenchmarkSeriesMax/10000000_Points-12         82         287263196 ns/op        984003137 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     225.430s

# New (using quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMax/1_Points-12          11400030              2061 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-12           719157             31761 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMax/100_Points-12          628326             37409 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMax/1000_Points-12         238548             93943 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMax/10000_Points-12         63762            374192 ns/op          717361 B/op         76 allocs/op
BenchmarkSeriesMax/100000_Points-12         7405           2907117 ns/op         8204850 B/op         94 allocs/op
BenchmarkSeriesMax/1000000_Points-12         918          24738828 ns/op        83358265 B/op        114 allocs/op
BenchmarkSeriesMax/10000000_Points-12         73         294510871 ns/op        984003147 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     204.229s

# New (using for loop over data, i, data)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMax/1_Points-12          11323375              2087 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-12           696939             32286 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMax/100_Points-12          639037             37299 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMax/1000_Points-12         253035             95058 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMax/10000_Points-12         63446            369802 ns/op          717361 B/op         76 allocs/op
BenchmarkSeriesMax/100000_Points-12         7399           2925740 ns/op         8204849 B/op         94 allocs/op
BenchmarkSeriesMax/1000000_Points-12         937          25084516 ns/op        83358259 B/op        114 allocs/op
BenchmarkSeriesMax/10000000_Points-12         69         298242797 ns/op        984003132 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     205.204s

# New (using for loop over data, i only)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMax/1_Points-12          10567780              2080 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMax/10_Points-12           750111             32183 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMax/100_Points-12          628203             37620 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMax/1000_Points-12         242797             95403 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMax/10000_Points-12         64016            373840 ns/op          717360 B/op         76 allocs/op
BenchmarkSeriesMax/100000_Points-12         7443           2944097 ns/op         8204849 B/op         94 allocs/op
BenchmarkSeriesMax/1000000_Points-12         897          25070239 ns/op        83358267 B/op        114 allocs/op
BenchmarkSeriesMax/10000000_Points-12         73         292397659 ns/op        984003158 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     205.068s

# Fit



*old*

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
