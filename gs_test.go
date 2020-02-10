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
