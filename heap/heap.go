package heap

type Heap[T any] struct {
	data  []T
	n     int
	less  func(a, b T) bool
	equal func(a, b T) bool
}

func NewHeap[T any](less, equal func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:  make([]T, 0),
		n:     0,
		less:  less,
		equal: equal,
	}
}

func (h *Heap[T]) shiftDown(i int) {
	n := len(h.data)
	for 2*i+1 < n {
		left, right := 2*i+1, 2*i+2
		k := left
		if right < n && h.less(h.data[right], h.data[k]) {
			k = right
		}

		// if current element is less than his ance sons, heap is sorted
		if h.less(h.data[i], h.data[k]) || h.equal(h.data[i], h.data[k]) {
			break
		}
		h.data[i], h.data[k] = h.data[k], h.data[i]
		i = k
	}
}

func (h *Heap[T]) shiftUp(i int) {
	for h.less(h.data[i], h.data[(i-1)/2]) {
		h.data[i], h.data[(i-1)/2] = h.data[(i-1)/2], h.data[i]
		i = (i - 1) / 2
	}
}

func (h *Heap[T]) Top() T {
	return h.data[0]
}

func (h *Heap[T]) Pop() T {
	mn := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.n--
	h.shiftDown(0)
	return mn
}

func (h *Heap[T]) Push(x T) {
	h.n++
	if h.n > len(h.data) {
		h.data = append(h.data, x)
	} else {
		h.data[h.n-1] = x
	}

	h.shiftUp(h.n - 1)
}
