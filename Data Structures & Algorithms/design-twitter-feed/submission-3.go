type Twitter struct {
	data        map[int][]Tweet
	followerMap map[int]map[int]bool
	time        int
}

type Tweet struct {
	id   int
	time int
}

type HeapNode struct {
	tweet    Tweet
	userId   int
	tweetIdx int
}

type MaxHeap []HeapNode

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].tweet.time > h[j].tweet.time
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(val interface{}) {
	*h = append(*h, val.(HeapNode))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func Constructor() *Twitter {
	return &Twitter{
		data:        make(map[int][]Tweet),
		followerMap: make(map[int]map[int]bool),
		time:        0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++
	this.data[userId] = append(this.data[userId], Tweet{
		id:   tweetId,
		time: this.time,
	})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &MaxHeap{}
	heap.Init(h)

	userIds := []int{userId}

	for followee := range this.followerMap[userId] {
		if userId == followee {
			continue
		}
		userIds = append(userIds, followee)
	}

	for _, uid := range userIds {
		tweets := this.data[uid]
		if len(tweets) == 0 {
			continue
		}

		lastIdx := len(tweets) - 1

		heap.Push(h, HeapNode{
			tweet:    tweets[lastIdx],
			userId:   uid,
			tweetIdx: lastIdx,
		})
	}

	var result []int

	for h.Len() > 0 && len(result) < 10 {
		node := heap.Pop(h).(HeapNode)

		result = append(result, node.tweet.id)

		nextIdx := node.tweetIdx - 1

		if nextIdx >= 0 {
			heap.Push(h, HeapNode{
				userId:   node.userId,
				tweet:    this.data[node.userId][nextIdx],
				tweetIdx: nextIdx,
			})
		}

	}

	return result

}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followerMap[followerId] == nil {
		this.followerMap[followerId] = make(map[int]bool)
	}

	this.followerMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followerMap[followerId] != nil {
		delete(this.followerMap[followerId], followeeId)
	}
}