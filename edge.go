package tinymap

import "fmt"

type Edge struct {
	u, v   int64
	weight float64
	name   string
}

func (e *Edge) String() string {
	s := fmt.Sprintf("%+v\n", *e)
	return s
}
