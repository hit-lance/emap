package etaxi

type StreetMap struct {
	*Graph
	NodeSet
}

func NewStreetMapFrom(fn string, ns NodeSet) *StreetMap {
	var sm StreetMap
	sm.Graph = NewGraphFrom(fn)
	sm.NodeSet = ns
	for nid, n := range sm.nodes {
		if sm.Neighbors(nid) != nil {
			ns.Insert(n)
		}
	}

	return &sm
}

func (sm *StreetMap) Closest(lat, lon float64) int64 {
	return sm.Nearest(&Node{lat: lat, lon: lon}).id
}
