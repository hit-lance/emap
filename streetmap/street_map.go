package streetmap

import (
	"etaxi/streetmap/graph"

	nd "etaxi/streetmap/namedict"
	ns "etaxi/streetmap/nodeset"
)

type StreetMap struct {
	*graph.Graph
	ns.NodeSet
	nd.NameDict
}

func NewStreetMap(fn string) *StreetMap {
	return NewStreetMapFrom(fn, &ns.KDTree{}, &nd.Trie{})
}

func NewStreetMapFrom(fn string, ns ns.NodeSet, nd nd.NameDict) *StreetMap {
	sm := StreetMap{}
	sm.Graph = graph.NewGraphFrom(fn)
	sm.NodeSet = ns
	sm.NameDict = nd
	for _, nid := range sm.NodeIds() {
		n := sm.GetNodeById(nid)
		if sm.Neighbors(nid) != nil {
			ns.Insert(n)
		}
		if n.Name() != "" {
			sm.NameDict.Put(n.Name(), n.Id())
		}
	}

	return &sm
}

func (sm *StreetMap) Closest(lat, lon float64) int64 {
	return sm.NodeSet.Nearest(graph.NewNode(lat, lon)).Id()
}

func (sm *StreetMap) GetNodesByPrefix(name string) []string {
	return sm.NameDict.KeysWithPrefix(name)
}

func (sm *StreetMap) GetNodeIdByName(name string) int64 {
	return sm.NameDict.Get(name)
}
