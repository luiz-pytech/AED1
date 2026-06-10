package main

func mergeSort(array []int) []int {
	//divide, ordena a esquerda e direita e depois junta com o merge.

	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	left := array[:middle]
	right := array[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}
