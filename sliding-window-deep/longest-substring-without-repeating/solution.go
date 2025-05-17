package longestsubstringwithoutrepeating

/*
Дана строка s.
Найди длину самой длинной подстроки, в которой все символы уникальны (без повторений).

Input: "abcabcbb"
Output: 3
Explanation: "abc"
*/

func longestSubstringWithoutRepeating(str string) int {
	byteMap := make(map[byte]int)

	left := 0
	curLen, maxLen := 0, 0

	for right := range len(str) {
		if _, ok := byteMap[str[right]]; ok {
			left = max(left, byteMap[str[right]]+1)
		}

		byteMap[str[right]] = right

		curLen = right - left + 1
		maxLen = max(curLen, maxLen)

	}

	return maxLen
}
