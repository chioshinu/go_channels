package main

func deadlock() {
	ch := make(chan int)
	ch <- 5
}
