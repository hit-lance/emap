package tinymap

import (
	"testing"
)

func TestStreetMap(t *testing.T) {
	fn := "./data/berkeley.osm.xml"
	sm1 := NewStreetMapFrom(fn, &SimpleNodeSet{})
	sm2 := NewStreetMapFrom(fn, &KDTree{})

	t.Run("find_closest_brute_force", func(t *testing.T) {
		got := sm1.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		assertClosest(t, got, want)
	})

	t.Run("find_closest_kdtree)", func(t *testing.T) {
		got := sm2.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		assertClosest(t, got, want)
	})
}

func assertClosest(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}

func BenchmarkSimpleNodeSet(b *testing.B) {
	fn := "./data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &SimpleNodeSet{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Closest(37.875613, -122.26009)
	}
}

func BenchmarkKDTree(b *testing.B) {
	fn := "./data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &KDTree{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Closest(37.875613, -122.26009)
	}
}
