`go test -run=^$ -bench=SeriesMin -benchmem -benchtime=20s -timeout=1h`

# Original (using sort.Float64s)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMin/1_Points-8           24348496               932.9 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-8           5230076              4549 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMin/100_Points-8          2328558             10265 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMin/1000_Points-8          594964             39775 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMin/10000_Points-8          90112            254387 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMin/100000_Points-8          8869           2449231 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMin/1000000_Points-8         1051          21957241 ns/op        83357622 B/op        102 allocs/op
BenchmarkSeriesMin/10000000_Points-8          90         251021952 ns/op        984002526 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     216.958s

# New (using quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMin/1_Points-8           25697004               930.8 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-8           5212107              4556 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMin/100_Points-8          2347297             10159 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMin/1000_Points-8          590694             39841 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMin/10000_Points-8          92462            253447 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMin/100000_Points-8          9022           2456678 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMin/1000000_Points-8         1048          21895992 ns/op        83357619 B/op        102 allocs/op
BenchmarkSeriesMin/10000000_Points-8          88         253558259 ns/op        984002533 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     218.184s

# New (using for loop over data)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMin/1_Points-8           25296696               970.0 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-8           4523696              4610 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMin/100_Points-8          2350720             10171 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMin/1000_Points-8          605127             39624 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMin/10000_Points-8          93234            251407 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMin/100000_Points-8          9052           2440665 ns/op         8204208 B/op         82 allocs/op
BenchmarkSeriesMin/1000000_Points-8         1023          22001633 ns/op        83357619 B/op        102 allocs/op
BenchmarkSeriesMin/10000000_Points-8          87         255599807 ns/op        984002509 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     216.875s

# Fit
