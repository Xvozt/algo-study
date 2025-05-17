package longestrepeatingcharreplacement

func longestRepeatingCharReplacement(str string, k int) int {
	freqMap := make(map[byte]int)

	start, maxLength := 0, 0

	for end := range len(str) {
		if freqMap[str[end]] > k {
			start++
		} else {
			freqMap[str[end]]++
			maxLength = max(maxLength, end-start+1)
		}

	}

	return maxLength
}
