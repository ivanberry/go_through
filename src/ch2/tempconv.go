package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreeingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%gC", c)}
func (c Fahrenheit) String() string { return fmt.Sprintf("%gF", f)}