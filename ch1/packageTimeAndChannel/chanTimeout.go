package main

import (
	"fmt"
	"time"
)

type  A struct{
	a  <- chan int
	b chan <- int
}
func main() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	//定时器的复用
	timeout := 500 * time.Millisecond
	var timer *time.Timer

	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}

		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("chan closed")
				return
			}
			fmt.Println("recieved: ", e)
		case <- timer.C:
			fmt.Println("timeout")

		}
	}

}
