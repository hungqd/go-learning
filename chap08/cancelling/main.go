package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func doSleep(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Done....")
		return
	case <-time.After(1 * time.Hour):
		fmt.Println("Slept for an hour")
	}
}

func main() {
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 10; i++ {
		go doSleep(context)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
	fmt.Printf("Error: %v", context.Err())
	fmt.Println("end")
}
