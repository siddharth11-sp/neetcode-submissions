type MinHeap struct {
    arr []int
}


type KthLargest struct {
    k    int
    heap MinHeap
}

func Constructor(k int, nums []int) KthLargest {

    kl := KthLargest{
        k: k,
    }

    for _, num := range nums {
        kl.Add(num)
    }

    return kl
}

func (this *KthLargest) Add(val int) int {

    this.heap.Push(val)

    if len(this.heap.arr) > this.k {
        this.heap.Pop()
    }

    return this.heap.Peek()
}

func (h *MinHeap) Push(val int) {
    h.arr = append(h.arr, val)

    idx := len(h.arr) - 1

    for idx > 0 {
        parent := (idx - 1) / 2

        if h.arr[parent] <= h.arr[idx] {
            break
        }

        h.arr[parent], h.arr[idx] =
            h.arr[idx], h.arr[parent]

        idx = parent
    }
}

func (h *MinHeap) Pop() int {
    n := len(h.arr)

    minVal := h.arr[0]

    h.arr[0] = h.arr[n-1]
    h.arr = h.arr[:n-1]

    idx := 0

    for {
        left := 2*idx + 1
        right := 2*idx + 2

        smallest := idx

        if left < len(h.arr) &&
            h.arr[left] < h.arr[smallest] {
            smallest = left
        }

        if right < len(h.arr) &&
            h.arr[right] < h.arr[smallest] {
            smallest = right
        }

        if smallest == idx {
            break
        }

        h.arr[idx], h.arr[smallest] =
            h.arr[smallest], h.arr[idx]

        idx = smallest
    }

    return minVal
}

func (h *MinHeap) Peek() int {
    return h.arr[0]
}