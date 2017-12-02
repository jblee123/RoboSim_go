package behaviors

import (
	"math"
	"math/rand"
	"robosim/common"
)

// Wander TODO: comment
type Wander struct {
	Behavior
	Persistence        IntGetter
	sameDirectionCount int
}

// Output TODO: comment
func (b *Wander) Output() common.Vector {
	output := b.Behavior.Output()
	return output.(common.Vector)
}

// CreateWander TODO: comment
func CreateWander(name string, persistence IntGetter) *Wander {

	var behavior Wander

	outputFunc := func() {
		computeWander(&behavior)
	}

	behavior.Behavior.Init(name, nil, outputFunc)
	behavior.Persistence = persistence
	behavior.sameDirectionCount = 0

	return &behavior
}

func computeWander(behavior *Wander) {

	if behavior.sameDirectionCount >= behavior.Persistence.Output() {
		behavior.sameDirectionCount = 0
	}

	if behavior.sameDirectionCount == 0 {
		output := common.Vector{X: 1, Y: 0, Z: 0}
		dir := common.Radians(rand.Float64() * 2 * math.Pi)
		behavior.output = output.RotateZ(dir)
	}

	behavior.sameDirectionCount++
}
