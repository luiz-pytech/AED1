package main

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	maior := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maior {
			maior = arr[i]
		}
	}

	count := make([]int, maior+1)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	sortedIndex := 0
	for i := 0; i < len(count); i++ {
		for count[i] > 0 {
			arr[sortedIndex] = i
			sortedIndex++
			count[i]--
		}
	}
	return arr
}
