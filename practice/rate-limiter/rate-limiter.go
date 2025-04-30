package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 1)

	for req := range requests {
		limiter.Wait(context.Background())
		fmt.Println("request:", req, time.Now())
	}

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	burstyLimiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 3)

	for req := range burstyRequests {
		burstyLimiter.Wait(context.Background())
		fmt.Println("bursty request:", req, time.Now())
	}
}
