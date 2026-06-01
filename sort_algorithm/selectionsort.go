package main

func selectionSort(v []int) []int {
	for i := 0; i < len(v); i++ {
		min_index := i
		for j := i + 1; j <= len(v)-1; j++ {
			if v[j] < v[min_index] {
				min_index = j
			}
		}
		v[i], v[min_index] = v[min_index], v[i]
	}

	return v

}
