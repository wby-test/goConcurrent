package main

import (
	"fmt"
	"time"
)

func main() {
	//var a int
	//fmt.Println(time.Now())
	//ticker := time.NewTicker(500 * time.Millisecond)
	//for b:= range ticker.C {
	//	fmt.Println(a, b)
	//	a += 2
	//	if a > 10 {
	//		break
	//	}
	//}
	//fmt.Println(time.Now())

	ticker := time.NewTicker(500 * time.Millisecond)
	a := <- ticker.C
	b := <- ticker.C
	fmt.Println(a, "\n", b)

}
