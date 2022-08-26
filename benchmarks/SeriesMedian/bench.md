`go test -run=^$ -bench=SeriesMedian -benchmem -benchtime=20s -timeout=1h`

# Original (using sort.Float64s)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMedian/1_Points-12               358115401               63.83 ns/op           32 B/op          2 allocs/op
BenchmarkSeriesMedian/10_Points-12              66023378               349.5 ns/op           272 B/op          6 allocs/op
BenchmarkSeriesMedian/100_Points-12              6426342              3657 ns/op            2064 B/op          9 allocs/op
BenchmarkSeriesMedian/1000_Points-12              280629             83836 ns/op           25232 B/op         13 allocs/op
BenchmarkSeriesMedian/10000_Points-12              21092           1132475 ns/op          357648 B/op         20 allocs/op
BenchmarkSeriesMedian/100000_Points-12              1695          13969096 ns/op         4101392 B/op         29 allocs/op
BenchmarkSeriesMedian/1000000_Points-12              144         165653059 ns/op        41678096 B/op         39 allocs/op
BenchmarkSeriesMedian/10000000_Points-12              12        1951916083 ns/op        492000544 B/op        50 allocs/op
PASS
ok      github.com/jpoly1219/gambas     241.930s
# New (using quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMedian/1_Points-12               766098285               30.53 ns/op            8 B/op          1 allocs/op
BenchmarkSeriesMedian/10_Points-12              100000000              224.9 ns/op           248 B/op          5 allocs/op
BenchmarkSeriesMedian/100_Points-12             18944134              1227 ns/op            2040 B/op          8 allocs/op
BenchmarkSeriesMedian/1000_Points-12             2154841             11096 ns/op           25208 B/op         12 allocs/op
BenchmarkSeriesMedian/10000_Points-12             131322            174420 ns/op          357624 B/op         19 allocs/op
BenchmarkSeriesMedian/100000_Points-12             10000           2020213 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesMedian/1000000_Points-12             1069          21498679 ns/op        41678073 B/op         38 allocs/op
BenchmarkSeriesMedian/10000000_Points-12             100         232671137 ns/op        492000504 B/op        49 allocs/op
PASS
ok      github.com/jpoly1219/gambas     214.180s

*old*
# Original (using sort.Float64s)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMedian/1_Points-8                380852155               61.79 ns/op           32 B/op          2 allocs/op
BenchmarkSeriesMedian/10_Points-8               70906406               316.1 ns/op           272 B/op          6 allocs/op
BenchmarkSeriesMedian/100_Points-8               7100576              3327 ns/op            2064 B/op          9 allocs/op
BenchmarkSeriesMedian/1000_Points-8               304057             75786 ns/op           25232 B/op         13 allocs/op
BenchmarkSeriesMedian/10000_Points-8               23196           1027268 ns/op          357648 B/op         20 allocs/op
BenchmarkSeriesMedian/100000_Points-8               1845          12739672 ns/op         4101392 B/op         29 allocs/op
BenchmarkSeriesMedian/1000000_Points-8               156         152445541 ns/op        41678096 B/op         39 allocs/op
BenchmarkSeriesMedian/10000000_Points-8               13        1773671247 ns/op        492000528 B/op        50 allocs/op
PASS
ok      github.com/jpoly1219/gambas     237.264s

# New (using quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMedian/1_Points-8                863071578               26.94 ns/op            8 B/op          1 allocs/op
BenchmarkSeriesMedian/10_Points-8               100000000              206.2 ns/op           248 B/op          5 allocs/op
BenchmarkSeriesMedian/100_Points-8              18125022              1267 ns/op            2040 B/op          8 allocs/op
BenchmarkSeriesMedian/1000_Points-8              2199795             10953 ns/op           25208 B/op         12 allocs/op
BenchmarkSeriesMedian/10000_Points-8              124786            181436 ns/op          357624 B/op         19 allocs/op
BenchmarkSeriesMedian/100000_Points-8              12926           1859029 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesMedian/1000000_Points-8              1332          17135331 ns/op        41678072 B/op         38 allocs/op
BenchmarkSeriesMedian/10000000_Points-8              100         208230142 ns/op        492000504 B/op        49 allocs/op
PASS
ok      github.com/jpoly1219/gambas     230.655s

# Fit
