package main

import "fmt"

func FindMaxMin(array []int, n int) (int, int) {
	var maximum int
	var minimum int
	var i int

	if n%2 == 0 {
		if array[0] > array[1] {
			maximum = array[0]
			minimum = array[1]
		} else {
			maximum = array[1]
			minimum = array[0]
		}

		i = 2
	} else {
		maximum = array[0]
		minimum = array[0]

		i = 1
	}

	for i < n-1 {
		if array[i] > array[i+1] {
			if array[i] > maximum {
				maximum = array[i]
			}

			if array[i+1] < minimum {
				minimum = array[i+1]
			}
		} else {
			if array[i+1] > maximum {
				maximum = array[i+1]
			}

			if array[i] < minimum {
				minimum = array[i]
			}
		}

		i += 2
	}

	return maximum, minimum
}

func main() {
	Array := []int{34, 7, 92, 15, 63, 28, 101, 4, 76, 55, 19, 88, 2, 47, 69, 31, 120, 11, 39, 84}

	max, min := FindMaxMin(Array, len(Array))

	fmt.Println(max, min)
}
