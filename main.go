package main

import (
	"fmt"
	"time"
)

type fn func([]int)

func main() {
	a := 12
	b := 5
	fmt.Println("Swap two int variables: ", a, b)
	a, b = swapTwo(a, b)
	fmt.Println("Swap two int end: ", a, b)

	sortData := generateRadomArray(100000, 0, 10000)

	runSortFn(bubbleSort, "BubbleSort", sortData)

	runSortFn(quickSort, "QuickSort", sortData)

	runSortFn(selectSort, "SelectSort", sortData)

	runSortFn(shellSort, "ShellSort", sortData)

	runSortFn(radixSort, "RadixSort", sortData)

	radixSort(sortData)
	index, found := binarySearch(sortData, 55, 0, len(sortData)-1)
	fmt.Println("BinarySearch", index, found)
}

func runSortFn(f fn, name string, sortData []int) {
	copySortData := make([]int, len(sortData))
	copy(copySortData, sortData)
	start := time.Now()
	fmt.Println(name, "start", "length", len(copySortData))
	f(copySortData)
	// fmt.Println(copySortData)
	duration := time.Since(start)
	fmt.Println(name, "end", duration)
}

//swap two arguments
func swapTwo(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b

	return a, b
}
