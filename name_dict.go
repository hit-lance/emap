package etaxi

type NameDict interface {
	put(s string, v int64)
	get(s string) int64
	keys() []string
	keysWithPrefix(pre string) []string
}
