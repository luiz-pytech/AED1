package main

// vector is sorted ascending
func binarySearch(val int, v []int) int {
	low := 0
	high := len(v) - 1

	for low <= high {
		mid := (low + high) / 2

		if v[mid] == val {
			return mid
		} else if val < v[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
