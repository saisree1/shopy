package util

import "math"

func RoundFloat(val float64) float64 {
	ratio := math.Pow(10, float64(2))
	return math.Round(val*ratio) / ratio
}
