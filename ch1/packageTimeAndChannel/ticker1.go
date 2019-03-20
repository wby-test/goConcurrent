package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int , 1)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("end sender")
	}()

	var sum int
	for e := range intChan {
		fmt.Printf("recieved : %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v.\n", sum)
			break
		}
	}
	fmt.Println("end recieved")
}
