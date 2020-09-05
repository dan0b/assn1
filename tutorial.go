package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hey lets get you started! \n First we'll need some information.")

	fmt.Println("\n What is your first name?")
	var first string
	fmt.Scanln(&first)

	fmt.Println("Now, what is your last name?")
	var second string
	fmt.Scanln(&second)

	fmt.Println("So, your full name is:")
	fmt.Print(first + " " + second)

	fmt.Printf("\n Is this correct? Type %q or %q:", "yes", "no")
	var yes string
	fmt.Scanln(&yes)

	if yes == "yes" {
		fmt.Println("aii bet")
	} else {
		fmt.Println("damn.")
	}

}
