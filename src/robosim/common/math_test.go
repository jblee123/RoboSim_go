package common

import (
	"testing"
)

func TestMinI(t *testing.T) {
	type Input struct{ a, b int }

	tests := []struct {
		input Input
		want  int
	}{
		{Input{-5, -4}, -5},
		{Input{-5, 4}, -5},
		{Input{-4, -5}, -5},
		{Input{4, -5}, -5},
	}

	for _, test := range tests {
		if got := MinI(test.input.a, test.input.b); got != test.want {
			t.Errorf("MinI(%v, %v) = %v, want %v",
				test.input.a, test.input.b, got, test.want)
		}
	}
}
