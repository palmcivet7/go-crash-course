package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// // Assign keyvalues
	// emails["Bob"] = "bob@gmail.com"
	// emails["Alice"] = "alice@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and key values
	emails := map[string]string{"Bob":"bob@gmail.com", "Alice":"alice@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
	fmt.Println(len(emails))
}