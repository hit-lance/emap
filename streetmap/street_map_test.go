package streetmap

import (
	nd "etaxi/streetmap/namedict"
	ns "etaxi/streetmap/nodeset"
	"testing"
)

func TestStreetMap(t *testing.T) {
	fn := "../data/berkeley.osm.xml"
	sm_kdtree_naive := NewStreetMapFrom(fn, &ns.KDTree{}, &nd.NaiveNameDict{})
	sm_naive_trie := NewStreetMapFrom(fn, &ns.NaiveNodeSet{}, &nd.Trie{})
	sm_kdtree_trie := NewStreetMapFrom(fn, &ns.KDTree{}, &nd.Trie{})

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
		keys := sm_kdtree_naive.NameDict.Keys()
		want := 1939
		got := len(keys)
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})

	t.Run("keys_with_prefix_trie", func(t *testing.T) {
		keys := sm_kdtree_trie.NameDict.Keys()
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
	sm := NewStreetMapFrom(fn, &ns.NaiveNodeSet{}, &nd.Trie{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Closest(37.875613, -122.26009)
	}
}

func BenchmarkKDTree(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &ns.KDTree{}, &nd.Trie{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Closest(37.875613, -122.26009)
	}
}

func BenchmarkNaiveNameDict(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &ns.NaiveNodeSet{}, &nd.NaiveNameDict{})
	keys := sm.NameDict.Keys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			sm.GetNodesByPrefix(key)
		}
	}
}

func BenchmarkTrie(b *testing.B) {
	fn := "../data/berkeley.osm.xml"
	sm := NewStreetMapFrom(fn, &ns.NaiveNodeSet{}, &nd.Trie{})
	keys := sm.NameDict.Keys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			sm.GetNodesByPrefix(key)
		}      
	}
}
