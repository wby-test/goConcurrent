package main

import "fmt"

func main() {
	chanOwner := func() <-chan int{
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}
	resulfStream := chanOwner()
	for result := range resulfStream {
		fmt.Println(result)
	}

	fmt.Println("Done")
}
