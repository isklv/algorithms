package main

func binarySearch(data []int, target int, low int, high int) (index int, found bool) {
	mid := (high + low) / 2
	if low > high {
		return -1, false
	}

	if target < data[mid] {
		return binarySearch(data, target, low, mid-1)
	} else if target > data[mid] {
		return binarySearch(data, target, mid+1, high)
	} else {
		return mid, true
	}
}
