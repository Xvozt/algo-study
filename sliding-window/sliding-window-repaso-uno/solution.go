package max_consecutive_ones_iii

/*
Дан бинарный массив nums (только 0 и 1) и число k.
Можно "перевернуть" до k нулей в единицы.
Найди максимальное количество подряд идущих единиц, которое можно получить.

Input:  nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Можно перевернуть два нуля — [1,1,1,0,0,1,1,1,1,1,1] → шесть подряд
*/

func maxConsecutiveOnesIII(nums []int, k int) int {

	start, zeroes, maxLength := 0, 0, 0

	for end := range nums {
		if nums[end] == 0 {
			zeroes++
		}
		for zeroes > k {
			if nums[start] == 0 {
				zeroes--
			}
			start++
		}

		maxLength = end - start + 1

	}

	return maxLength
}
