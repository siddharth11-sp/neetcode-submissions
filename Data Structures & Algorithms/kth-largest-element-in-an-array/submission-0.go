
type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *maxHeap) Pop() any {
	old := *m
	n := len(*m)
	val := old[n-1]
	*m = old[:n-1]
	return val
}

func (m *maxHeap) Push(val interface{}) {
	*m = append(*m, val.(int))
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func findKthLargest(nums []int, k int) int {
	h := maxHeap(nums)
	heap.Init(&h)

	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}
	
	return heap.Pop(&h).(int)
}