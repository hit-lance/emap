package streetmap

type NameDict interface {
	put(s string, v int64)
	get(s string) int64
	vals() []int64
	valsWithPrefix(pre string) []int64
}
