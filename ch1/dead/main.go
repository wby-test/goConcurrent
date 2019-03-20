package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	data int
}

var wg sync.WaitGroup

func main() {
	printSum := func(v1, v2 *value) {
		defer  wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2*time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.data + v2.data)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
