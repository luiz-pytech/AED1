package main

func bubbleSort(v []int) []int {
	trocou := true

	for trocou {
		trocou = false

		for i := 0; i < len(v)-1; i++ {
			if v[i] > v[i+1] {
				v[i], v[i+1] = v[i+1], v[i]
				trocou = true
			}
		}
	}

	return v
}
