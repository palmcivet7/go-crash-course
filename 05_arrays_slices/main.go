package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	// Declare and assign
	// fruitArr := [2]string{"apple", "orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"apple", "orange", "grape", "cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}