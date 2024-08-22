package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("\nExercise 1")
	fmt.Println(os.Args[0])

	fmt.Println("\nExercise 2")
	for idx, arg := range os.Args {
		fmt.Printf("Index %d: %s\n", idx, arg)
	}

	fmt.Println("\nExercise 3")

	var start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], "\n"))
	fmt.Printf("Using strings.Join. Execute time: %dμs\n", time.Since(start).Microseconds())

	start = time.Now()
	for idx, arg := range os.Args {
		fmt.Printf("Index %d: %s\n", idx, arg)
	}
	fmt.Printf("Using for range. Execute time: %dμs\n", time.Since(start).Microseconds())

	start = time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Using for loop. Execute time: %dμs\n", time.Since(start).Microseconds())
}
