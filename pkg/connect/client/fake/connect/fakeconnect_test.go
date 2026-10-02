package fakeconnect

import (
	"slices"
	"testing"
)

func TestSortedValues(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}

	// Run repeatedly, since an unsorted implementation could pass by chance.
	for range 20 {
		got := slices.Collect(SortedValues(m))
		if want := []int{1, 2, 3}; !slices.Equal(got, want) {
			t.Fatalf("SortedValues() = %v, want %v", got, want)
		}
	}
}

func TestSortedValues_earlyBreak(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}

	var got []int
	for v := range SortedValues(m) {
		got = append(got, v)
		if len(got) == 2 {
			break
		}
	}
	if want := []int{1, 2}; !slices.Equal(got, want) {
		t.Fatalf("SortedValues() with break = %v, want %v", got, want)
	}
}
