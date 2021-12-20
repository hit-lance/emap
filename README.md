# etaxi

## Quickstart

### 启动server
```shell
cd server
go build
./server
```

### 测试

- 根据name查询点的id（前缀查询）
```shell
curl --request GET 'localhost:9000/locations?name=天坛北'
```
```json
{
    "node_ids": [
        734521002,
        2383146360
    ]
}
```

- 根据id查询点的具体信息
```shell
curl --request GET 'localhost:9000/locations/734521002'
```
```json
{
    "node_ids": [
        734521002,
        2383146360
    ]
}
```

- 导航
```shell
curl --request GET 'localhost:9000/direction?slat=39.9322003&slon=116.3978560&dlat=39.8868562&dlon=116.4046622'```
```json
{
    "nodes": [
        [
            39.9320951,
            116.3978921
        ],
        [
            39.9321088,
            116.3983186
        ],
        ...
        [
            39.8870435,
            116.3978756
        ]
    ],
    "text": "全程5.849公里\n从起点出发，进入*地安门东大街*，继续前行0.141公里\n右转，进入*北河沿大街*，继续前行2.026公里\n直行，进入*南河沿大街*，继续前行0.811公里\n右转，继续前行0.022公里\n左转，进入*正义路*，继续前行0.780公里\n向右前方行驶，继续前行0.012公里\n直行，继续前行0.015公里\n向左前方行驶，继续前行0.046公里\n直行，继续前行0.064公里\n向右前方行驶，进入*北翔凤胡同*，继续前行0.250公里\n向左前方行驶，进入*草厂三条*，继续前行0.532公里\n右转，进入*珠市口东大街*，继续前行0.390公里\n左转，继续前行0.016公里\n右转，进入*珠市口东大街*，继续前行0.034公里\n左转，继续前行0.011公里\n向右前方行驶，继续前行0.010公里\n向左前方行驶，继续前行0.043公里\n直行，继续前行0.076公里\n直行，继续前行0.020公里\n直行，继续前行0.028公里\n左转，继续前行0.189公里\n右转，继续前行0.073公里\n左转，继续前行0.085公里\n右转，继续前行0.111公里\n直行，继续前行0.009公里\n右转，进入*粉厂胡同*，继续前行0.053公里\n到达目的地\n"
}
```

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
