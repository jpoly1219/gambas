`go test -run=^$ -bench=SeriesMin -benchmem -benchtime=20s -timeout=1h`

# Original (using sort.Flaot64s)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMin/1_Points-12          10982055              2080 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-12           727849             32806 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMin/100_Points-12          628476             37267 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMin/1000_Points-12         248702             94780 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMin/10000_Points-12         63840            368031 ns/op          717360 B/op         76 allocs/op
BenchmarkSeriesMin/100000_Points-12         7564           2891917 ns/op         8204849 B/op         94 allocs/op
BenchmarkSeriesMin/1000000_Points-12         925          24553011 ns/op        83358260 B/op        114 allocs/op
BenchmarkSeriesMin/10000000_Points-12         67         300010681 ns/op        984003145 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     203.983s

# New (using quickselect)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMin/1_Points-12          11430456              2084 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-12           692506             32354 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMin/100_Points-12          668431             37555 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMin/1000_Points-12         241129             95014 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMin/10000_Points-12         64989            371507 ns/op          717361 B/op         76 allocs/op
BenchmarkSeriesMin/100000_Points-12         7630           2923864 ns/op         8204850 B/op         94 allocs/op
BenchmarkSeriesMin/1000000_Points-12         930          24822006 ns/op        83358271 B/op        114 allocs/op
BenchmarkSeriesMin/10000000_Points-12         68         302538172 ns/op        984003163 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     206.301s

# New (using for loop, i only)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSeriesMin/1_Points-12          11744286              2065 ns/op             368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-12           696724             32472 ns/op            2288 B/op         42 allocs/op
BenchmarkSeriesMin/100_Points-12          644619             37376 ns/op            6192 B/op         54 allocs/op
BenchmarkSeriesMin/1000_Points-12         244724             94495 ns/op           52528 B/op         62 allocs/op
BenchmarkSeriesMin/10000_Points-12         65012            369962 ns/op          717361 B/op         76 allocs/op
BenchmarkSeriesMin/100000_Points-12         7826           2945804 ns/op         8204849 B/op         94 allocs/op
BenchmarkSeriesMin/1000000_Points-12         930          24940658 ns/op        83358262 B/op        114 allocs/op
BenchmarkSeriesMin/10000000_Points-12         79         308226599 ns/op        984003133 B/op       136 allocs/op
PASS
ok      github.com/jpoly1219/gambas     230.829s


*old*

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

# New (using for loop over data, i, data)
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

# New (using for loop over data, i only)
goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMin/1_Points-8           25311508               933.1 ns/op           368 B/op          7 allocs/op
BenchmarkSeriesMin/10_Points-8           5208818              4564 ns/op            1488 B/op         27 allocs/op
BenchmarkSeriesMin/100_Points-8          2348414             10166 ns/op            5552 B/op         42 allocs/op
BenchmarkSeriesMin/1000_Points-8          583315             39729 ns/op           51888 B/op         50 allocs/op
BenchmarkSeriesMin/10000_Points-8          92572            253265 ns/op          716720 B/op         64 allocs/op
BenchmarkSeriesMin/100000_Points-8          9838           2424566 ns/op         8204209 B/op         82 allocs/op
BenchmarkSeriesMin/1000000_Points-8         1065          21801945 ns/op        83357620 B/op        102 allocs/op
BenchmarkSeriesMin/10000000_Points-8          92         242466478 ns/op        984002533 B/op       124 allocs/op
PASS
ok      github.com/jpoly1219/gambas     237.016s

# Fit
