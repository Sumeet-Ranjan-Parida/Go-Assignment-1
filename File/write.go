package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//test.txt
