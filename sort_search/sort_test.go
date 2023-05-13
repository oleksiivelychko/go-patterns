package sort_search

import (
	"sort"
	"testing"
)

func TestSort_Int(t *testing.T) {
	var unsorted = []int{11, -13, 15, 17, 19, 1, 3, 5, -7, 9}
	sort.Ints(unsorted)

	if !sort.IntsAreSorted(unsorted) {
		t.Error("array is not sorted")
	}

	t.Log(unsorted)
}

func TestSort_String(t *testing.T) {
	var unsorted = []string{"f", "h", "a", "abc", "x", "o", "1", "u", "12", "s"}
	sort.Strings(unsorted)

	if !sort.StringsAreSorted(unsorted) {
		t.Error("array is not sorted")
	}

	t.Log(unsorted)
}
