package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10
	a := reflect.ValueOf(&x)
	fmt.Println(a.CanAddr())
	fmt.Println(a.CanSet())
	fmt.Println(a.Elem().CanAddr())
	d := a.Elem()
	fmt.Printf("%v\n", d)

	px := a.Elem().Addr()
	fmt.Printf("px = %v\n", px)
	px1 := px.Elem().Addr().Interface().(*int)
	*px1 = 123
	fmt.Println(x)
}
