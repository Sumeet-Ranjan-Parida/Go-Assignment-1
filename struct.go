package main

import (
	"fmt"
)

type Address struct {
	city  string
	state string
}

type User struct {
	name    string
	age     int
	address Address
}

func main() {
	u1 := User{
		name: "Naveen",
		age:  50,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	u2 := User{
		name: "Akash",
		age:  36,
		address: Address{
			city:  "Chennai",
			state: "Tamil Nadu",
		},
	}

	u3 := User{
		name: "Sam",
		age:  24,
		address: Address{
			city:  "Mumbai",
			state: "Maharashtra",
		},
	}

	u4 := User{
		name: "Mahesh",
		age:  40,
		address: Address{
			city:  "Bengaluru",
			state: "Karnataka",
		},
	}

	fmt.Println("Name:", u1.name)
	fmt.Println("Age:", u1.age)
	fmt.Println("City:", u1.address.city)
	fmt.Println("State:", u1.address.state)

	fmt.Println("\nName:", u2.name)
	fmt.Println("Age:", u2.age)
	fmt.Println("City:", u2.address.city)
	fmt.Println("State:", u2.address.state)

	fmt.Println("\nName:", u3.name)
	fmt.Println("Age:", u3.age)
	fmt.Println("City:", u3.address.city)
	fmt.Println("State:", u3.address.state)

	fmt.Println("\nName:", u4.name)
	fmt.Println("Age:", u4.age)
	fmt.Println("City:", u4.address.city)
	fmt.Println("State:", u4.address.state)
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//Name: Naveen
// Age: 50
// City: Chicago
// State: Illinois

// Name: Akash
// Age: 36
// City: Chennai
// State: Tamil Nadu

// Name: Sam
// Age: 24
// City: Mumbai
// State: Maharashtra

// Name: Mahesh
// Age: 40
// City: Bengaluru
// State: Karnataka
