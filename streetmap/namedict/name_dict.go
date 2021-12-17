package streetmap

type NameDict interface {
	Put(s string, v int64)
	Get(s string) int64
	Keys() []string
	KeysWithPrefix(pre string) []string
}
