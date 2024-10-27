package main

import "fmt"

// Type declarations
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) Stringify() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	c := CToF(BoilingC)
	fmt.Printf("Boiling point = %g°F or %g°C\n", c, FToC(c))

	d := FToC(212.0)
	fmt.Println(d.Stringify())
}
