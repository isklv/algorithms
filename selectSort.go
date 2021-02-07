package main

func selectSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[min] > data[j] {
				min = j
			}
		}

		if min != i {
			swap(data, i, min)
		}
	}
}
