package move_zeroes_in_place

/*
Input:  [0,1,0,3,12]
Output: [1,3,12,0,0]
*/

func moveZeroes(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	write := 0

	// placing non zero element on write pointer making nums look like [1,3,12,3,12]
	for read := range nums {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}

	// now we need to make all elements above write position zeroes
	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}
