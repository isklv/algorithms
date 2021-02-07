package main

//bubble sort
func bubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for y := i + 1; y < len(data); y++ {
			if data[y] < data[i] {
				swap(data, i, y)
			}
		}
	}
}
