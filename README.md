# tiny-map

## Find Closet
- SimpleNodeSet vs KDTree
### benchmark
```shell
go test -benchmem -run=^$ -bench=. -benchtime=1000x
```
### result
```shell
BenchmarkSimpleNodeSet      1000           4037436 ns/op              48 B/op          1 allocs/op
BenchmarkKDTree             1000              8561 ns/op              48 B/op          1 allocs/op
```