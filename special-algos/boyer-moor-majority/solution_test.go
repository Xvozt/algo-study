package boyermoormajority

/*
🧠 Задача: Majority Element

Дан массив целых чисел nums, в котором всегда существует элемент, встречающийся более чем ⌊n / 2⌋ раз.

Нужно вернуть этот элемент.


Input:  nums = [3,2,3]
Output: 3

Input:  nums = [2,2,1,1,1,2,2]
Output: 2
*/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majority_element_boyer_moore(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		result int
	}{
		{
			name:   "test 1",
			nums:   []int{3, 2, 3},
			result: 3,
		},
		{
			name:   "test 2",
			nums:   []int{2, 2, 1, 1, 1, 2, 2},
			result: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElementBoyerMoore(tt.nums)
			assert.Equal(t, tt.result, got)
		})
	}
}
