package Sorting

import (
	"fmt"
	"log"
)

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
	// -10^9 <= nums[i] <= 10^9
	const nVal = int(1e9 + 7)

	type DJVal struct{ root, rank int }

	djSet := map[int]*DJVal{}

	AddSet := func(n int) {
		djSet[n] = &DJVal{root: n, rank: 0}
	}
	FindSet := func(n int) int {
		if _, ok := djSet[n]; !ok {
			return nVal
		}

		var Root func(n int) *DJVal
		Root = func(n int) *DJVal {
			N := djSet[n]
			if N.root != n {
				N.root = Root(N.root).root
			}
			return N
		}

		return Root(n).root
	}
	Union := func(x, y int) {
		if FindSet(y) == nVal {
			return
		}
		X, Y := djSet[FindSet(x)], djSet[FindSet(y)]
		if X.rank >= Y.rank {
			Y.root = X.root
			if X.rank == Y.rank {
				X.rank++
			}
		} else {
			X.root = Y.root
		}
	}

	for _, n := range nums {
		AddSet(n)
		Union(n, n+1)
		Union(n, n-1)
	}

	Draw := func(m string) {
		fmt.Print("DJS :: ")
		for n, v := range djSet {
			fmt.Printf("{%d:%d %d} ", n, v.root, v.rank)
		}
		fmt.Printf("/ [%s]\n", m)
	}

	Draw("")

	lcs, fSet, Dup := 0, map[int]int{}, map[int]struct{}{}
	for _, n := range nums {
		if _, ok := Dup[n]; !ok { // [0 0 1 2 2] => [0 1 2] :: lcs:3 !5
			Dup[n] = struct{}{}

			r := FindSet(n)
			fSet[r]++
			lcs = max(fSet[r], lcs)
		}
	}

	Draw("Path Compression")

	log.Print("fSet :: ", fSet)

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
