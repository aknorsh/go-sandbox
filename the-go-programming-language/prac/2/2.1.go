package main

import "fmt"

// Celsius tempature
type Celsius float64

// Fahrenheit tempature
type Fahrenheit float64

// Kelvin tempature
type Kelvin float64

// Some useful constant tempature
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// CToF converts Celcius -> Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit -> Celcius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts Kelvin -> Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

// CToK converts Celsius -> Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func main() {
	var c Celsius
	c = 27.5
	f := Fahrenheit(CToF(c))
	k := Kelvin(CToK(c))

	fmt.Printf("c: %v, f: %v, k: %v\n", c, f, k)
}
