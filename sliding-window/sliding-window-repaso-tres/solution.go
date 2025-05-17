package max_average_subarray

/*
Дан массив целых чисел nums и целое число k.
Найди максимальное среднее значение подмассива длины k.

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75
// Подмассив [12,-5,-6,50] даёт сумму 51 → среднее 12.75
*/
func maxAverageSubarray(nums []int, k int) float64 {
	curSum := 0

	for i := range k {
		curSum += nums[i]
	}

	maxSum := curSum

	for i := k; i < len(nums); i++ {
		curSum -= nums[i-k]
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}

	return float64(maxSum) / float64(k)
}
