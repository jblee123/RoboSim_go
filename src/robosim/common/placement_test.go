package common

import (
	"testing"
)

func TestPlacementString(t *testing.T) {
	tests := []struct {
		input Placement
		want  string
	}{
		{
			Placement{Vector{0, 0, 0}, 0},
			"[(0.00, 0.00, 0.00), 0.00]",
		},
		{
			Placement{Vector{1, 2, -3}, 45.6723},
			"[(1.00, 2.00, -3.00), 45.67]",
		},
		{
			Placement{Vector{1.2, 2.45, -3.67}, 350.5689},
			"[(1.20, 2.45, -3.67), 350.57]",
		},
		{
			Placement{Vector{1.20, 2.451, 3.456}, 450},
			"[(1.20, 2.45, 3.46), 450.00]",
		},
		{
			Placement{Vector{1.20, 2.451, 3.456}, -20},
			"[(1.20, 2.45, 3.46), -20.00]",
		},
		{
			Placement{Vector{1.20, 2.451, 3.456}, -450},
			"[(1.20, 2.45, 3.46), -450.00]",
		},
	}

	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf("[%v, %v].String() = %q, want %q",
				test.input.Pos, test.input.Heading,
				got, test.want)
		}
	}
}
