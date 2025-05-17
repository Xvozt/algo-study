package search_insert_position

/*

Дан отсортированный по возрастанию массив nums и target.
Найди индекс, по которому должен стоять target.
Если target уже есть — верни его индекс.
Если нет — верни индекс, куда его вставить, чтобы сохранить порядок.


Input:  nums = [1,3,5,6], target = 5
Output: 2

Input:  nums = [1,3,5,6], target = 2
Output: 1

Input:  nums = [1,3,5,6], target = 7
Output: 4

*/

func searchInsertPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left

}
