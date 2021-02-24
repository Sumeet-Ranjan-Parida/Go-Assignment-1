package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Error Handling: To check a file exists or not\n")
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println("Error: File not found")
		return
	}
	fmt.Println(f.Name(), "Opened Successfully")
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//Error Handling: To check a file exists or not

//Error: File not found
