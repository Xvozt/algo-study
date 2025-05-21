package firstandlastposition

/*
Input:  nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Input:  nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
func searchRange(nums []int, target int) []int {

	first := binarySearch(nums, target, true)
	second := binarySearch(nums, target, false)

	if first == -1 || second == -1 {
		return []int{-1, -1}
	}

	return []int{first, second}
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
