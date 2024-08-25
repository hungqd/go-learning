package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := test()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
	}
}

func test() (value int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	panic("Test panic")
}
