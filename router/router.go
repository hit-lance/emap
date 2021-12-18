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

func GetDirectionsText(nd []NavigationDirection) (s string) {
	dist := 0.0
	for _, d := range nd {
		dist += d.distance
	}
	s += fmt.Sprintf("全程%.3f公里\n", dist)
	for _, d := range nd {
		s += fmt.Sprintln(d)
	}
	s += fmt.Sprintln("到达目的地")
	return
}

func (r Router) RouteDirections(m *sm.StreetMap, route *list.List) (nd []NavigationDirection) {
	if route == nil || route.Len() < 1 {
		fmt.Fprintln(os.Stderr, "got wrong input route.")
		os.Exit(1)
	}

	p := route.Front()
	var direction DirectionType
	var distance, prevBearing float64
	var prevWayName string

	for {
		cur, next := p.Value.(int64), p.Next().Value.(int64)
		way := m.GetEdge(cur, next)
		curBearing := m.GetNode(cur).Bearing(m.GetNode(next))

		if p == route.Front() {
			direction = Start
			prevWayName = way.Name()
			distance = way.Weight()
		} else {
			if prevWayName != "" && way.Name() == prevWayName {
				distance += way.Weight()
			} else {
				nd = append(nd, NavigationDirection{direction: direction, way: prevWayName, distance: distance})
				direction = getDirection(prevBearing, curBearing)
				prevWayName = way.Name()
				distance = way.Weight()
			}
		}
		if p.Next().Next() == nil {
			nd = append(nd, NavigationDirection{direction: direction, way: prevWayName, distance: distance})
			break
		}
		prevBearing = curBearing
		p = p.Next()

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
		for _, nbr := range m.Neighbors(p) {
			q := nbr.To()
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
		for _, nbr := range m.Neighbors(p) {
			q := nbr.To()

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
