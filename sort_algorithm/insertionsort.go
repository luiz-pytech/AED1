package main

func insertionSort(v []int) []int {

	for i := 1; i < len(v); i++ {
		pivo := v[i]
		j := i - 1

		for j >= 0 && v[j] > pivo {
			v[j+1] = v[j]
			j--
		}
		v[j+1] = pivo

	}

	return v
}
