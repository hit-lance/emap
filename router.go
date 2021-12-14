package tinymap

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"os"
)

type solve func(sm *StreetMap, src, dst int64) *list.List

type Router struct {
}

func (r Router) ShortestPath(s solve, sm *StreetMap, slat, slon, dlat, dlon float64) *list.List {
	src := sm.Closest(slat, slon)
	dst := sm.Closest(dlat, dlon)
	return s(sm, src, dst)
}

func dijkstra(sm *StreetMap, src, dst int64) (sol *list.List) {
	if _, ok := sm.nodes[src]; !ok {
		fmt.Fprintf(os.Stderr, "node id %d doesn't exist", src)
		os.Exit(1)
	}

	sol = list.New()
	pq := make(PriorityQueue, 0)

	itemPtrs := make(map[int64]*Item)
	distTo := make(map[int64]float64)
	edgeTo := make(map[int64]int64)

	for nid := range sm.nodes {
		var dist float64
		if nid == src {
			dist = 0.0
		} else {
			dist = math.MaxFloat64
		}
		item := &Item{value: nid, priority: dist}
		heap.Push(&pq, item)
		itemPtrs[nid] = item
		distTo[nid] = dist
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.value == dst {
			pos := dst
			for pos != src {
				sol.PushFront(pos)
				pos = edgeTo[pos]
			}
			sol.PushFront(pos)
			return
		}

		p := item.value
		for _, nbr := range sm.Neighbors(item.value) {
			var q int64
			if p == nbr.u {
				q = nbr.v
			} else {
				q = nbr.u
			}
			if item.priority+nbr.weight < distTo[q] {
				distTo[q] = item.priority + nbr.weight
				edgeTo[q] = p
				pq.update(itemPtrs[q], itemPtrs[q].value, distTo[q])
			}
		}
	}
	return
}

func aStar(sm *StreetMap, src, dst int64) (sol *list.List) {
	if _, ok := sm.nodes[src]; !ok {
		fmt.Fprintf(os.Stderr, "node id %d doesn't exist", src)
		os.Exit(1)
	}

	sol = list.New()
	pq := make(PriorityQueue, 0)

	itemPtrs := make(map[int64]*Item)
	distTo := make(map[int64]float64)
	edgeTo := make(map[int64]int64)
	heuristic := make(map[int64]float64)

	for nid := range sm.nodes {
		if nid != src {
			distTo[nid] = math.MaxFloat64
		}
	}

	heuristic[src] = distance(sm.nodes[src], sm.nodes[dst])
	item := &Item{value: src, priority: heuristic[src]}
	heap.Push(&pq, item)
	itemPtrs[src] = item
	distTo[src] = 0.0

	for pq.Len() > 0 {
		item = heap.Pop(&pq).(*Item)
		if item.value == dst {
			pos := dst
			for pos != src {
				sol.PushFront(pos)
				pos = edgeTo[pos]
			}
			sol.PushFront(pos)
			return
		}

		p := item.value
		for _, nbr := range sm.Neighbors(item.value) {
			var q int64
			if p == nbr.u {
				q = nbr.v
			} else {
				q = nbr.u
			}

			if distTo[p]+nbr.weight < distTo[q] {
				distTo[q] = distTo[p] + nbr.weight
				edgeTo[q] = p

				if _, ok := heuristic[q]; !ok {
					heuristic[q] = distance(sm.nodes[q], sm.nodes[dst])
				}

				if ptr, ok := itemPtrs[q]; ok {
					pq.update(ptr, ptr.value, distTo[q]+heuristic[q])
				} else {
					i := &Item{value: q, priority: distTo[q] + heuristic[q]}
					heap.Push(&pq, i)
					itemPtrs[q] = i
				}
			}
		}
	}
	return
}
