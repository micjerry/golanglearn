package tempconv

import (
	"fmt"
	"flag"
)

type Celsius float64
type Fahrenheit float64
type celsiusFlag struct { Celsius }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &uint)
	switch uint {
		case "C", "°C":
			f.Celsius = Celsius(value)
			return nil
		case "F", "°F":
			f.Celsius = FToC(Fahrenheit(value))
			return nil
	}
	
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

