package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(linearSearch(5, v))
	fmt.Println(binarySearch(5, v))

	fmt.Println(linearSearch(11, v))
	fmt.Println(binarySearch(11, v))
}
