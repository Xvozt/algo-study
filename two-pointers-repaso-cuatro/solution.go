package valid_palindrome

import (
	"unicode"
)

func validPalindrome(text string) bool {
	cleaned := []rune{}

	for _, v := range []rune(text) {
		if unicode.IsLetter(v) {
			cleaned = append(cleaned, unicode.ToLower(v))
		} else if unicode.IsDigit(v) {
			cleaned = append(cleaned, v)
		}
	}

	left, right := 0, len(cleaned)-1

	for left <= right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}
