package sort_search

import (
	"sort"
	"testing"
)

func TestSearch_Int(t *testing.T) {
	var sorted = []int{-13, -7, 1, 3, 5, 9, 11, 15, 17, 19}
	index := sort.SearchInts(sorted, 17)

	if index != 8 {
		t.Errorf("index %d is wrong", index)
	}
}

func TestSearch_String(t *testing.T) {
	var sorted = []string{"1", "12", "a", "abc", "f", "h", "o", "s", "u", "x"}
	index := sort.SearchStrings(sorted, "abc")

	if index != 3 {
		t.Errorf("index %d is wrong", index)
	}
}
