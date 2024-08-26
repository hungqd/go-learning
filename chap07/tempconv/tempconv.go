package main

import (
	"flag"
	"fmt"

	"github.com/hungqd/go-learning/chap07/tempconv/temp"
)

type celsiusFlag struct{ temp.Celcius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celcius = temp.Celcius(value)
		return nil
	case "F", "°F":
		f.Celcius = temp.FToC(temp.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c *celsiusFlag) String() string {
	return "example value"
}

func CelciusFlag(name string, value temp.Celcius, usage string) *temp.Celcius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcius
}

var tempF = CelciusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*tempF)
}
