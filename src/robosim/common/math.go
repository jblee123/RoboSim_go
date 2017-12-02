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

func WrapInInterval(value, interval_start, interval_end float64) float64 {
	if value < interval_start || value >= interval_end {
		length := interval_end - interval_start
		value = value - math.Floor((value-interval_start)/length)*length
	}

	return value
}

func NormalizeAngle(angle float64) float64 {
	return WrapInInterval(angle, 0, 360)
}
