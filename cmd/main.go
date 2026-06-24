package main

import (
	algo "algogo/algo"
	"fmt"
)

func main() {
	unsorted := []int{9, 9, 1, 1, 3, 4, 5, 2, 3, 4}

	sorted := algo.QuickSort(unsorted)

	fmt.Println(sorted)
}
