package main

func shellSort(data []int) {
	length := len(data)
	gap := int(length / 2)
	for gap >= 1 {
		for i := gap; i < length; i++ {
			value := data[i]
			for j := i; (j-gap) >= 0 && value < data[j-gap]; data[j] = value {
				data[j] = data[j-gap]
				j = j - gap
			}
		}
		gap = int(gap / 2)
	}
}
