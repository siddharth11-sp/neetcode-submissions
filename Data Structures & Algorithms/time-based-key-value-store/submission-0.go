type Pair struct {
	Value     string
	TimeStamp int
}

type TimeMap struct {
	m map[string][]Pair
}

func Constructor() *TimeMap {
	return &TimeMap{
		m: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Pair{
		Value:     value,
		TimeStamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, ok := this.m[key]
	if !ok {
		return ""
	}

	left := 0
	right := len(values) - 1
	ans := ""
	for left <= right {
		mid := left + (right-left)/2

		if values[mid].TimeStamp <= timestamp {
			ans = values[mid].Value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}