package boyermoormajority

func majorityElementBoyerMoore(nums []int) int {
	candidate := nums[0]
	count := 1

	for _, num := range nums[1:] {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	return candidate
}
