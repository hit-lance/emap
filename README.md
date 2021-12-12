# tiny-map

## Find Closest
- SimpleNodeSet vs KDTree
测benchmark
```shell
go test -benchmem -run=^$ -bench=. -benchtime=1000x
```
KDTree使查询速度提高了500倍
```shell
BenchmarkSimpleNodeSet      1000           4037436 ns/op              48 B/op          1 allocs/op
BenchmarkKDTree             1000              8561 ns/op              48 B/op          1 allocs/op
```