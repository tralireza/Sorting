package Sorting

import "log"

func init() {
	log.SetFlags(0)
}

// 75m Sort Colors
func sortColors(nums []int) {
	l, m, h := 0, 0, len(nums)-1

	// [m...h] -> Under Concideration ~ <Not yet sorted>
	// [0...l] -> Visited .L.ow values
	// [h+1...] -> Visited .H.igh values

	for m <= h {
		switch nums[m] {
		case 2:
			nums[h], nums[m] = 2, nums[h]
			h--
		case 1:
			m++
		case 0:
			nums[l], nums[m] = 0, nums[l]
			l++
			m++
		}
	}
}

// 128m Longest Consecutive Sequence
func longestConsecutive(nums []int) int {
	lcs := 0
	return lcs
}

// 2037 Minimum Number of Moves to Sear Everyone
func minMovesToSeat(seats []int, students []int) int {
	// 1 <= seats[i], students[i] <= 100
	var fP []int

	for _, Pos := range [][]int{seats, students} {
		fP = make([]int, 100+1)
		for _, p := range Pos {
			fP[p]++
		}

		i := 0
		for p, f := range fP {
			for range f {
				Pos[i] = p
				i++
			}
		}
	}

	moves := 0
	for i := 0; i < len(students) && i < len(seats); i++ {
		m := students[i] - seats[i]
		if m < 0 {
			m *= -1
		}
		moves += m
	}
	return moves
}
