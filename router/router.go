package router

import (
	"container/heap"
	"container/list"
	sm "etaxi/streetmap"
	"fmt"
	"math"
	"os"
)

type solve func(m *sm.StreetMap, src, dst int64) *list.List

type Router struct {
}

func (r Router) ShortestPath(s solve, m *sm.StreetMap, slat, slon, dlat, dlon float64) *list.List {
	src := m.Closest(slat, slon)
	dst := m.Closest(dlat, dlon)
	return s(m, src, dst)
}

func (r Router) RouteDirections(m *sm.StreetMap, route *list.List) (res []NavigationDirection) {
	if route == nil || route.Len() < 1 {
		fmt.Fprintln(os.Stderr, "got wrong input route.")
		os.Exit(1)
	}
	if route.Len() == 1 {
		return
	}

	for p := route.Front(); p.Next() != nil; p = p.Next() {
		fmt.Println(m.GetEdge(p.Value.(int64), p.Next().Value.(int64)).Name())
	}
	return
}

func dijkstra(m *sm.StreetMap, src, dst int64) (sol *list.List) {
	if !m.Contains(src) {
		fmt.Fprintf(os.Stderr, "node id %d doesn't exist", src)
		os.Exit(1)
	}

	sol = list.New()
	pq := make(PriorityQueue, 0)

	itemPtrs := make(map[int64]*Item)
	distTo := make(map[int64]float64)
	edgeTo := make(map[int64]int64)

	for _, nid := range m.NodeIds() {
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
		for _, nbr := range m.Neighbors(item.value) {
			var q int64
			if p == nbr.U() {
				q = nbr.V()
			} else {
				q = nbr.U()
			}
			if item.priority+nbr.Weight() < distTo[q] {
				distTo[q] = item.priority + nbr.Weight()
				edgeTo[q] = p
				pq.update(itemPtrs[q], itemPtrs[q].value, distTo[q])
			}
		}
	}
	return
}

func aStar(m *sm.StreetMap, src, dst int64) (sol *list.List) {
	if !m.Contains(src) {
		fmt.Fprintf(os.Stderr, "node id %d doesn't exist", src)
		os.Exit(1)
	}

	sol = list.New()
	pq := make(PriorityQueue, 0)

	itemPtrs := make(map[int64]*Item)
	distTo := make(map[int64]float64)
	edgeTo := make(map[int64]int64)
	heuristic := make(map[int64]float64)

	for _, nid := range m.NodeIds() {
		if nid != src {
			distTo[nid] = math.MaxFloat64
		}
	}

	heuristic[src] = m.GetNode(src).Distance(m.GetNode(dst))
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
		for _, nbr := range m.Neighbors(item.value) {
			var q int64
			if p == nbr.U() {
				q = nbr.V()
			} else {
				q = nbr.U()
			}

			if distTo[p]+nbr.Weight() < distTo[q] {
				distTo[q] = distTo[p] + nbr.Weight()
				edgeTo[q] = p

				if _, ok := heuristic[q]; !ok {
					heuristic[q] = m.GetNode(q).Distance(m.GetNode(dst))
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
