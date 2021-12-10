package tinymap

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {

	// t.Run("add node", func(t *testing.T) {
	// 	g := NewGraph()
	// 	n := Node{1, 32.9697, -96.80322, "a"}
	// 	g.AddNode(&n)
	// 	n1 := Node{2, 29.46786, -98.53506, "b"}
	// 	g.AddNode(&n1)
	// 	g.AddEdge(1, 2, "ab")

	// 	// fmt.Println(g.adj[1][0])
	// 	// fmt.Println(g.adj[2][0])
	// 	fmt.Print(g)
	// })

	t.Run("new graph from xml file", func(t *testing.T) {
		fn := "./tiny-clean.osm.xml"
		g := NewGraphFrom(fn)
		// n := Node{1, 32.9697, -96.80322, "a"}
		// g.AddNode(&Node{1, 32.9697, -96.80322, "a"})
		// g.AddNode(&Node{2, 29.46786, -98.53506, "b"})
		// g.AddEdge(1, 2, "ab")

		// fmt.Println(g.adj[1][0])
		// fmt.Println(g.adj[2][0])
		fmt.Print(g)
	})

}
