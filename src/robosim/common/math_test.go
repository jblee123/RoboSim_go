package common

import (
	"math"
	"testing"
)

const (
	deg45 = 45.0
	rad45 = deg45 * (2 * math.Pi / 360)
)

func TestDegToRad(t *testing.T) {
	if rad := degToRad(deg45); rad != rad45 {
		t.Errorf("degToRad(%v) = %v, want %v", deg45, rad, rad45)
	}
}

func TestRadToDeg(t *testing.T) {
	if deg := radToDeg(rad45); deg != deg45 {
		t.Errorf("radToDeg(%v) = %v, want %v", rad45, deg, deg45)
	}
}

func TestNormalizeAngle(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{-5, 355},
		{0, 0},
		{180, 180},
		{360, 0},
		{365, 5},
		{725, 5},
	}

	for _, test := range tests {
		if got := normalizeAngle(test.input); got != test.want {
			t.Errorf("normalizeAngle(%v) = %v, want %v",
				test.input, got, test.want)
		}
	}
}
