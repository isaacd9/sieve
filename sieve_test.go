package sieve

import (
	"testing"
)

func TestTo(t *testing.T) {
	tests := []struct {
		test   uint32
		result []uint32
	}{
		{
			test:   3,
			result: []uint32{2},
		},
		{
			test:   10,
			result: []uint32{2, 3, 5, 7},
		},
		{
			test:   20,
			result: []uint32{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}

	for _, test := range tests {
		got := To(test.test)

		for i, want := range test.result {
			if got[i] != want {
				t.Fatalf("Test %v, Expected %v, got %v", i, test.test, test.result)
			}
		}
	}
}
