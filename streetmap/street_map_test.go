package streetmap

import (
	nd "etaxi/streetmap/namedict"
	ns "etaxi/streetmap/nodeset"
	"reflect"
	"testing"
)

func TestStreetMap(t *testing.T) {
	fn := "../data/berkeley.osm.xml"
	smKDtreeNaive := NewStreetMapFrom(fn, &ns.KDTree{}, &nd.NaiveNameDict{})
	smNaiveTrie := NewStreetMapFrom(fn, &ns.NaiveNodeSet{}, &nd.Trie{})
	smKDtreeTrie := NewStreetMapFrom(fn, &ns.KDTree{}, &nd.Trie{})

	t.Run("find_closest_naive", func(t *testing.T) {
		got := smNaiveTrie.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		assertNodeIDEqual(t, got, want)
	})

	t.Run("find_closest_kdtree", func(t *testing.T) {
		got := smKDtreeTrie.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		assertNodeIDEqual(t, got, want)
	})

	t.Run("keys_with_prefix_naive", func(t *testing.T) {
		keys := smKDtreeNaive.NameDict.Keys()
		want := 1939
		got := len(keys)
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})

	t.Run("keys_with_prefix_trie", func(t *testing.T) {
		keys := smKDtreeTrie.NameDict.Keys()
		want := 1939
		got := len(keys)
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
		assertSliceEqual(t, smKDtreeTrie.GetNodeIDByName("Starbucks Coffee"), []int64{1467717295, 343610934})
	})
}

func assertNodeIDEqual(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}

func assertSliceEqual(t testing.TB, got, want []int64) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but expected %v", got, want)
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
