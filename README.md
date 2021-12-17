# tiny-map

## 测benchmark

```shell
go test -benchmem -run=^$ -bench=. -benchtime=1000x
```

- cpu: Intel(R) Core(TM) i7-8700K CPU @ 3.70GHz

### Find Closest
#### NaiveNodeSet vs KDTree

KDTree使查询速度提高了500多倍。
```shell
BenchmarkNaiveNodeSet-12            1000           2005871 ns/op              48 B/op          1 allocs/op
BenchmarkKDTree-12                  1000              2734 ns/op              48 B/op          1 allocs/op

```

#### Routing
- Dijkstra vs A*

A*使路径查询速度提高了3倍，内存操作也少了许多（因为做了内存优化）。
```shell
BenchmarkDijkstra-12                1000          13260664 ns/op         8112721 B/op      30536 allocs/op
BenchmarkAStar-12                   1000           5809084 ns/op         3572189 B/op       5394 allocs/op
```

#### Autocomplete
Trie使前缀搜索速度提高了70多倍。
```shell
BenchmarkNaiveNameDict-12           1000          55106339 ns/op           38736 B/op       2048 allocs/op
BenchmarkTrie-12                    1000            723659 ns/op          119336 B/op       4628 allocs/op

```
