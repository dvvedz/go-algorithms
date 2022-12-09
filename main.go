package main

import (
	"algo/module01"
	"fmt"
)

func main() {
	// num in list
	var l1 = []int{1, 2, 3, 3, 4}
	fmt.Println("NumInList:", module01.NumInList(l1, 3))

	// sum of list
	var sumList = []int{5, 5, 5, 5}
	fmt.Println("Sum of list:", module01.Sum(sumList))
	fmt.Println("Recursion sum", module01.RecursionSum(sumList))

	// Reverse string
	fmt.Println("Reverse string", module01.RevString("String"))
}
