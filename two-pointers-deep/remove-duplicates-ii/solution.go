package removeduplicatesii

/*
Дан отсортированный массив nums.
Удалите дубликаты на месте, так чтобы каждый элемент встречался не более двух раз.
Верните новую длину массива. Остальная часть массива может быть любой.
Input:  [1,1,1,2,2,3]
Output: 5 // массив становится [1,1,2,2,3,_]
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	write := 2

	for read := 2; read < len(nums); read++ {
		if nums[read] != nums[write-2] {
			nums[write] = nums[read]
			write++
		}
	}

	return write
}
