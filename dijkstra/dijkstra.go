package dijkstra

import (
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

type Heap struct {
	n int    // heap len
	a []Item // heap elements
}

func NewHeap() *Heap {
	return &Heap{n: 0, a: make([]Item, 0)}
}

func (h *Heap) shiftDown(i int) {
	for 2*i+1 < h.n {
		left, right := 2*i+1, 2*i+2
		k := left
		if right < h.n && Less(h.a[right], h.a[k]) {
			k = right
		}

		// if current element is less than his ance sons, heap is sorted
		if LessOrEqual(h.a[i], h.a[k]) {
			break
		}
		h.a[i], h.a[k] = h.a[k], h.a[i]
		i = k
	}
}

func (h *Heap) shiftUp(i int) {
	for Less(h.a[i], h.a[(i-1)/2]) {
		h.a[i], h.a[(i-1)/2] = h.a[(i-1)/2], h.a[i]
		i = (i - 1) / 2
	}
}

func (h *Heap) Min() Item {
	min := h.a[0]
	h.a[0] = h.a[h.n-1]
	h.n--
	h.shiftDown(0)
	return min
}

func (h *Heap) Append(x Item) {
	h.n++
	if h.n > len(h.a) {
		h.a = append(h.a, x)
	} else {
		h.a[h.n-1] = x
	}

	h.shiftUp(h.n - 1)
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

	h := NewHeap()
	h.Append(Item{x: u, d: dist[u]})

	for h.n > 0 {
		v := h.Min().x

		for _, e := range gr[v] {
			to, cost := e.To, e.Cost
			if dist[to] > dist[v]+cost {
				dist[to] = dist[v] + cost
				par[to] = v
				h.Append(Item{to, dist[v] + cost})
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
