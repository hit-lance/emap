package tinymap

import (
	"container/list"
	"testing"
)

func TestRouter(t *testing.T) {

	t.Run("dijkstra", func(t *testing.T) {
		fn := "./tiny-clean.osm.xml"
		sm := NewStreetMapFrom(fn, &KDTree{})
		r := Router{}
		sol := r.ShortestPath(dijkstra, sm, 38.1, 0.4, 38.6, 0.4)
		want := []int64{41, 63, 66, 46}

		assertEqual(t, sol, want)
	})

}

func assertEqual(t testing.TB, got *list.List, want []int64) {
	t.Helper()
	e, i := got.Front(), 0
	for ; e != nil && i < len(want); e, i = e.Next(), i+1 {
		if e.Value != want[i] {
			t.Errorf("assertEqual failed")
		}
	}
	if e != nil || i < len(want) {
		t.Errorf("assertEqual failed")
	}
}
