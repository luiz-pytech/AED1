package main

func linearSearch(val int, v []int) int {
	for i := 0; i < len(v); i++ {
		if v[i] == val {
			return i
		}
	}
	return -1
}
