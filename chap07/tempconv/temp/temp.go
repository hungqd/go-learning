package temp

import "fmt"

type Celcius float64
type Fahrenheit float64

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

const (
	AbsoluteZero Celcius = -273.15
	FreezingC    Celcius = 0
	BoilingC     Celcius = 100
)

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
