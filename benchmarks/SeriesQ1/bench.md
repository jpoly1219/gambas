`go test -run=^$ -bench=SeriesQ1 -benchmem -benchtime=20s -timeout=1h`

`go test -run=^$ -bench=SeriesQ1 -benchmem -benchtime=20s -timeout=1h`

# Original (old median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesQ1/1_Points-12           100000000              203.4 ns/op           104 B/op          5 allocs/op
BenchmarkSeriesQ1/10_Points-12          49656849               459.5 ns/op           296 B/op          7 allocs/op
BenchmarkSeriesQ1/100_Points-12          5538144              4324 ns/op            2088 B/op         10 allocs/op
BenchmarkSeriesQ1/1000_Points-12          262766             89911 ns/op           25256 B/op         14 allocs/op
BenchmarkSeriesQ1/10000_Points-12          19815           1205243 ns/op          357672 B/op         21 allocs/op
BenchmarkSeriesQ1/100000_Points-12          1614          14744908 ns/op         4101416 B/op         30 allocs/op
BenchmarkSeriesQ1/1000000_Points-12          134         176018326 ns/op        41678120 B/op         40 allocs/op
BenchmarkSeriesQ1/10000000_Points-12          10        2091866430 ns/op        492000552 B/op        51 allocs/op
PASS
ok      github.com/jpoly1219/gambas     234.016s

# New (quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesQ1/4_Points-12           219200240              106.5 ns/op            56 B/op          3 allocs/op
BenchmarkSeriesQ1/10_Points-12          100000000              213.6 ns/op           248 B/op          5 allocs/op
BenchmarkSeriesQ1/100_Points-12         21532389              1023 ns/op            2040 B/op          8 allocs/op
BenchmarkSeriesQ1/1000_Points-12         2748222              8756 ns/op           25208 B/op         12 allocs/op
BenchmarkSeriesQ1/10000_Points-12         154868            147325 ns/op          357624 B/op         19 allocs/op
BenchmarkSeriesQ1/100000_Points-12         15508           1534313 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesQ1/1000000_Points-12         1220          16925232 ns/op        41678072 B/op         38 allocs/op
BenchmarkSeriesQ1/10000000_Points-12         136         178804705 ns/op        492000507 B/op        49 allocs/op
PASS
ok      github.com/jpoly1219/gambas     253.269s

# Fit



*old*

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

# Fit
