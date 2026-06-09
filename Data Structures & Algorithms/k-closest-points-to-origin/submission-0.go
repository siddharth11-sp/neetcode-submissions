type Point struct {
	x int
	y int
	dis int
}

type MaxHeap []Point

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].dis > h[j].dis // max heap
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, p := range points {
		x, y := p[0], p[1]
		dist := x*x + y*y

		heap.Push(h, Point{
			x : x,
			y : y,
			dis : dist,
		})

		if h.Len() > k{
			heap.Pop(h)
		}
	}

	ans := make([][]int,0,k)
	for h.Len() > 0 {
		p := heap.Pop(h).(Point)
		ans = append(ans, []int{p.x, p.y})
	}

	return ans
}
