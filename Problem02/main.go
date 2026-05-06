package main

import "fmt"

func SumPairs(array []int, n int, value int) (int, int) {
	numbers := make(map[int]int)

	for i := 0; i < n; i++ {
		valuePair := value - array[i]
		val, exist := numbers[valuePair]
		if exist && val != i {
			return i, val
		}
		numbers[array[i]] = i
	}

	return -1, -1
}

func main() {

	array := []int{7, 3, 10, 2, 5, 14, 7, 1, 9, 6}
	value := 15

	i, j := SumPairs(array, len(array), value)

	fmt.Println("[", i, ",", j, "]")
}
