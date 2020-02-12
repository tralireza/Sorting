package Sorting

import (
	"log"
	"sort"
	"testing"
)

func init() {
	log.Print("> Sort")
}

// 75m Sort Colors
func Test75(t *testing.T) {
	for _, nums := range [][]int{{2, 0, 2, 1, 1, 0}, {2, 0, 1}} {
		sortColors(nums)
		log.Print(" -> ", nums)
	}
}

// 128m Longest Consecutive Sequence
func Test128(t *testing.T) {
	WithSort := func(nums []int) int {
		sort.Ints(nums)

		lcs, cur := 0, 0
		prv := int(1e9 + 1)
		for _, n := range nums {
			if n == prv+1 {
				cur++
				lcs = max(cur, lcs)
			} else {
				cur = 0
			}
			prv = n
		}

		return lcs + 1
	}

	for _, f := range []func([]int) int{longestConsecutive, WithSort} {
		log.Print("==")
		log.Print("4 ?= ", f([]int{100, 4, 200, 1, 3, 2}))
		log.Print("9 ?= ", f([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	}
}

// 2037 Minimum Number of Moves to Sear Everyone
func Test2037(t *testing.T) {
	log.Print("4 ?= ", minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	log.Print("7 ?= ", minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
}
