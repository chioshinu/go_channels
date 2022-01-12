package main

import (
	"math/rand"
	"time"
)

func process() {
	// Seed the random number generator using the current time (nanoseconds since epoch)
	rand.Seed(time.Now().UnixNano())

	// Much harder to predict...but it is still possible if you know the day, and hour, minute...
	delay := rand.Intn(1000)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func main() {
	bar := ProgressBar.New(1000)
	bar.Start()
	for i := 0; i <= 1000; i++ {
		process()

	}
}
