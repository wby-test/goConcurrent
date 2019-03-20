package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		intChan <- 1
	}()

	select {
	case e := <-intChan:
		fmt.Println("recieved: ", e)
	case <- time.NewTimer(5000 * time.Millisecond).C:
		fmt.Println("timeout")

	}
	time.AfterFunc(100*time.Millisecond, func() {
		fmt.Println("aaa")
	})
	time.Sleep(1 * time.Second)
}
