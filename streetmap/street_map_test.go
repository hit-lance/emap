package streetmap

import (
	"testing"
)

func TestStreetMap(t *testing.T) {
	fn := "../data/berkeley.osm.xml"
	sm_kdtree_naive := NewStreetMapFrom(fn, &KDTree{}, &NaiveNameDict{})
	sm_naive_trie := NewStreetMapFrom(fn, &NaiveNodeSet{}, &Trie{})
	sm_kdtree_trie := NewStreetMapFrom(fn, &KDTree{}, &Trie{})

	t.Run("find_closest_naive", func(t *testing.T) {
		got := sm_naive_trie.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		assertClosest(t, got, want)
	})

	t.Run("find_closest_kdtree", func(t *testing.T) {
		got := sm_kdtree_trie.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		assertClosest(t, got, want)
	})

	t.Run("keys_with_prefix_naive", func(t *testing.T) {
		keys := sm_kdtree_naive.nd.keys()
		want := 1939
		got := len(keys)
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})

	t.Run("keys_with_prefix_trie", func(t *testing.T) {
		keys := sm_kdtree_trie.nd.keys()
		want := 1939
		got := len(keys)
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})
}

func assertClosest(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}

func BenchmarkNaiveNodeSet(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &NaiveNodeSet{}, &Trie{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Closest(37.875613, -122.26009)
	}
}

func BenchmarkKDTree(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &KDTree{}, &Trie{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Closest(37.875613, -122.26009)
	}
}

func BenchmarkNaiveNameDict(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &NaiveNodeSet{}, &NaiveNameDict{})
	keys := sm.nd.keys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			sm.getNodesByPrefix(k)
		}
	}
}

func BenchmarkTrie(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &NaiveNodeSet{}, &Trie{})
	keys := sm.nd.keys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			sm.getNodesByPrefix(k)
		}
	}
}
