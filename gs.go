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
