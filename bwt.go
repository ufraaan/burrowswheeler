package bwt

import "sort"

func Transform(input []byte) ([]byte, int) {
	n := len(input)
	if n == 0 {
		return []byte{}, 0
	}

	rotations := make([]int, n)
	for i := range rotations {
		rotations[i] = i
	}
	sort.Slice(rotations, func(i, j int) bool {
		left, right := rotations[i], rotations[j]
		for k := 0; k < n; k++ {
			a := input[(left+k)%n]
			b := input[(right+k)%n]
			if a != b {
				return a < b
			}
		}
		return left < right
	})

	last := make([]byte, n)
	primary := 0
	for row, start := range rotations {
		last[row] = input[(start+n-1)%n]
		if start == 0 {
			primary = row
		}
	}
	return last, primary
}
