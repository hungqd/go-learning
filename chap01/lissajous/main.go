package main

import (
	"fmt"
	"os"

	lissajous "github.com/hungqd/go-learning/chap01/lissajous/lib"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
		} else {
			lissajous.Lissajous(file)
			file.Close()
		}
	} else {
		lissajous.Lissajous(os.Stdout)
	}
}
