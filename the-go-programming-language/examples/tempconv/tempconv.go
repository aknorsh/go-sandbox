package tempconv

import "fmt"

// Celsius tempature
type Celsius float64

// Fahrenheit tempature
type Fahrenheit float64

// Some useful constant tempature
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
