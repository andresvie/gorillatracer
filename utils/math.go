package utils

import "math"

func Clamp(value, min, max REAL) REAL {
	newMin := math.Max(float64(value), float64(min))
	return REAL(math.Min(float64(max), newMin))
}
