package algo_test

import (
	sort "algogo/algo/sort"
	"slices"
	"testing"
)

var unsorted = []int{9, 9, 1, 1, 3, 4, 5, 2, 3, 4}
var sorted = []int{1, 1, 2, 3, 3, 4, 4, 5, 9, 9}

func TestBlock(t *testing.T) {

	sortedBlock := sort.BlockSort(unsorted)

	if !slices.Equal(sortedBlock, sorted) {
		t.Errorf("Doesn`t equal %v and %v", sortedBlock, sorted)
	}
}

func TestBubble(t *testing.T) {
	sortedBlock := sort.BubbleSort(unsorted)

	if !slices.Equal(sortedBlock, sorted) {
		t.Errorf("Doesn`t equal %v and %v", sortedBlock, sorted)
	}
}

func TestQuick(t *testing.T) {
	sortedBlock := sort.QuickSort(unsorted)

	if !slices.Equal(sortedBlock, sorted) {
		t.Errorf("Doesn`t equal %v and %v", sortedBlock, sorted)
	}
}
