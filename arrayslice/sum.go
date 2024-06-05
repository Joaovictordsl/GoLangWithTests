package main

func Sum(array []int) int {

	var result int

	for _, number := range array {
		result += number
	}

	return result
}
