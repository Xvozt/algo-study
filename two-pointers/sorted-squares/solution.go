package sorted_squares

import "math"

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	left, right, pos := 0, len(nums)-1, len(nums)-1

	for left <= right {

		if math.Abs(float64(nums[left])) > math.Abs(float64(nums[right])) {
			result[pos] = nums[left] * nums[left]
			left++
		} else {
			result[pos] = nums[right] * nums[right]
			right--
		}

		pos--

	}

	return result
}
