type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(val interface{}) {
	*m = append(*m, val.(int))
}

func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)

	val := old[n-1]
	*m = old[:n-1]

	return val
}


func lastStoneWeight(stones []int) int {
h := maxHeap(stones)

	heap.Init(&h)

	for h.Len() > 1 {

		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)

		if y != x {
			heap.Push(&h, y-x)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return h[0]
}
