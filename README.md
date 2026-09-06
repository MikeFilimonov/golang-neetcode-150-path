 NeetCode 150 — Go

Solutions to the [NeetCode 150](https://neetcode.io/practice) in Go, one package per
problem, organized by the roadmap's topic order.

**Solved: 0 / 150**

## Layout

```
.
├── go.mod                       # module github.com/mike/neetcode
├── Makefile
├── <category>/
│   └── <NNNN>-<slug>/
│       ├── solution.go          # package <slug without hyphens>
│       ├── solution_test.go     # same package, table-driven
│       └── README.md            # problem link + approach, complexity
```

Categories follow the NeetCode roadmap: `arrays-hashing`, `two-pointers`,
`sliding-window`, `stack`, `binary-search`, `linked-list`, `trees`, `tries`, `heap`,
`backtracking`, `graphs`, `advanced-graphs`, `1d-dp`, `2d-dp`, `greedy`, `intervals`,
`math-geometry`, `bit-manipulation`.

## Conventions

- **One package per problem.** Directory is `two-sum`; package is `twosum`
  (Go identifiers can't contain hyphens or start with a digit — the directory name
  doesn't need to match the package name).
- **Keep LeetCode's signature.** Use the unexported function name LeetCode gives you
  (`func twoSum(...)`), so files paste in and out of the editor unchanged.
- **Define `ListNode` / `TreeNode` locally** in each package that needs them. The
  duplication keeps every solution self-contained and paste-ready.
- **Table-driven tests** in the same package so they can reach unexported helpers.
- **Per-problem README** records the LeetCode link, the idea in a sentence or two,
  and time/space complexity.

## Adding a problem

```sh
mkdir -p arrays-hashing/two-sum
cd arrays-hashing/two-sum
$EDITOR solution.go solution_test.go README.md
go test ./...
```

`solution.go`:

```go
package twosum

func twoSum(nums []int, target int) []int {
}
```

`solution_test.go`:

```go
package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
```

## Running

```sh
make test      # go test ./...
make lint      # go vet ./... + gofmt check
go test ./arrays-hashing/...        # one category
go test -run TwoSum ./...           # one problem
```

# NeetCode 150 — Go

Solutions to the [NeetCode 150](https://neetcode.io/practice) in Go, one package per
problem, organized by the roadmap's topic order.

**Solved: 0 / 150**

## Layout

```
.
├── go.mod                    # module github.com/mike/neetcode
├── Makefile
├── <category>/
│   └── <slug>/
│       ├── solution.go       # package <slug without hyphens>
│       ├── solution_test.go  # same package, table-driven
│       └── README.md         # problem link + approach, complexity
```

Categories follow the NeetCode roadmap: `arrays-hashing`, `two-pointers`,
`sliding-window`, `stack`, `binary-search`, `linked-list`, `trees`, `tries`, `heap`,
`backtracking`, `graphs`, `advanced-graphs`, `1d-dp`, `2d-dp`, `greedy`, `intervals`,
`math-geometry`, `bit-manipulation`. Within a category, the checklist below is in
roadmap order; on disk the folders sort alphabetically.

## Conventions

- **One package per problem.** Directory is `two-sum`; package is `twosum`
  (Go identifiers can't contain hyphens — the directory name doesn't need to match
  the package name).
- **Keep LeetCode's signature.** Use the unexported function name LeetCode gives you
  (`func twoSum(...)`), so files paste in and out of the editor unchanged.
- **Define `ListNode` / `TreeNode` locally** in each package that needs them. The
  duplication keeps every solution self-contained and paste-ready.
- **Table-driven tests** in the same package so they can reach unexported helpers.
- **Per-problem README** records the LeetCode link, the idea in a sentence or two,
  and time/space complexity.

## Adding a problem

```sh
mkdir -p arrays-hashing/two-sum
cd arrays-hashing/two-sum
$EDITOR solution.go solution_test.go README.md
go test ./...
```

`solution.go`:

```go
package twosum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
```

`solution_test.go`:

```go
package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
```

## Running

```sh
make test      # go test ./...
make lint      # go vet ./... + gofmt check
go test ./arrays-hashing/...        # one category
go test -run TwoSum ./...           # one problem
```

Suggested `Makefile`:

```makefile
.PHONY: test lint fmt
test:
	go test ./...
lint:
	go vet ./...
	test -z "$$(gofmt -l .)"
fmt:
	gofmt -w .
```

## Progress

### Arrays & Hashing
- [ ] [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) — `arrays-hashing/contains-duplicate`
- [ ] [Valid Anagram](https://leetcode.com/problems/valid-anagram/) — `arrays-hashing/valid-anagram`
- [ ] [Two Sum](https://leetcode.com/problems/two-sum/) — `arrays-hashing/two-sum`
- [ ] [Group Anagrams](https://leetcode.com/problems/group-anagrams/) — `arrays-hashing/group-anagrams`
- [ ] [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) — `arrays-hashing/top-k-frequent-elements`
- [ ] [Encode and Decode Strings](https://leetcode.com/problems/encode-and-decode-strings/) — `arrays-hashing/encode-and-decode-strings`
- [ ] [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) — `arrays-hashing/product-of-array-except-self`
- [ ] [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) — `arrays-hashing/valid-sudoku`
- [ ] [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/) — `arrays-hashing/longest-consecutive-sequence`

### Two Pointers
- [ ] [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) — `two-pointers/valid-palindrome`
- [ ] [Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) — `two-pointers/two-sum-ii`
- [ ] [3Sum](https://leetcode.com/problems/3sum/) — `two-pointers/3sum`
- [ ] [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) — `two-pointers/container-with-most-water`
- [ ] [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) — `two-pointers/trapping-rain-water`

### Sliding Window
- [ ] [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) — `sliding-window/best-time-to-buy-and-sell-stock`
- [ ] [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) — `sliding-window/longest-substring-without-repeating-characters`
- [ ] [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) — `sliding-window/longest-repeating-character-replacement`
- [ ] [Permutation in String](https://leetcode.com/problems/permutation-in-string/) — `sliding-window/permutation-in-string`
- [ ] [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) — `sliding-window/minimum-window-substring`
- [ ] [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) — `sliding-window/sliding-window-maximum`

### Stack
- [ ] [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) — `stack/valid-parentheses`
- [ ] [Min Stack](https://leetcode.com/problems/min-stack/) — `stack/min-stack`
- [ ] [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) — `stack/evaluate-reverse-polish-notation`
- [ ] [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) — `stack/generate-parentheses`
- [ ] [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/) — `stack/daily-temperatures`
- [ ] [Car Fleet](https://leetcode.com/problems/car-fleet/) — `stack/car-fleet`
- [ ] [Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) — `stack/largest-rectangle-in-histogram`

### Binary Search
- [ ] [Binary Search](https://leetcode.com/problems/binary-search/) — `binary-search/binary-search`
- [ ] [Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/) — `binary-search/search-a-2d-matrix`
- [ ] [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/) — `binary-search/koko-eating-bananas`
- [ ] [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) — `binary-search/find-minimum-in-rotated-sorted-array`
- [ ] [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) — `binary-search/search-in-rotated-sorted-array`
- [ ] [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) — `binary-search/time-based-key-value-store`
- [ ] [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) — `binary-search/median-of-two-sorted-arrays`

### Linked List
- [ ] [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) — `linked-list/reverse-linked-list`
- [ ] [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) — `linked-list/merge-two-sorted-lists`
- [ ] [Reorder List](https://leetcode.com/problems/reorder-list/) — `linked-list/reorder-list`
- [ ] [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) — `linked-list/remove-nth-node-from-end-of-list`
- [ ] [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/) — `linked-list/copy-list-with-random-pointer`
- [ ] [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) — `linked-list/add-two-numbers`
- [ ] [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) — `linked-list/linked-list-cycle`
- [ ] [Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/) — `linked-list/find-the-duplicate-number`
- [ ] [LRU Cache](https://leetcode.com/problems/lru-cache/) — `linked-list/lru-cache`
- [ ] [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) — `linked-list/merge-k-sorted-lists`
- [ ] [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) — `linked-list/reverse-nodes-in-k-group`

### Trees
- [ ] [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) — `trees/invert-binary-tree`
- [ ] [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) — `trees/maximum-depth-of-binary-tree`
- [ ] [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) — `trees/diameter-of-binary-tree`
- [ ] [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) — `trees/balanced-binary-tree`
- [ ] [Same Tree](https://leetcode.com/problems/same-tree/) — `trees/same-tree`
- [ ] [Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/) — `trees/subtree-of-another-tree`
- [ ] [Lowest Common Ancestor of a BST](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) — `trees/lowest-common-ancestor-of-a-bst`
- [ ] [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) — `trees/binary-tree-level-order-traversal`
- [ ] [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/) — `trees/binary-tree-right-side-view`
- [ ] [Count Good Nodes in Binary Tree](https://leetcode.com/problems/count-good-nodes-in-binary-tree/) — `trees/count-good-nodes-in-binary-tree`
- [ ] [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) — `trees/validate-binary-search-tree`
- [ ] [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) — `trees/kth-smallest-element-in-a-bst`
- [ ] [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) — `trees/construct-binary-tree-from-preorder-and-inorder-traversal`
- [ ] [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) — `trees/binary-tree-maximum-path-sum`
- [ ] [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) — `trees/serialize-and-deserialize-binary-tree`

### Tries
- [ ] [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) — `tries/implement-trie-prefix-tree`
- [ ] [Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/) — `tries/design-add-and-search-words-data-structure`
- [ ] [Word Search II](https://leetcode.com/problems/word-search-ii/) — `tries/word-search-ii`

### Heap / Priority Queue
- [ ] [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) — `heap/kth-largest-element-in-a-stream`
- [ ] [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/) — `heap/last-stone-weight`
- [ ] [K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/) — `heap/k-closest-points-to-origin`
- [ ] [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) — `heap/kth-largest-element-in-an-array`
- [ ] [Task Scheduler](https://leetcode.com/problems/task-scheduler/) — `heap/task-scheduler`
- [ ] [Design Twitter](https://leetcode.com/problems/design-twitter/) — `heap/design-twitter`
- [ ] [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) — `heap/find-median-from-data-stream`

### Backtracking
- [ ] [Subsets](https://leetcode.com/problems/subsets/) — `backtracking/subsets`
- [ ] [Combination Sum](https://leetcode.com/problems/combination-sum/) — `backtracking/combination-sum`
- [ ] [Permutations](https://leetcode.com/problems/permutations/) — `backtracking/permutations`
- [ ] [Subsets II](https://leetcode.com/problems/subsets-ii/) — `backtracking/subsets-ii`
- [ ] [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/) — `backtracking/combination-sum-ii`
- [ ] [Word Search](https://leetcode.com/problems/word-search/) — `backtracking/word-search`
- [ ] [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/) — `backtracking/palindrome-partitioning`
- [ ] [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) — `backtracking/letter-combinations-of-a-phone-number`
- [ ] [N-Queens](https://leetcode.com/problems/n-queens/) — `backtracking/n-queens`

### Graphs
- [ ] [Number of Islands](https://leetcode.com/problems/number-of-islands/) — `graphs/number-of-islands`
- [ ] [Max Area of Island](https://leetcode.com/problems/max-area-of-island/) — `graphs/max-area-of-island`
- [ ] [Clone Graph](https://leetcode.com/problems/clone-graph/) — `graphs/clone-graph`
- [ ] [Walls and Gates](https://leetcode.com/problems/walls-and-gates/) — `graphs/walls-and-gates`
- [ ] [Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) — `graphs/rotting-oranges`
- [ ] [Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/) — `graphs/pacific-atlantic-water-flow`
- [ ] [Surrounded Regions](https://leetcode.com/problems/surrounded-regions/) — `graphs/surrounded-regions`
- [ ] [Course Schedule](https://leetcode.com/problems/course-schedule/) — `graphs/course-schedule`
- [ ] [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) — `graphs/course-schedule-ii`
- [ ] [Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/) — `graphs/graph-valid-tree`
- [ ] [Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/) — `graphs/number-of-connected-components-in-an-undirected-graph`
- [ ] [Redundant Connection](https://leetcode.com/problems/redundant-connection/) — `graphs/redundant-connection`
- [ ] [Word Ladder](https://leetcode.com/problems/word-ladder/) — `graphs/word-ladder`

### Advanced Graphs
- [ ] [Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/) — `advanced-graphs/reconstruct-itinerary`
- [ ] [Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/) — `advanced-graphs/min-cost-to-connect-all-points`
- [ ] [Network Delay Time](https://leetcode.com/problems/network-delay-time/) — `advanced-graphs/network-delay-time`
- [ ] [Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/) — `advanced-graphs/swim-in-rising-water`
- [ ] [Alien Dictionary](https://leetcode.com/problems/alien-dictionary/) — `advanced-graphs/alien-dictionary`
- [ ] [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/) — `advanced-graphs/cheapest-flights-within-k-stops`

### 1-D Dynamic Programming
- [ ] [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) — `1d-dp/climbing-stairs`
- [ ] [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) — `1d-dp/min-cost-climbing-stairs`
- [ ] [House Robber](https://leetcode.com/problems/house-robber/) — `1d-dp/house-robber`
- [ ] [House Robber II](https://leetcode.com/problems/house-robber-ii/) — `1d-dp/house-robber-ii`
- [ ] [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) — `1d-dp/longest-palindromic-substring`
- [ ] [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) — `1d-dp/palindromic-substrings`
- [ ] [Decode Ways](https://leetcode.com/problems/decode-ways/) — `1d-dp/decode-ways`
- [ ] [Coin Change](https://leetcode.com/problems/coin-change/) — `1d-dp/coin-change`
- [ ] [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/) — `1d-dp/maximum-product-subarray`
- [ ] [Word Break](https://leetcode.com/problems/word-break/) — `1d-dp/word-break`
- [ ] [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) — `1d-dp/longest-increasing-subsequence`
- [ ] [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/) — `1d-dp/partition-equal-subset-sum`

### 2-D Dynamic Programming
- [ ] [Unique Paths](https://leetcode.com/problems/unique-paths/) — `2d-dp/unique-paths`
- [ ] [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) — `2d-dp/longest-common-subsequence`
- [ ] [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) — `2d-dp/best-time-to-buy-and-sell-stock-with-cooldown`
- [ ] [Coin Change II](https://leetcode.com/problems/coin-change-ii/) — `2d-dp/coin-change-ii`
- [ ] [Target Sum](https://leetcode.com/problems/target-sum/) — `2d-dp/target-sum`
- [ ] [Interleaving String](https://leetcode.com/problems/interleaving-string/) — `2d-dp/interleaving-string`
- [ ] [Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) — `2d-dp/longest-increasing-path-in-a-matrix`
- [ ] [Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/) — `2d-dp/distinct-subsequences`
- [ ] [Edit Distance](https://leetcode.com/problems/edit-distance/) — `2d-dp/edit-distance`
- [ ] [Burst Balloons](https://leetcode.com/problems/burst-balloons/) — `2d-dp/burst-balloons`
- [ ] [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) — `2d-dp/regular-expression-matching`

### Greedy
- [ ] [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) — `greedy/maximum-subarray`
- [ ] [Jump Game](https://leetcode.com/problems/jump-game/) — `greedy/jump-game`
- [ ] [Jump Game II](https://leetcode.com/problems/jump-game-ii/) — `greedy/jump-game-ii`
- [ ] [Gas Station](https://leetcode.com/problems/gas-station/) — `greedy/gas-station`
- [ ] [Hand of Straights](https://leetcode.com/problems/hand-of-straights/) — `greedy/hand-of-straights`
- [ ] [Merge Triplets to Form Target Triplet](https://leetcode.com/problems/merge-triplets-to-form-target-triplet/) — `greedy/merge-triplets-to-form-target-triplet`
- [ ] [Partition Labels](https://leetcode.com/problems/partition-labels/) — `greedy/partition-labels`
- [ ] [Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/) — `greedy/valid-parenthesis-string`

### Intervals
- [ ] [Insert Interval](https://leetcode.com/problems/insert-interval/) — `intervals/insert-interval`
- [ ] [Merge Intervals](https://leetcode.com/problems/merge-intervals/) — `intervals/merge-intervals`
- [ ] [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) — `intervals/non-overlapping-intervals`
- [ ] [Meeting Rooms](https://leetcode.com/problems/meeting-rooms/) — `intervals/meeting-rooms`
- [ ] [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/) — `intervals/meeting-rooms-ii`
- [ ] [Minimum Interval to Include Each Query](https://leetcode.com/problems/minimum-interval-to-include-each-query/) — `intervals/minimum-interval-to-include-each-query`

### Math & Geometry
- [ ] [Rotate Image](https://leetcode.com/problems/rotate-image/) — `math-geometry/rotate-image`
- [ ] [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/) — `math-geometry/spiral-matrix`
- [ ] [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/) — `math-geometry/set-matrix-zeroes`
- [ ] [Happy Number](https://leetcode.com/problems/happy-number/) — `math-geometry/happy-number`
- [ ] [Plus One](https://leetcode.com/problems/plus-one/) — `math-geometry/plus-one`
- [ ] [Pow(x, n)](https://leetcode.com/problems/powx-n/) — `math-geometry/powx-n`
- [ ] [Multiply Strings](https://leetcode.com/problems/multiply-strings/) — `math-geometry/multiply-strings`
- [ ] [Detect Squares](https://leetcode.com/problems/detect-squares/) — `math-geometry/detect-squares`

### Bit Manipulation
- [ ] [Single Number](https://leetcode.com/problems/single-number/) — `bit-manipulation/single-number`
- [ ] [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) — `bit-manipulation/number-of-1-bits`
- [ ] [Counting Bits](https://leetcode.com/problems/counting-bits/) — `bit-manipulation/counting-bits`
- [ ] [Reverse Bits](https://leetcode.com/problems/reverse-bits/) — `bit-manipulation/reverse-bits`
- [ ] [Missing Number](https://leetcode.com/problems/missing-number/) — `bit-manipulation/missing-number`
- [ ] [Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/) — `bit-manipulation/sum-of-two-integers`
- [ ] [Reverse Integer](https://leetcode.com/problems/reverse-integer/) — `bit-manipulation/reverse-integer`
