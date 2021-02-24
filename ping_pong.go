package main

import (
	"fmt"
	"time"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
	}
}

// - Sumeet Ranjan Parida (Batch - 9A)

//Output:

//ping
//pong
//ping
//pong
//ping
//pong
