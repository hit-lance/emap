package etaxi

type StreetMap struct {
	*Graph
	NodeSet
	trie *Trie
}

func NewStreetMapFrom(fn string, ns NodeSet) *StreetMap {
	sm := StreetMap{}
	sm.Graph = NewGraphFrom(fn)
	sm.NodeSet = ns
	sm.trie = &Trie{}
	for nid, n := range sm.nodes {
		if sm.Neighbors(nid) != nil {
			ns.Insert(n)
		}
		if n.name != "" {
			sm.trie.put(n.name, n.id)
		}
	}

	return &sm
}

func (sm *StreetMap) Closest(lat, lon float64) int64 {
	return sm.Nearest(&Node{lat: lat, lon: lon}).id
}
