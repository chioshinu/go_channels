package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/cheggaaa/pb"
)

func process(number int) {
	// Seed the random number generator using the current time (nanoseconds since epoch)
	rand.Seed(time.Now().UnixNano())

	// Much harder to predict...but it is still possible if you know the day, and hour, minute...
	delay := rand.Intn(100) + number
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func iterate_process() {
	bar := pb.New(1000)
	bar.Start()
	for i := 0; i <= 1000; i++ {
		process(i)
		bar.Increment()
	}
	bar.Finish()
}

func parallel_process() {
	workers := 10
	c := make(chan os.Signal, 1)
	t := make(chan int, 1)
	workChan := make(chan int, 1000)

	signal.Notify(c, syscall.SIGINT)

	for i := 0; i < 1000; i++ {
		workChan <- i
	}


	bar := pb.New(1000)
	bar.Start()

	fmt.Printf("Start")

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(bar *pb.ProgressBar) {
			defer wg.Done()
			cancelled(c, t)
			for {
				if cancelled(c, t) {
					return
				}
				select {
				case w := <-workChan:
					if w < 30 {
						t <- w
					}
					process(w)
					bar.Increment()
				default:
					return
				}
			}
		}(bar)
	}
	wg.Wait()
	bar.Finish()
	fmt.Println("All goroutines stopped")
}

func cancelled(c chan os.Signal, t chan int) bool {
	if len(c) != 0 || len(t) != 0 {
		return true
	}

	return false
}
