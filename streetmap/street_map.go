package streetmap

type StreetMap struct {
	*Graph
	ns NodeSet
	nd NameDict
}

func NewStreetMapFrom(fn string, ns NodeSet, nd NameDict) *StreetMap {
	sm := StreetMap{}
	sm.Graph = NewGraphFrom(fn)
	sm.ns = ns
	sm.nd = nd
	for nid, n := range sm.nodes {
		if sm.Neighbors(nid) != nil {
			ns.Insert(n)
		}
		if n.name != "" {
			sm.nd.put(n.name, n.id)
		}
	}

	return &sm
}

func (sm *StreetMap) Closest(lat, lon float64) int64 {
	return sm.ns.Nearest(&Node{lat: lat, lon: lon}).id
}

// func (sm *StreetMap) getNodesByPrefix(name string) []string {
// 	return sm.nd.keysWithPrefix("")
// }

// func (sm *StreetMap) GetNodeIdByName(name string) int64 {
// 	return sm.nd.get(name)
// }
