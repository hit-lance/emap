package etaxi

import (
	"testing"
)

func TestTrie(t *testing.T) {

	t.Run("simple test", func(t *testing.T) {
		trie := &Trie{}
		strs := []string{"a", "awls", "sad", "sam", "same", "sap"}
		for i, s := range strs {
			trie.put(s, int64(i))
		}
		
		for i, s := range strs {
			assertNodeIdEqual(t, int64(i), trie.get(s))
		}
		assertNodeIdEqual(t, -1, trie.get("aw"))
		assertNodeIdEqual(t, -1, trie.get("awlsa"))
	})

}

func assertNodeIdEqual(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but expected %d", got, want)
	}
}

