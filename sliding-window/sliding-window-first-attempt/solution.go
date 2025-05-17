package max_consecutive_ones

/*
Input:  [1,1,0,1,1,1]
Output: 3
*/

func countMaxConsecutiveOnes(nums []int) int {
	maxCount, currentCount := 0, 0

	for i := range nums {
		if nums[i] == 1 {
			currentCount++
			maxCount = max(currentCount, maxCount)
		} else {
			currentCount = 0
		}

	}

	return maxCount
}
