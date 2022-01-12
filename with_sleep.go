package main

import (
	"fmt"
	"time"
)

func with_sleep_hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func with_sleep() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go with_sleep_hello(done)
	<-done
	fmt.Println("Main received data")
}
