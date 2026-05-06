package main

import "fmt"

func SumPairs(array []int, n int, value int) (int, int) {

	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if array[i]+array[j] == value {
				return i, j
			}
		}
	}

	return -1, -1
}

func main() {

	array := []int{8, 3, 10, 2, 5, 14, 7, 1, 9, 6}
	value := 10

	i, j := SumPairs(array, len(array), value)

	fmt.Println("[", i, ",", j, "]")
}
