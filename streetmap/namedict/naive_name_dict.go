package streetmap

import (
	"strings"
)

type NaiveNameDict map[string]*[]int64

func (nnd NaiveNameDict) Put(s string, v int64) {
	if n, ok := nnd[s]; ok {
		*n = append(*n, v)
	} else {
		nnd[s] = &[]int64{v}
	}
}

func (nnd NaiveNameDict) Get(s string) (ret []int64) {
	if v, ok := nnd[s]; ok {
		ret = *v
	}
	return
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

func (nnd NaiveNameDict) ValsWithPrefix(pre string) []int64 {
	s := []int64{}
	for k, v := range nnd {
		if strings.HasPrefix(k, pre) {
			for i := range *v {
				s = append(s, (*v)[i])
			}
		}
	}
	return s
}
