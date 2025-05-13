package longest_subarray_after_deletion

/*
Тебе дан бинарный массив nums (только 0 и 1).
Ты можешь удалить ровно один элемент (любой, не обязательно 0).
Найди максимальную длину подмассива, состоящего только из 1, после одного удаления.

Input:  [0,1,1,1,0,1,1,0,1]
Output: 5  // удаляем 0 между двумя сериями
*/
func longestSubArrayAfterDeletion(nums []int) int {

	start, zeroes, maxLength := 0, 0, 0

	for end := range nums {
		if nums[end] == 0 {
			zeroes++
		}

		for zeroes > 1 {
			if nums[start] == 0 {
				zeroes--

			}
			start++
		}

		maxLength = max(maxLength, (end - start + 1))
	}

	return maxLength - 1
}
