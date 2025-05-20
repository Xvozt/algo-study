package targetsearch

/*
Input:  nums = [-1,0,3,5,9,12], target = 9
Output: 4
*/

func targetSearch(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
