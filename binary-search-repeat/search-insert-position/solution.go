package searchinsertposition

/*
Дан отсортированный массив nums и число target.
Найди индекс, по которому target должен быть вставлен, чтобы порядок сохранился.

Если target уже есть — верни его индекс.
Если нет — верни индекс, куда бы ты его вставил.

Input:  nums = [1,3,5,6], target = 5
Output: 2

Input:  nums = [1,3,5,6], target = 2
Output: 1

Input:  nums = [1,3,5,6], target = 7
Output: 4
*/
func searchInsert(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}

	}
	return left
}
