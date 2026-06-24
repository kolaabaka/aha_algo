package algo

// Sort array with block algorithm, count each element and his index,
// could use with arrays with value more than length array, but not interested
//
// Parameters:
//   - unsorted: slice of integers to be sorted
//
// Returns:
//   - []int: sorted slice
//
// Example:
//   arr := []int{3, 1, 4, 1, 5}
//   sorted := BlockSort(arr)
//   fmt.Println(sorted) // [1 1 3 4 5]
func BlockSort(unsorted []int) []int {
	leng := len(unsorted)

	blockSlice := make([]int, leng)

	blockResult := make([]int, 0)

	for _, value := range unsorted {
		blockSlice[value]++
	}

	for index, value := range blockSlice {
		for range value {
			blockResult = append(blockResult, index)
		}
	}

	return blockResult
}

// Sort array with bubble algorithm, step by step, index by index
// swap elemnent by element
//
// Parameters:
//   - unsorted: slice of integers to be sorted
//
// Returns:
//   - []int: sorted slice
//
// Example:
//   arr := []int{3, 1, 4, 1, 5}
//   sorted := BlockSort(arr)
//   fmt.Println(sorted) // [1 1 3 4 5]
func BubbleSort(unsorted []int) []int {

	for i := 0; i < len(unsorted); i++ {
		for j := 1; j < len(unsorted)-i; j++ {
			if unsorted[j] < unsorted[j-1] {
				swap(unsorted, j, j-1)
			}
		}
	}

	return unsorted
}

func QuickSort(unsorted []int) []int {
	return quickSortHelper(unsorted, 0, len(unsorted)-1)
}

func quickSortHelper(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	pivot := arr[left]

	i := left
	j := right

	for i != j {
		for i < j && arr[j] >= pivot {
			j--
		}

		for i < j && arr[i] <= pivot {
			i++
		}

		if i < j {
			swap(arr, i, j)
		}
	}

	arr[left] = arr[i]
	arr[i] = pivot

	quickSortHelper(arr, left, i-1)
	quickSortHelper(arr, i+1, right)

	return arr
}

func swap(arr []int, first, second int) {
	temp := arr[first]
	arr[first] = arr[second]
	arr[second] = temp
}
