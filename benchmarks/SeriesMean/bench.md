# No concurrency

goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMean/459_Points-8                 4469834              2595 ns/op            8184 B/op         10 allocs/op
BenchmarkSeriesMean/90837_Points-8                 13407            880933 ns/op         4101368 B/op         28 allocs/op
BenchmarkSeriesMean/4857378_Points-8                 282          41962914 ns/op        160716027 B/op        44 allocs/op
PASS
ok      github.com/jpoly1219/gambas     66.618s

# Concurrency 1

goos: linux
goarch: amd64
pkg: github.com/jpoly1219/gambas
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
BenchmarkSeriesMean/459_Points-8                 1805366              6823 ns/op            9224 B/op         29 allocs/op
BenchmarkSeriesMean/90837_Points-8                 14415            826585 ns/op         4102408 B/op         47 allocs/op
BenchmarkSeriesMean/4857378_Points-8                 298          39537941 ns/op        160717081 B/op        63 allocs/op
PASS
ok      github.com/jpoly1219/gambas     68.401s

# Concurrency 2