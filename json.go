package main

import (
	"encoding/json"
	"fmt"
)

type Preson struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Gender string   `json:"gender"`
	City   []string `json:"city"`
}

func main() {

	p1 := &Preson{Name: "Sumeet", Age: 22, Gender: "Male", City: []string{"Chennai", "TN"}}
	data, _ := json.Marshal(p1)
	fmt.Println("Encode (marshal) struct to JSON:")
	fmt.Println(string(data))

	p2 := `{"Name":"Sumeet","Age":22,"Gender":"Male","City":["Chennai","TN"]}`
	data2 := &Preson{}

	json.Unmarshal([]byte(p2), data2)
	fmt.Println("\nDecode (unmarshal) JSON to struct:")
	fmt.Println(data2)
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//Encode (marshal) struct to JSON:
//{"name":"Sumeet","age":22,"gender":"Male","city":["Chennai","TN"]}

//Decode (unmarshal) JSON to struct:
//&{Sumeet 22 Male [Chennai TN]}
