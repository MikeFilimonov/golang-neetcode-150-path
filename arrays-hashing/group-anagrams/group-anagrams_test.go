package groupanagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	cases := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{"positive", []string{"act", "pots", "tops", "cat", "stop", "hat"},
			[][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}}},
		{"empty", []string{}, [][]string{}},
		{"onlyAnagrams", []string{"acb", "abc", "cba", "bac"},
			[][]string{{"acb", "abc", "cba", "bac"}}},
		{"noAnagrams", []string{"doge", "eth", "btc", "sol", "xrp"},
			[][]string{{"doge"}, {"eth"}, {"btc"}, {"sol"}, {"xrp"}}},
		{"oneEmptyString", []string{""}, [][]string{{""}}},
		{"twins", []string{"bro", "bro"}, [][]string{{"bro", "bro"}}},
	}

	for _, testCase := range cases {

		t.Run(testCase.name, func(t *testing.T) {

			got := groupAnagrams(testCase.input)
			got = normalize(got)
			testCase.expected = normalize(testCase.expected)
			if !reflect.DeepEqual(got, testCase.expected) {
				t.Errorf("case:%s: expected: %v, got %v",
					testCase.name, testCase.expected, got)
			}
		})

	}

}

func normalize(groups [][]string) [][]string {

	for _, item := range groups {
		sort.Strings(item)
	}

	sort.Slice(groups, func(i, j int) bool {
		return strings.Join(groups[i], ",") < strings.Join(groups[j], ",")
	})

	return groups

}
