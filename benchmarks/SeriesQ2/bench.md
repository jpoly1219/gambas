`go test -run=^$ -bench=SeriesQ2 -benchmem -benchtime=20s -timeout=1h`

# Original (old median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesQ2/1_Points-12           334727898               69.86 ns/op           32 B/op          2 allocs/op
BenchmarkSeriesQ2/10_Points-12          52953936               420.3 ns/op           272 B/op          6 allocs/op
BenchmarkSeriesQ2/100_Points-12          6036704              3877 ns/op            2064 B/op          9 allocs/op
BenchmarkSeriesQ2/1000_Points-12          257306             88058 ns/op           25232 B/op         13 allocs/op
BenchmarkSeriesQ2/10000_Points-12          20095           1184451 ns/op          357648 B/op         20 allocs/op
BenchmarkSeriesQ2/100000_Points-12          1612          14696094 ns/op         4101392 B/op         29 allocs/op
BenchmarkSeriesQ2/1000000_Points-12          136         173306347 ns/op        41678096 B/op         39 allocs/op
BenchmarkSeriesQ2/10000000_Points-12          10        2076909050 ns/op        492000547 B/op        50 allocs/op
PASS
ok      github.com/jpoly1219/gambas     241.523s

# New (new median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesQ2/1_Points-12           619558501               37.18 ns/op            8 B/op          1 allocs/op
BenchmarkSeriesQ2/10_Points-12          82552210               260.1 ns/op           248 B/op          5 allocs/op
BenchmarkSeriesQ2/100_Points-12         16797900              1385 ns/op            2040 B/op          8 allocs/op
BenchmarkSeriesQ2/1000_Points-12         2107887             11350 ns/op           25208 B/op         12 allocs/op
BenchmarkSeriesQ2/10000_Points-12         117078            190936 ns/op          357624 B/op         19 allocs/op
BenchmarkSeriesQ2/100000_Points-12         10000           2185847 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesQ2/1000000_Points-12          928          24445240 ns/op        41678072 B/op         38 allocs/op
BenchmarkSeriesQ2/10000000_Points-12          96         267341501 ns/op        492000508 B/op        49 allocs/op
PASS
ok      github.com/jpoly1219/gambas     218.721s


*old*

# Original (old median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesQ2/1_Points-8            361497385               65.79 ns/op           32 B/op          2 allocs/op
BenchmarkSeriesQ2/10_Points-8           76772324               305.9 ns/op           272 B/op          6 allocs/op
BenchmarkSeriesQ2/100_Points-8           7529334              3115 ns/op            2064 B/op          9 allocs/op
BenchmarkSeriesQ2/1000_Points-8           325093             73682 ns/op           25232 B/op         13 allocs/op
BenchmarkSeriesQ2/10000_Points-8           23610           1012182 ns/op          357648 B/op         20 allocs/op
BenchmarkSeriesQ2/100000_Points-8           1892          12576280 ns/op         4101392 B/op         29 allocs/op
BenchmarkSeriesQ2/1000000_Points-8           157         150244940 ns/op        41678096 B/op         39 allocs/op
BenchmarkSeriesQ2/10000000_Points-8           13        1766600801 ns/op        492000529 B/op        50 allocs/op
PASS
ok      github.com/jpoly1219/gambas     239.737s

# New (new median)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesQ2/1_Points-8            718953512               32.90 ns/op            8 B/op          1 allocs/op
BenchmarkSeriesQ2/10_Points-8           93087078               257.6 ns/op           248 B/op          5 allocs/op
BenchmarkSeriesQ2/100_Points-8          21726249              1042 ns/op            2040 B/op          8 allocs/op
BenchmarkSeriesQ2/1000_Points-8          2156119             11005 ns/op           25208 B/op         12 allocs/op
BenchmarkSeriesQ2/10000_Points-8          150321            154636 ns/op          357624 B/op         19 allocs/op
BenchmarkSeriesQ2/100000_Points-8          13902           1723354 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesQ2/1000000_Points-8          1160          20179472 ns/op        41678072 B/op         38 allocs/op
BenchmarkSeriesQ2/10000000_Points-8          100         211940489 ns/op        492000504 B/op        49 allocs/op
PASS
ok      github.com/jpoly1219/gambas     234.073s

# New (new median + quickselect)


# Fit
