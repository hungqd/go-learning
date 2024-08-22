package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "quit" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
