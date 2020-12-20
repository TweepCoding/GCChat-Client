package util

import "math"

const (
	WIDTH  int     = 1200
	HEIGHT int     = 800
	MARGIN float64 = 0.1
)

func GetIntPortion(whole int, fraction float64) int {
	return int(math.Round(float64(whole) * fraction))
}