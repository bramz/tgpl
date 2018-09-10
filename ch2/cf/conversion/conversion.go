package conversion

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64
type Pound float64
type Kilo float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	OneK          Kelvin  = 273.27
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
func (f Feet) String() string       { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g meters", m) }
func (p Pound) String() string      { return fmt.Sprintf("%g lbs", p) }
func (k Kilo) String() string       { return fmt.Sprintf("%g kg", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.27) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.27) }

func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
func FToK(f Fahrenheit) Kelvin { return CToK(CToF(f)) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

func PToK(p Pound) Kilo { return Kilo(p * 0.453592) }
func KToP(k Kilo) Pound { return Pound(k / 0.453592) }
