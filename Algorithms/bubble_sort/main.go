package main

import (
	"fmt"
	"reflect"
)

func sweep(numbers []int) []int {
	//Swap using reflect
	swapF := reflect.Swapper(numbers)
	endFlag := false

	for i := 0; i+1 < len(numbers); i++ {
		for j := 0; j+1 < len(numbers)-i; j++ {
			if numbers[j] > numbers[j+1] {
				swapF(j, j+1)
				endFlag = true
			}
		}
		if endFlag == false {
			break
		}
		fmt.Printf("Sweep %d result: %v \n", i, numbers)
		endFlag = false
	}
	return numbers
}

func main() {
	n_slice := []int{11, 11, 11, 11}
	fmt.Printf("Original array is: %v \n", n_slice)

	arranged := sweep(n_slice)
	fmt.Printf("Resulting array is %v \n", arranged)
}
