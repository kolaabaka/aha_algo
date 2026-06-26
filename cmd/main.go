package main

import (
	sort "algogo/algo/sort"
	"fmt"
)

func main() {
	unsorted := []int{9, 9, 1, 1, 3, 4, 5, 2, 3, 4}

	sorted := sort.QuickSort(unsorted)

	fmt.Println(sorted)
}
