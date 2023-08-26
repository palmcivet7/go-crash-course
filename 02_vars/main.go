package main

import "fmt"

func main() {
	var age = 33
	var isCool = true
	isCool = false
	size := 1.3

	name, email := "palmcivet", "palmcivet@proton.me"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}