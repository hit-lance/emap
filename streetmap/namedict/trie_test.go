package streetmap

import (
	"reflect"
	"testing"
)

func TestTrie(t *testing.T) {

	t.Run("keys_with_prefix", func(t *testing.T) {
		trie := &Trie{}
		strs := []string{"a", "awls", "sad", "sam", "same", "sap"}
		for i, s := range strs {
			trie.Put(s, int64(i))
		}

		assertSliceEqual(t, trie.KeysWithPrefix("sam"), []string{"sam", "same"})
		assertSliceEqual(t, trie.KeysWithPrefix("c"), []string{})
	})

	t.Run("chinese", func(t *testing.T) {
		trie := &Trie{}
		strs := []string{"中", "中国", "中国广西", "中国广东", "中国广东深圳"}
		for i, s := range strs {
			trie.Put(s, int64(i))
		}

		for i, s := range strs {
			assertNodeIDEqual(t, int64(i), trie.Get(s)[0])
		}

		trie.Put("中国", 5)
		assert(t, len(trie.Get("中")) == 1)
		assert(t, len(trie.Get("中国")) == 2)
		assert(t, len(trie.Get("美国")) == 0)
		assert(t, len(trie.Get("中国广东深圳南山")) == 0)

	})

}

func assert(t testing.TB, b bool) {
	if !b {
		t.Errorf("assert failed")
	}
}

func assertNodeIDEqual(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}

func assertSliceEqual(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but expected %v", got, want)
	}
}
