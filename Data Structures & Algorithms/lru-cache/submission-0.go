type LRUCache struct {
	data     map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

func Constructor(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		data:     make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) moveToFront(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head

	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) remove(node *Node) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if ok {
		this.remove(node)
		this.moveToFront(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.data[key]
	if ok {
		node.Value = value
		this.remove(node)
		this.moveToFront(node)
		return
	}

	// does not exsist
	newNode := &Node{
		Key:   key,
		Value: value,
	}

	this.data[key] = newNode
	this.moveToFront(newNode)

	if len(this.data) > this.capacity {
		lru := this.tail.Prev
		this.remove(lru)
		delete(this.data, lru.Key)
	}
}