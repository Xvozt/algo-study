package mergesortedarrays

/*
Даны два отсортированных массива nums1 и nums2 размером m и n соответственно.
Массив nums1 имеет длину m + n, где последние n элементов заполнены нулями и предназначены для размещения элементов из nums2.
Объедините массивы nums1 и nums2 в один отсортированный массив внутри nums1.
Inputs:
nums1 := []int{1,2,3,0,0,0}
m := 3
nums2 := []int{2,5,6}
n := 3

Result:
nums1 = []int{1,2,2,3,5,6}
*/
func mergeSortedArrays(numsOne []int, m int, numsTwo []int, n int) []int {
	p := m + n - 1
	p1 := m - 1
	p2 := n - 1

	for p2 >= 0 {
		if p1 >= 0 && numsOne[p1] > numsTwo[p2] {
			numsOne[p] = numsOne[p1]
			p1--
		} else {
			numsOne[p] = numsTwo[p2]
			p2--
		}

		p--
	}

	return numsOne
}
