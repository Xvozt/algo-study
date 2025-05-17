package maximum_sum_subarray

/*
Input: nums = [2, 1, 5, 1, 3, 2], k = 3
Output: 9
Explanation: [5, 1, 3] даёт максимум
*/

func maximumSum(nums []int, k int) int {
	if len(nums) < k {
		k = len(nums)
	}

	curSum := 0

	for i := range k {
		curSum += nums[i]
	}

	maxSum := curSum

	for i := k; i < len(nums); i++ {
		curSum = curSum + nums[i] - nums[i-k]
		maxSum = max(curSum, maxSum)
	}

	return maxSum

}
