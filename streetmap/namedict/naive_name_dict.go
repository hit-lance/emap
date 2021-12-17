package streetmap

import (
	"etaxi/streetmap/graph"
	"strings"
)

type NaiveNameDict map[string]int64

func (nnd NaiveNameDict) Put(s string, v int64) {
	nnd[s] = v
}

func (nnd NaiveNameDict) Get(s string) (v int64) {
	if v, ok := nnd[s]; ok {
		return v
	} else {
		return graph.INVALID_NODE_ID
	}
}

func (nnd *NaiveNameDict) Keys() []string {
	return nnd.KeysWithPrefix("")
}

func (nnd NaiveNameDict) KeysWithPrefix(pre string) []string {
	s := []string{}
	for k := range nnd {
		if strings.HasPrefix(k, pre) {
			s = append(s, k)
		}
	}
	return s
}
