package containsduplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {

	cases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"positive", []int{9, 9, 5, 9}, true},
		{"negative", []int{1, 2, 3, 4}, false},
		{"empty", nil, false},
		{"non-adjancent positive", []int{9, 5, 3, 9}, true},
		{"single", []int{8}, false},
	}

	for _, testCase := range cases {

		t.Run(testCase.name, func(t *testing.T) {
			got := hasDuplicate(testCase.input)
			if got != testCase.expected {
				t.Errorf("case:%s: expected: %v, got %v",
					testCase.name, testCase.expected, got)
			}
		})

	}

}
