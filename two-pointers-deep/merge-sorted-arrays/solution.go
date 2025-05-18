package mergesortedarrays

/*

Дан отсортированный массив nums1 длиной m + n,
в котором первые m элементов заполнены, а последние n — нули.

Также дан второй отсортированный массив nums2 длиной n.

Нужно объединить nums2 в nums1, чтобы результат был отсортирован.
Сделать это на месте (в nums1), не создавая новый массив.

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	p := m + n - 1
	p1 := m - 1
	p2 := n - 1

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	return nums1
}
