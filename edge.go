package etaxi

import "fmt"

type Edge struct {
	u, v   int64
	weight float64
	name   string
}

func (e *Edge) String() string {
	return fmt.Sprintf("%+v\n", *e)
}
