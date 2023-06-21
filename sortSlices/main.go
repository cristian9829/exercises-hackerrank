package main

import (
	"fmt"
)

func main() {
	slices1 := []int{3, 1, 5}
	slices2 := []int{7, 4, 2}

	slice_ordenado := sortSlices(slices1, slices2)
	fmt.Println(slice_ordenado) //[1 2 3 4 5 7]
}

func sortSlices(slice1, slice2 []int) []int {
	mergeSlice := append(slice1, slice2...)
	for i := 1; i < len(mergeSlice); i++ {
		j := i
		for j > 0 && mergeSlice[j-1] > mergeSlice[j] {
			copySlice := mergeSlice[j]
			mergeSlice[j] = mergeSlice[j-1]
			mergeSlice[j-1] = copySlice
			j--
		}
	}
	return mergeSlice
}
