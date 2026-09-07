package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {

	cases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"basic", []int{1, 8, 6, 9}, 10, []int{0, 3}},
		{"twins", []int{9, 8, 6, 9}, 18, []int{0, 3}},
		{"empty", []int{}, 9, []int{}},
		{"one-liner", []int{9}, 9, []int{}},
		{"trio", []int{9, 9, 9}, 18, []int{0, 1}},
	}

	for _, testCase := range cases {

		t.Run(testCase.name, func(t *testing.T) {
			got := twoSum(testCase.nums, testCase.target)
			if !slices.Equal(got, testCase.expected) {
				t.Errorf("case:%s: expected: %v, got %v",
					testCase.name, testCase.expected, got)
			}
		})

	}

}
