package graph

import "fmt"

type Edge struct {
	from, to int64
	weight   float64
	name     string
}

func (e *Edge) String() string {
	return fmt.Sprintf("%+v\n", *e)
}

func (e *Edge) From() int64 {
	return e.from
}

func (e *Edge) To() int64 {
	return e.to
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) Name() string {
	return e.name
}
