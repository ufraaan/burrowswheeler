package bwt

import (
	"errors"
	"sort"
)

var ErrInvalidPrimaryIndex = errors.New("invalid primary index")

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

func InverseTransform(lastColumn []byte, primary int) ([]byte, error) {
	n := len(lastColumn)
	if n == 0 {
		if primary != 0 {
			return nil, ErrInvalidPrimaryIndex
		}
		return []byte{}, nil
	}
	if primary < 0 || primary >= n {
		return nil, ErrInvalidPrimaryIndex
	}

	occurrences := make([]int, n)
	counts := [256]int{}
	for i, value := range lastColumn {
		occurrences[i] = counts[value]
		counts[value]++
	}

	starts := [256]int{}
	for value := 1; value < len(starts); value++ {
		starts[value] = starts[value-1] + counts[value-1]
	}

	result := make([]byte, n)
	row := primary
	for i := n - 1; i >= 0; i-- {
		value := lastColumn[row]
		result[i] = value
		row = starts[value] + occurrences[row]
	}
	return result, nil
}
