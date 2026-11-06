package bwt

import (
	"bytes"
	"testing"
)

func TestTransformBanana(t *testing.T) {
	last, primary := Transform([]byte("banana"))
	if !bytes.Equal(last, []byte("nnbaaa")) {
		t.Fatalf("Transform() = %q, want %q", last, "nnbaaa")
	}
	if primary != 3 {
		t.Fatalf("Transform() primary = %d, want 3", primary)
	}
}

func TestTransformInverseRoundTrip(t *testing.T) {
	tests := [][]byte{
		{},
		[]byte("a"),
		[]byte("aaaaaa"),
		[]byte("banana"),
		[]byte("ufraan"),
		[]byte("mississippi"),
		{0, 255, 1, 0, 128, 255},
	}

	for _, input := range tests {
		last, primary := Transform(input)
		got, err := InverseTransform(last, primary)
		if err != nil {
			t.Fatalf("InverseTransform(%v, %d) error = %v", last, primary, err)
		}
		if !bytes.Equal(got, input) {
			t.Errorf("round trip for %v = %v, want %v", input, got, input)
		}
	}
}

func TestInverseTransformInvalidPrimary(t *testing.T) {
	tests := []struct {
		last    []byte
		primary int
	}{
		{last: []byte("a"), primary: -1},
		{last: []byte("a"), primary: 1},
		{last: []byte{}, primary: 1},
	}

	for _, test := range tests {
		if _, err := InverseTransform(test.last, test.primary); err != ErrInvalidPrimaryIndex {
			t.Errorf("InverseTransform(%v, %d) error = %v, want %v", test.last, test.primary, err, ErrInvalidPrimaryIndex)
		}
	}
}
