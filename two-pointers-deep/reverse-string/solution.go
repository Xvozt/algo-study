package reversestring

/*
Дан массив символов s — разверни его на месте.
Input:  ['h','e','l','l','o']
Output: ['o','l','l','e','h']
*/
func reverseString(str []byte) []byte {

	left, right := 0, len(str)-1

	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}

	return str
}
