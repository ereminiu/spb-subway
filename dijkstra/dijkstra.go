package dijkstra

import (
	"github.com/ereminiu/spb-subway/heap"
	. "github.com/ereminiu/spb-subway/models"
)

type Item struct {
	x int
	d int
}

func Less(e, other Item) bool {
	return e.d < other.d
}

func LessOrEqual(e, other Item) bool {
	return e.d <= other.d
}

func Dijkstra(gr [][]*Edge, u, v int) (int, []int) {
	n := len(gr)
	inf := int(1e9 + 228)

	dist := make([]int, n)
	par := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[u] = 0
	par[u] = -1

	h := heap.NewHeap(func(a, b Item) bool { return a.d < b.d }, func(a, b Item) bool { return a.d <= b.d })
	h.Push(Item{x: u, d: dist[u]})

	for !h.Empty() {
		v := h.Pop().x

		for _, e := range gr[v] {
			to, cost := e.To, e.Cost
			if dist[to] > dist[v]+cost {
				dist[to] = dist[v] + cost
				par[to] = v
				h.Push(Item{to, dist[v] + cost})
			}
		}
	}

	path := make([]int, 0)
	cur := v
	for cur != -1 {
		path = append(path, cur)
		cur = par[cur]
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}

	return dist[v], path
}
