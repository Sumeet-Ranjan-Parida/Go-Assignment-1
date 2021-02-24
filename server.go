package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Ready,")
	fmt.Fprint(w, " Enter myname to view your name")
}

func myName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Sumeet")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/myname", myName)
	http.ListenAndServe(":8080", nil)

}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:
//localhost:8080  Server is Ready, Enter myname to view your name
//localhost:8080/myname Hello, Sumeet
