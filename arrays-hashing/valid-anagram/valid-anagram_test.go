package validanagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {

	cases := []struct {
		name     string
		op1      string
		op2      string
		expected bool
	}{
		{"positive", "hacker", "hckare", true},
		{"negativeEqual", "hacker", "badger", false},
		{"empty", "f", "long", false},
		{"random items", "texas", "montgolfier", false},
	}

	for _, testCase := range cases {

		t.Run(testCase.name, func(t *testing.T) {
			got := isAnagram(testCase.op1, testCase.op2)
			if got != testCase.expected {
				t.Errorf("case:%s: expected: %v, got %v",
					testCase.name, testCase.expected, got)
			}
		})

	}

}
