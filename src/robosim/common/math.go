package common

import (
	"math"
)

func MinI(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func WrapInInterval(value, intervalStart, intervalEnd float64) float64 {
	if value < intervalStart || value >= intervalEnd {
		length := intervalEnd - intervalStart
		value = value - math.Floor((value-intervalStart)/length)*length
	}

	return value
}
