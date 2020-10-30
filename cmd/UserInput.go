package main 

import "fmt"

func main() {
	fmt.Println("Enter Your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Your name is:" , name)
}

// Use your first-name only. Using both names will trigger an error.
// add another user input for the last name if you want the last name too.