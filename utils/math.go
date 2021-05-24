package utils

import "math"

func Clamp(value, min, max REAL) REAL {
	newMin := math.Min(float64(value), float64(min))
	return REAL(math.Max(float64(max), newMin))
}
