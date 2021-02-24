package main

import (
	"fmt"
	"net/http"
)

func main() {

	c := make(chan string)

	websites := []string{
		"https://github.com/",
		"https://golang.org/",
		"https://www.google.co.in/",
	}

	for _, website := range websites {
		go getWebsite(website, c)
	}

	for msg := range c {
		fmt.Println(msg)
	}
}
func getWebsite(website string, c chan string) {
	if _, err := http.Get(website); err != nil {
		c <- website + "is down"

	} else {
		c <- website + "is up"
	}

}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:
// https://github.com/is up
// https://www.google.co.in/is up
// https://golang.org/is up
