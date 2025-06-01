package majorityelement

/*
🧠 Задача: Majority Element

Дан массив целых чисел nums, в котором всегда существует элемент, встречающийся более чем ⌊n / 2⌋ раз.

Нужно вернуть этот элемент.


Input:  nums = [3,2,3]
Output: 3

Input:  nums = [2,2,1,1,1,2,2]
Output: 2
*/

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
		if frequencyMap[num] > len(nums)/2 {
			return num
		}

	}
	return -1
}
