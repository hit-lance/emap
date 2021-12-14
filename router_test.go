package tinymap

import (
	"container/list"
	"fmt"
	"os"
	"testing"
)

func TestRouter(t *testing.T) {
	tiny_osm_path := "./data/tiny-clean.osm.xml"
	berkeley_osm_path := "./data/berkeley.osm.xml"
	params_path := "./data/path_params.txt"
	results_path := "./data/path_results.txt"

	t.Run("dijkstra", func(t *testing.T) {
		sm := NewStreetMapFrom(tiny_osm_path, &KDTree{})
		r := Router{}
		sol := r.ShortestPath(dijkstra, sm, 38.1, 0.4, 38.6, 0.4)
		want := list.New()
		want.PushBack(int64(41))
		want.PushBack(int64(63))
		want.PushBack(int64(66))
		want.PushBack(int64(46))

		assertEqual(t, sol, want)
	})

	t.Run("dijkstra_large_scale", func(t *testing.T) {
		sm := NewStreetMapFrom(berkeley_osm_path, &KDTree{})
		r := Router{}

		testParas := paramsFromFile(params_path)
		expectedResults := resultsFromFile(results_path)
		testsNum := 8

		if len(testParas) != testsNum || len(expectedResults) != testsNum {
			fmt.Fprintln(os.Stderr, "Failed to read files")
			os.Exit(1)
		}

		for i := 0; i < testsNum; i++ {
			sol := r.ShortestPath(dijkstra, sm, testParas[i][0], testParas[i][1], testParas[i][2], testParas[i][3])
			want := expectedResults[i]
			assertEqual(t, sol, want)
		}
	})

}

func assertEqual(t testing.TB, got *list.List, want *list.List) {
	t.Helper()
	g, w := got.Front(), want.Front()
	for ; g != nil && w != nil; g, w = g.Next(), w.Next() {
		if g.Value != w.Value {
			t.Errorf("assertEqual failed")
			return
		}
	}
	if g != nil || w != nil {
		t.Errorf("assertEqual failed")
	}
}

func paramsFromFile(params_path string) (params [][4]float64) {
	f, _ := os.Open(params_path)
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

func resultsFromFile(results_path string) (results []*list.List) {
	f, _ := os.Open(results_path)
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
