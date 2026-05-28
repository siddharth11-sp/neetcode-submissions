func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x := len(nums1)
	y := len(nums2)

	low, high := 0, x

	for low <= high {

		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX != 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX != x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY != 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY != y {
			minRightY = nums2[partitionY]
		}

		// Correct partition found
		if maxLeftX <= minRightY && maxLeftY <= minRightX {

			// Even length
			if (x+y)%2 == 0 {
				return float64(
					max(maxLeftX, maxLeftY)+
						min(minRightX, minRightY),
				) / 2.0
			}

			// Odd length
			return float64(max(maxLeftX, maxLeftY))
		}

		// Move left
		if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			// Move right
			low = partitionX + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}