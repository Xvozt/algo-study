package remove_duplicates_sorted_array

/*
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	write := 1

	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
		}
	}

	return write
}
