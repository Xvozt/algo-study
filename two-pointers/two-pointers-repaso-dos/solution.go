package remove_element

/*

Дан массив nums и значение val. Удалите все вхождения val на месте и верните новую длину.
Порядок можно не сохранять, но нельзя использовать дополнительный массив.
Input:  nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]

*/

func removeElement(nums []int, val int) int {

	write := 0

	for read := range nums {
		if nums[read] != val {
			nums[write] = nums[read]
			write++
		}
	}

	return write

}
