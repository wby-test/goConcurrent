package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	//time.Timer{} :不能直接对time.Timer进行初始化，因为Timer中有私有数据
	timer := time.NewTimer(2 *time.Second)
	//timer.Stop()
	fmt.Println(reflect.TypeOf(timer))
	fmt.Printf("present time: %v.\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("expiration time: %v.\n", expirationTime)
	fmt.Printf("stop time: %v.\n", time.Now())
}
