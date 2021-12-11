package tinymap

import (
	"testing"
)

func TestGraph(t *testing.T) {

	t.Run("add node", func(t *testing.T) {
		g := NewGraph()

		g.AddNode(&Node{1, 32.9697, -96.80322, "a"})
		g.AddNode(&Node{2, 29.46786, -98.53506, "b"})

		g.AddEdge(1, 2, "ab")

		got := *g.Neighbors(1)[0]
		want := Edge{u: 1, v: 2, weight: 422.7592707099526, name: "ab"}
		if got != want {
			t.Errorf("got %+v sent but expected %+v", got, want)
		}
	})

	t.Run("new graph from xml file", func(t *testing.T) {
		fn := "./tiny-clean.osm.xml"
		g := NewGraphFrom(fn)

		got := len(g.nodes)
		want := 7
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})

}
