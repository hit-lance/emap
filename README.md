# tiny-map

## 测benchmark
```shell
go test -benchmem -run=^$ -bench=. -benchtime=1000x
```
### Find Closest
- SimpleNodeSet vs KDTree
KDTree使查询速度提高了500倍
```shell
BenchmarkSimpleNodeSet      1000           4194527 ns/op              48 B/op          1 allocs/op
BenchmarkKDTree             1000              7252 ns/op              48 B/op          1 allocs/op
```

### Routing
- Dijkstra vs A*
A*使路径查询速度提高了3倍，相应内存操作也少了许多（因为做了内存优化）
```shell
BenchmarkDijkstra           1000          48377061 ns/op         7890971 B/op      30531 allocs/op
BenchmarkAStar              1000          18315887 ns/op         3350914 B/op       5392 allocs/op
```