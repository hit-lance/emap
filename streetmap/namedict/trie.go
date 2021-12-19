package streetmap

type trieNode struct {
	val  *[]int64
	next map[rune]*trieNode
}

func newTrieNode() *trieNode {
	var n trieNode
	n.val = &[]int64{}
	n.next = make(map[rune]*trieNode)
	return &n
}

type Trie struct {
	root *trieNode
}

func (t *Trie) Put(s string, v int64) {
	t.root = putHelper(t.root, []rune(s), v, 0)
}

func putHelper(t *trieNode, r []rune, v int64, d int) *trieNode {
	if t == nil {
		t = newTrieNode()
	}
	if len(r) == d {
		*t.val = append(*t.val, v)
	} else {
		t.next[r[d]] = putHelper(t.next[r[d]], r, v, d+1)
	}
	return t
}

func (t *Trie) Get(s string) (v []int64) {
	n := getHelper(t.root, []rune(s), 0)
	if n != nil {
		v = *n.val
	}
	return
}

func getHelper(t *trieNode, r []rune, d int) *trieNode {
	if t == nil || len(r) == d {
		return t
	}

	return getHelper(t.next[r[d]], r, d+1)
}

func (t *Trie) Keys() []string {
	return t.KeysWithPrefix("")
}

func (t *Trie) KeysWithPrefix(pre string) []string {
	s := []string{}
	collect(getHelper(t.root, []rune(pre), 0), pre, &s)
	return s
}

func collect(t *trieNode, pre string, s *[]string) {
	if t == nil {
		return
	}
	if len(*t.val) != 0 {
		*s = append(*s, pre)
	}
	for r, n := range t.next {
		collect(n, pre+string(r), s)
	}
}

func (t *Trie) Vals() []int64 {
	return t.ValsWithPrefix("")
}

func (t *Trie) ValsWithPrefix(pre string) []int64 {
	s := []int64{}
	collectVal(getHelper(t.root, []rune(pre), 0), &s)
	return s
}

func collectVal(t *trieNode, s *[]int64) {
	if t == nil {
		return
	}
	if len(*t.val) != 0 {
		*s = append(*s, *t.val...)
	}
	for _, n := range t.next {
		collectVal(n, s)
	}
}
