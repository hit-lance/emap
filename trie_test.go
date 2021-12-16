package etaxi

import (
	"reflect"
	"testing"
)

func TestTrie(t *testing.T) {

	t.Run("test_keys_with_prefix", func(t *testing.T) {
		trie := &Trie{}
		strs := []string{"a", "awls", "sad", "sam", "same", "sap"}
		for i, s := range strs {
			trie.put(s, int64(i))
		}

		assertSliceEqual(t, trie.keysWithPrefix("sam"), []int64{3, 4})
		assertSliceEqual(t, trie.keysWithPrefix("c"), []int64{})
	})

	t.Run("chinese", func(t *testing.T) {
		trie := &Trie{}
		strs := []string{"中", "中国", "中国广西", "中国广东", "中国广东深圳"}
		for i, s := range strs {
			trie.put(s, int64(i))
		}

		for i, s := range strs {
			assertNodeIdEqual(t, int64(i), trie.get(s))
		}
		assertNodeIdEqual(t, -1, trie.get("美国"))
		assertNodeIdEqual(t, -1, trie.get("中国广东深圳南山"))
	})

}

func assertNodeIdEqual(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}

func assertSliceEqual(t testing.TB, got, want []int64) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d but expected %d", got, want)
	}
}
