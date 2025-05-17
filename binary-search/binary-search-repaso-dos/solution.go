package is_bad_version

/**
Дан отсортированный массив букв letters и символ target.
Найди первую букву, которая строго больше target.
Если такой нет — верни первую букву массива (wrap around).

*/

func smallestLetterGreaterThanTarget(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	for left <= right {
		mid := left + (right-left)/2

		if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left == len(letters) {
		return letters[0]
	}

	return letters[left]
}
