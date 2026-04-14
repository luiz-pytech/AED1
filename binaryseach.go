package main

import "fmt"

func linearSearch(val int, v []int) int {
    for i := 0; i < len(v); i++ {
        if v[i] == val {
            return i
        }
    }
    return -1
}

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

func main() {
    arr1 := [10]int{1,2,3,4,5,6,7,8,9,10}

    fmt.Println(linearSearch(11, arr1[:]))
    fmt.Println(binarySearch(5, arr1[:]))
}
