`go test -run=^$ -bench=SeriesMean -benchmem -benchtime=20s -timeout=1h`

# No concurrency
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMean/1_Points-8          877157317               26.78 ns/op            8 B/op          1 allocs/op
BenchmarkSeriesMean/10_Points-8         140136961              170.9 ns/op           248 B/op          5 allocs/op
BenchmarkSeriesMean/100_Points-8        30074172               666.8 ns/op          2040 B/op          8 allocs/op
BenchmarkSeriesMean/1000_Points-8        3953356              5972 ns/op           25208 B/op         12 allocs/op
BenchmarkSeriesMean/10000_Points-8        291363             74675 ns/op          357624 B/op         19 allocs/op
BenchmarkSeriesMean/100000_Points-8        28250            842610 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesMean/1000000_Points-8        2692           8552740 ns/op        41678072 B/op         38 allocs/op
BenchmarkSeriesMean/10000000_Points-8        236         100084086 ns/op        492000507 B/op        49 allocs/op
PASS
ok      github.com/jpoly1219/gambas     242.262s

# Concurrency 1 (Divide and conquer, no channels)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMean/1_Points-8          37734870               624.1 ns/op           376 B/op          6 allocs/op
BenchmarkSeriesMean/10_Points-8         12429267              1836 ns/op            1000 B/op         18 allocs/op
BenchmarkSeriesMean/100_Points-8         7387249              3200 ns/op            3080 B/op         27 allocs/op
BenchmarkSeriesMean/1000_Points-8        2465499              9624 ns/op           26248 B/op         31 allocs/op
BenchmarkSeriesMean/10000_Points-8        283830             76012 ns/op          358664 B/op         38 allocs/op
BenchmarkSeriesMean/100000_Points-8        29038            823423 ns/op         4102408 B/op         47 allocs/op
BenchmarkSeriesMean/1000000_Points-8        2778           8061964 ns/op        41679114 B/op         57 allocs/op
BenchmarkSeriesMean/10000000_Points-8        253          94242493 ns/op        492001562 B/op        68 allocs/op
PASS
ok      github.com/jpoly1219/gambas     231.312s

# Concurrency 2 (Divide and conquer, using channels)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMean/1_Points-8          47568241               497.6 ns/op           248 B/op          4 allocs/op
BenchmarkSeriesMean/10_Points-8         11205786              2115 ns/op             680 B/op         12 allocs/op
BenchmarkSeriesMean/100_Points-8         6173605              3800 ns/op            2616 B/op         18 allocs/op
BenchmarkSeriesMean/1000_Points-8        2403882              9935 ns/op           25784 B/op         22 allocs/op
BenchmarkSeriesMean/10000_Points-8        271948             77136 ns/op          358200 B/op         29 allocs/op
BenchmarkSeriesMean/100000_Points-8        28719            826327 ns/op         4101944 B/op         38 allocs/op
BenchmarkSeriesMean/1000000_Points-8        2766           8039915 ns/op        41678651 B/op         48 allocs/op
BenchmarkSeriesMean/10000000_Points-8        258          93566885 ns/op        492001106 B/op        59 allocs/op
PASS
ok      github.com/jpoly1219/gambas     233.017s