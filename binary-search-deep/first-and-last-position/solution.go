package firstandlastposition

func firstAndLastPosition(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false)

	if left == -1 || right == -1 {
		return []int{-1, -1}
	}
	return []int{left, right}

}

func binarySearch(nums []int, target int, first bool) int {
	left, right := 0, len(nums)-1
	if first {
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] >= target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			}
		}

		if left < len(nums) && nums[left] == target {
			return left
		}
	} else {
		for left <= right {
			mid := left + (right-left)/2

			if nums[mid] <= target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			}
		}
		if right >= 0 && nums[right] == target {
			return right
		}
	}

	return -1
}
