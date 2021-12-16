package streetmap

import "strings"


type NaiveNameDict map[string]int64

func (nnd NaiveNameDict) put(s string, v int64) {
	nnd[s] = v
}

func (nnd NaiveNameDict) get(s string) (v int64) {
	if v, ok := nnd[s]; ok {
		return v
	} else {
		return INVALID_NODE_ID
	}
}

func (nnd *NaiveNameDict) keys() []string {
	return nnd.keysWithPrefix("")
}

func (nnd NaiveNameDict) keysWithPrefix(pre string) []string {
	s := []string{}
	for k := range nnd {
		if strings.HasPrefix(k, pre) {
			s = append(s, k)
		}
	}
	return s
}
