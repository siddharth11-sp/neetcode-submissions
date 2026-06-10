type Task struct {
	count int
	char  byte
}
type maxHeap []Task

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxHeap) Less(i, j int) bool {
	return m[i].count > m[j].count
}

func (m *maxHeap) Push(val interface{}) {
	*m = append(*m, val.(Task))
}

func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)

	val := old[n-1]
	*m = old[:n-1]
	return val
}

type CoolDown struct {
	task      Task
	readyTime int
}

func leastInterval(tasks []byte, n int) int {
	h := &maxHeap{}
	heap.Init(h)

	freq := make(map[byte]int)

	for _, t := range tasks {
		freq[t]++
	}

	for ch, freq := range freq {
		heap.Push(h, Task{
			char:  ch,
			count: freq,
		})
	}

	queue := []CoolDown{}
	time := 0

	for h.Len() > 0 || len(queue) > 0 {
		time++

		if len(queue) > 0 && queue[0].readyTime == time {
			heap.Push(h, queue[0].task)
			queue = queue[1:]
		}

		if h.Len() > 0 {
			curr := heap.Pop(h).(Task)
			curr.count--

			if curr.count > 0 {
				queue = append(queue, CoolDown{
					task:      curr,
					readyTime: time + n + 1,
				})
			}

		}
	}
	return time
}