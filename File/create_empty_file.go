package main

import (
	"log"
	"os"
)

func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//empty.txt
