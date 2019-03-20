package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)
	go func() {
	
		for i := 0; i < 10000000; i++ {

		}
		fmt.Println("join" +
			"ddd")
		ch <- 5
	}()
	//<- ch
	//go func() {
	//	fmt.Println("no schduler")
	//}()
	//go fmt.Println("saa")
	fmt.Println(runtime.NumGoroutine())
}
