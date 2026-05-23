func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := max(piles)
	minSpeed := right
	for left <= right {
		mid := left + (right-left)/2
		if canFinish(piles,h,mid){
			right = mid -1 
			if mid < minSpeed{
				minSpeed = mid 
			}
		} else {
			left = mid + 1
		}
	}

	return minSpeed

}

func canFinish(piles []int, h int, speed int) bool {
	hours := 0
	for _, pile := range piles {
		hours += (speed+pile-1)/speed
		if hours > h {
			return false
		}
	}
	return true
}

func max(piles []int) int {
	m := piles[0]
	for i := 1 ; i < len(piles); i++ {
		if piles[i] > m {
			m = piles[i]
		}
	} 
	return m
}