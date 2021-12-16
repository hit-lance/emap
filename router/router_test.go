package router

import (
	"container/list"
	sm "etaxi/streetmap"
	"fmt"
	"os"
	"testing"
)

func TestRouter(t *testing.T) {
	tinyOsmPath := "../data/tiny-clean.osm.xml"
	berkeleyOsmPath := "../data/berkeley.osm.xml"

	paramsPath := "../data/path_params.txt"
	resultsPath := "../data/path_results.txt"
	testParas := paramsFromFile(paramsPath)
	expectedResults := resultsFromFile(resultsPath)
	testsNum := 8

	if len(testParas) != testsNum || len(expectedResults) != testsNum {
		fmt.Fprintln(os.Stderr, "Failed to read files")
		os.Exit(1)
	}

	t.Run("dijkstra", func(t *testing.T) {
		m := sm.NewStreetMapFrom(tinyOsmPath, &sm.KDTree{}, &sm.Trie{})
		r := Router{}
		sol := r.ShortestPath(dijkstra, m, 38.1, 0.4, 38.6, 0.4)
		want := list.New()
		want.PushBack(int64(41))
		want.PushBack(int64(63))
		want.PushBack(int64(66))
		want.PushBack(int64(46))

		assertListEqual(t, sol, want)
	})

	t.Run("dijkstra_large_scale", func(t *testing.T) {
		m := sm.NewStreetMapFrom(berkeleyOsmPath, &sm.KDTree{}, &sm.Trie{})
		r := Router{}
		for i := 0; i < testsNum; i++ {
			sol := r.ShortestPath(dijkstra, m, testParas[i][0], testParas[i][1], testParas[i][2], testParas[i][3])
			want := expectedResults[i]
			assertListEqual(t, sol, want)
		}
	})

	t.Run("astar", func(t *testing.T) {
		m := sm.NewStreetMapFrom(tinyOsmPath, &sm.KDTree{}, &sm.Trie{})
		r := Router{}
		sol := r.ShortestPath(aStar, m, 38.1, 0.4, 38.6, 0.4)

		want := list.New()
		want.PushBack(int64(41))
		want.PushBack(int64(63))
		want.PushBack(int64(66))
		want.PushBack(int64(46))

		assertListEqual(t, sol, want)
	})

	t.Run("astar_large_scale", func(t *testing.T) {
		m := sm.NewStreetMapFrom(berkeleyOsmPath, &sm.KDTree{}, &sm.Trie{})
		r := Router{}
		for i := 0; i < testsNum; i++ {
			sol := r.ShortestPath(aStar, m, testParas[i][0], testParas[i][1], testParas[i][2], testParas[i][3])
			want := expectedResults[i]
			assertListEqual(t, sol, want)
		}
	})

}

func assertListEqual(t testing.TB, got *list.List, want *list.List) {
	t.Helper()
	g, w := got.Front(), want.Front()
	for ; g != nil && w != nil; g, w = g.Next(), w.Next() {
		if g.Value != w.Value {
			t.Errorf("assertListEqual failed")
			return
		}
	}
	if g != nil || w != nil {
		t.Errorf("assertListEqual failed")
	}
}

func paramsFromFile(paramsPath string) (params [][4]float64) {
	f, _ := os.Open(paramsPath)
	var param [4]float64
	for {
		n, err := fmt.Fscan(f, &param[1], &param[0], &param[3], &param[2])
		if n == 0 || err != nil {
			break
		}
		params = append(params, param)
	}
	return
}

func resultsFromFile(resultsPath string) (results []*list.List) {
	f, _ := os.Open(resultsPath)
	var num int
	var path *list.List
	for {
		n, err := fmt.Fscanln(f, &num)
		if n == 0 || err != nil {
			break
		}
		path = list.New()
		var p int64
		for i := 0; i < num; i++ {
			n, err := fmt.Fscanln(f, &p)
			if n == 0 || err != nil {
				fmt.Fprintln(os.Stderr, "Failed to read results file")
				os.Exit(1)
			}
			path.PushBack(p)
		}
		results = append(results, path)
	}
	return
}

func BenchmarkDijkstra(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	m := sm.NewStreetMapFrom(fn, &sm.KDTree{}, &sm.Trie{})
	r := Router{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.ShortestPath(dijkstra, m, 37.87383979834944, -122.23354274523257, 37.86020837234193, -122.23307272570244)
	}
}

func BenchmarkAStar(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	m := sm.NewStreetMapFrom(fn, &sm.KDTree{}, &sm.Trie{})
	r := Router{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.ShortestPath(aStar, m, 37.87383979834944, -122.23354274523257, 37.86020837234193, -122.23307272570244)
	}
}
