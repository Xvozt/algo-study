package minimuminrotatedarray

/*
Дан отсортированный, но повёрнутый массив nums без дубликатов.
Найди минимальный элемент.

Input:  [3,4,5,1,2]
Output: 1

Input:  [4,5,6,7,0,1,2]
Output: 0
*/

func minimumInRotated(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
