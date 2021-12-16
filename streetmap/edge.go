package streetmap

import "fmt"

type Edge struct {
	u, v   int64
	weight float64
	name   string
}

func (e *Edge) String() string {
	return fmt.Sprintf("%+v\n", *e)
}

func (e *Edge) U() int64 {
	return e.u
}

func (e *Edge) V() int64 {
	return e.v
}

func (e *Edge) Weight() float64 {
	return e.weight
}
