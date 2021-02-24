package main

import (
	"fmt"
	"regexp"
)

func main() {
	var password string

	fmt.Println("Verify your Password\n")

	fmt.Printf("Enter the password: ")
	fmt.Scanln(&password)

	match, _ := regexp.Compile("pass123")
	fmt.Println(match.MatchString(password))
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//Verify your Password

//Enter the password: pass123
//true
