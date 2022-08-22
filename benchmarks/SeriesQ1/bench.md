`go test -run=^$ -bench=SeriesQ1 -benchmem -benchtime=20s -timeout=1h`

# Original (old median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesQ1/1_Points-8            126074221              188.4 ns/op           104 B/op          5 allocs/op
BenchmarkSeriesQ1/10_Points-8           66445503               354.4 ns/op           296 B/op          7 allocs/op
BenchmarkSeriesQ1/100_Points-8           6277894              3785 ns/op            2088 B/op         10 allocs/op
BenchmarkSeriesQ1/1000_Points-8           270290             88845 ns/op           25256 B/op         14 allocs/op
BenchmarkSeriesQ1/10000_Points-8           20148           1189744 ns/op          357672 B/op         21 allocs/op
BenchmarkSeriesQ1/100000_Points-8           1647          14546270 ns/op         4101416 B/op         30 allocs/op
BenchmarkSeriesQ1/1000000_Points-8           136         173886980 ns/op        41678120 B/op         40 allocs/op
BenchmarkSeriesQ1/10000000_Points-8           10        2052660890 ns/op        492000552 B/op        51 allocs/op
PASS
ok      github.com/jpoly1219/gambas     255.472s

# New (new median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesQ1/1_Points-8            127461896              187.0 ns/op           104 B/op          5 allocs/op
BenchmarkSeriesQ1/10_Points-8           63334408               370.2 ns/op           296 B/op          7 allocs/op
BenchmarkSeriesQ1/100_Points-8           6408037              3705 ns/op            2088 B/op         10 allocs/op
BenchmarkSeriesQ1/1000_Points-8           269560             89125 ns/op           25256 B/op         14 allocs/op
BenchmarkSeriesQ1/10000_Points-8           20246           1178682 ns/op          357672 B/op         21 allocs/op
BenchmarkSeriesQ1/100000_Points-8           1638          14584931 ns/op         4101416 B/op         30 allocs/op
BenchmarkSeriesQ1/1000000_Points-8           138         172660367 ns/op        41678120 B/op         40 allocs/op
BenchmarkSeriesQ1/10000000_Points-8           10        2048735084 ns/op        492000552 B/op        51 allocs/op
PASS
ok      github.com/jpoly1219/gambas     255.076s

# New (new median + quickselect)


# Fit
