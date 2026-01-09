// 补充kelvin及其转化
package tempconv

import "fmt"

type Kelvin float64

const (
	AbsoluteZeroK Kelvin = 0.0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func (k Kelvin) string() string {
	return fmt.Sprintf("%g°K", k)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
