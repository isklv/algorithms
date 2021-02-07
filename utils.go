package main

import "math/rand"

func swap(data []int, i, y int) {
	data[i], data[y] = data[y], data[i]
}

func generateRadomArray(length int, min, max int) []int {
	data := make([]int, 0)
	for i := 0; i < length; i++ {
		data = append(data, randomInteger(min, max))
	}
	return data
}

func randomInteger(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
