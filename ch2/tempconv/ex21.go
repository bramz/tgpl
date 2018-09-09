/*
Exercise 2.1:
Add types, constants, and functions to tempconv for processing
temperatures in the Kelvin scale, where zero Kelvin is -273.27C and
a difference of 1K has the same magnitude as 1C.
*/
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Kelvin  = -273.27
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.27) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.27) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC()) }
func FToK(f Fahrenheit) Kelvin { return CToF(KToC()) }
