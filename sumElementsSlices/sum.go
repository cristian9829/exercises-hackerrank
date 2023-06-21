package main

import "fmt"

func sumElements(slice []int, k int) []int {
	numsSum := []int{}
	for _, numero := range slice {
		for i := 1; i < len(slice); i++ {
			if numero+slice[i] == k {
				numsSum = append(numsSum, numero, slice[i])
			}
		}
		if len(numsSum) >= 2 {
			break
		}
	}
	return numsSum
}

func main() {
	numeros := []int{2, 1, 7}
	response := sumElements(numeros, 8)
	fmt.Println(response) // [1,7]
}
