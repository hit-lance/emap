package etaxi

type trieNode struct {
	val  int64
	next map[rune]*trieNode
}

func NewTrieNode() *trieNode {
	var n trieNode
	n.val = INVALID_NODE_ID
	n.next = make(map[rune]*trieNode)
	return &n
}

type Trie struct {
	root *trieNode
}

func (t *Trie) put(s string, v int64) {
	t.root = putHelper(t.root, []rune(s), v, 0)
}

func putHelper(t *trieNode, r []rune, v int64, d int) *trieNode {
	if t == nil {
		t = NewTrieNode()
	}
	if len(r) == d {
		t.val = v
	} else {
		t.next[r[d]] = putHelper(t.next[r[d]], r, v, d+1)
	}
	return t
}

func (t *Trie) get(s string) (v int64) {
	n := getHelper(t.root, []rune(s), 0)
	if n == nil {
		return INVALID_NODE_ID
	}
	return n.val
}

func getHelper(t *trieNode, r []rune, d int) *trieNode {
	if t == nil || len(r) == d {
		return t
	}

	return getHelper(t.next[r[d]], r, d+1)
}

func (t *Trie) keys() []string {
	return t.keysWithPrefix("")
}

func (t *Trie) keysWithPrefix(pre string) []string {
	s := []string{}
	collect(getHelper(t.root, []rune(pre), 0), pre, &s)
	return s
}

func collect(t *trieNode, pre string, s *[]string) {
	if t == nil {
		return
	}
	if t.val != INVALID_NODE_ID {
		*s = append(*s, pre)
	}
	for r, n := range t.next {
		collect(n, pre+string(r), s)
	}
}
