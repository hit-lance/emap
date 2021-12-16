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

func (nnd *NaiveNameDict) vals() []int64 {
	return nnd.valsWithPrefix("")
}

func (nnd NaiveNameDict) valsWithPrefix(pre string) []int64 {
	s := []int64{}
	for k, v := range nnd {
		if strings.HasPrefix(k, pre) {
			s = append(s, v)
		}
	}
	return s
}
