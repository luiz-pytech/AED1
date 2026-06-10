package main

import rand "math/rand"

func partition(arr []int, low int, high int) int {
	// para aleatoriedade
	randomIndex := low + rand.Intn(high-low+1)
	arr[randomIndex], arr[high] = arr[high], arr[randomIndex]

	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}
