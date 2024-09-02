package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())
	name := stdout.FieldByName("name")
	fmt.Println(name.String())
	fmt.Println(name.CanAddr(), name.CanSet())
	name.SetString("Hung") // panic: unexported field
}
