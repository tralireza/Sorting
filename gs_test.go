package Sorting

import (
	"log"
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

// 2037 Minimum Number of Moves to Sear Everyone
func Test2037(t *testing.T) {
	log.Print("4 ?= ", minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	log.Print("7 ?= ", minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
}
