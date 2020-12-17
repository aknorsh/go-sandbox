package tempconv

// CToF converts Celcius -> Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit -> Celcius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
